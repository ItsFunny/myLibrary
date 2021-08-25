/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/3/21 7:08 上午
# @File : demo.go
# @Description :
# @Attention :
*/
package main

import (
	"fmt"
	"time"

	// "github.com/libp2p/go-libp2p"
	"github.com/tendermint/tendermint/libs/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// ctx := context.Background()
	// node, err := libp2p.New(ctx)
	// if nil != err {
	// 	panic(err)
	// }
	// // print the node's listening addresses
	// fmt.Println("Listen addresses:", node.Addrs())
	//
	// // shut the node down
	// if err := node.Close(); err != nil {
	// 	panic(err)
	// }
	// testMap()
	// testC()
	// testCCS()
	testGoto()
	fmt.Println("exit")
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGKILL, syscall.SIGINT)
	<-sigCh
}


type T struct {
	sync.RWMutex
	m map[int]int
}

func testGoto() {
	i := 0
	for {
		i++
		if i&1 == 1 {
			goto s
			goto ss
		}
	}
	fmt.Println(2)
s:
	fmt.Println(1)
ss:
	fmt.Println("ss")
}

func testCCS() {
	cs := make([]chan int, 0)
	for i := 0; i < 10; i++ {
		cs = append(cs, make(chan int, 1))
	}
	go func() {
		for i := 0; i < 10; i++ {
			go func(index int) {
				for {
					time.Sleep(time.Second * time.Duration(index+1))
					cs[index] <- index
				}
			}(i)
		}
	}()
	go func() {
		for {
			for i, _ := range cs {
				select {
				case v := <-cs[i]:
					fmt.Println(v)
				default:
				}
			}
		}
	}()

}
func testC() {
	c := make(chan int)
	go func() {
		time.Sleep(time.Second * 4)
		c <- 1
		time.Sleep(time.Second * 3)
		close(c)
	}()
	go func() {
		for {
			select {
			case v, ok := <-c:
				if ok {
					fmt.Println(v)
				} else {
					fmt.Println("no")
				}
			}
			time.Sleep(time.Second)
		}
	}()

}

func testMap() {
	t := &T{
		m: make(map[int]int),
	}

	go func() {
		for i := 0; i < 100; i++ {
			go func() {
				for {
					t.Lock()
					r := rand.Intn(1000000000)
					t.m[r] = r
					t.Unlock()
					time.Sleep(time.Second)
				}
			}()
		}
	}()

	go func() {
		for {
			if _, exist := t.m[100000000000000]; !exist {
				t.RLock()
				if _, exist := t.m[100000000000000]; !exist {
					fmt.Println("不存在")
				}
				t.RUnlock()
			}
			// for _, v := range t.m {
			// 	fmt.Println(v)
			// }
			time.Sleep(time.Second)
		}

	}()

}
