/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-08-23 13:32 
# @File : DistributeLock.go
# @Description :  分布式锁
*/
package redis

import (
	"github.com/gomodule/redigo/redis"
	"myLibrary/library/src/main/go/converters"
	"myLibrary/library/src/main/go/utils"
	"sync"
	"time"
)

type DistributeLock struct {
	conn redis.Conn
	lock sync.Mutex
	// 业务过期时间,秒
	bussTimeOut int64

	// 是否开启保护模式
	protectionMode bool

	// token
	token string

}

func NewDistributeLock(conn redis.Conn,bussTimeOutSeconds int64) *DistributeLock {
	l := new(DistributeLock)
	l.bussTimeOut=bussTimeOutSeconds
	l.token = utils.GenerateUUID()
	return l
}

func (this *DistributeLock) TryLock(conn redis.Conn, specialKey string, value []byte, tryLockTimeOut, expire int64, isAlwaysWait bool) (bool, error) {
	startTime := time.Now().Unix()
	bussExpireTime := startTime + this.bussTimeOut
	value = append(value, converter.BigEndianInt642Bytes(bussExpireTime)...)
	isGet := false
	if isAlwaysWait {
		for {
			if isGet {
				break
			}
			for {
				_, err := conn.Do("SET", specialKey, value, this.token, "EX", expire, "NX")
				if nil != err && err != redis.ErrNil {
					return false, err
				} else {
					isGet = true
				}
			}
		}
	} else {
		for {
			d := time.Now().Unix() - startTime
			if isGet || d >= tryLockTimeOut {
				break
			}
			_, err := conn.Do("SET", specialKey, value, "EX", expire, "NX")
			if nil != err {
				if err == redis.ErrNil {
					// 说明设置失败
				}
				return false, err
			}
			isGet = true
		}
	}

	// 保护模式,则需要判断业务时间和过期时间
	if this.protectionMode {
		// 最后再进行一次获取
		if !isGet {
			reply, err := conn.Do("GET", specialKey)
			if nil != err {
				// 说明key不存在,则尝试再次获取
				if err == redis.ErrNil {
					// 尝试再次获取
					if b, err := this.lockAgain(conn, specialKey, value, expire); nil != err {
						return false, err
					} else {
						return b, nil
					}
				}
				return false, err
			}
			// 说明获取到了,判断是否过期
			// 判定业务时间是否已经过期了
			bussExpire := this.getBussExpireTime(reply.([]byte))
			now := time.Now().Unix()
			if now-bussExpire < 0 {
				// 说明没过期,直接返回
				return false, nil
			}
			// 说明过期了
			// 则这里需要锁住,然后删除这个key
			// 并且需要double check
			// 若不加双重检测,则发现过期直接删除,同时删除之前,key刚好没了,有新的key进入了,则会把这个key给删了
			// 所以需要双重检测(并且因为release的时候是需要加锁的,所以能防止)
			this.lock.Lock()
			defer this.lock.Unlock()
			if reply, err := conn.Do("GET", specialKey); nil != err {
				if err == redis.ErrNil {
					return false, nil
				} else {
					return false, err
				}
			} else {
				if now-this.getBussExpireTime(reply.([]byte)) < 0 {
					// 这个错误不需要处理
					conn.Do("DELETE", specialKey)
				}
			}
		}
	}
	return isGet, nil
}
func (this *DistributeLock) lockAgain(conn redis.Conn, specialKey string, value []byte, expire int64) (bool, error) {
	startTime := time.Now().Unix()
	bussExpireTime := startTime + this.bussTimeOut
	value = append(value, converter.BigEndianInt642Bytes(bussExpireTime)...)
	_, err := conn.Do("SET", specialKey, value, "EX", expire, "NX")
	if nil != err {
		if err == redis.ErrNil {
			return false, nil
		}
	} else {
		return false, err
	}
	return true, nil
}

var unlockScript = redis.NewScript(1, `
	if redis.call("get", KEYS[1]) == ARGV[1]
	then
		return redis.call("del", KEYS[1])
	else
		return 0
	end
`)

func (this *DistributeLock) TryReleaseLock(conn redis.Conn, specialKey string) (err error) {
	_, err = unlockScript.Do(conn, specialKey, this.token)
	return
}

func (this *DistributeLock) getBussExpireTime(value []byte) int64 {
	l := len(value)
	bytes := value[l-7 : l]
	return int64(converter.BigEndianBytes2Int64(bytes))
}
