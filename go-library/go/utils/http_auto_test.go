/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-12 16:47 
# @File : http_test3.go
# @Description :
  用于自动批量校验深思返回的数据
# @Attention : 
*/
package utils

import (
	"github.com/prometheus/common/log"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var (
	cfgs []*BidsunKeyInfoConfiguration
)

func initByPath(path string) {
	defer func() {
		if e := recover(); nil != e {
			log.Fatal("执行该计划失败,plan=" + path)
		}
	}()
	bytes, e := ioutil.ReadFile(path)
	if nil != e {
		log.Fatal("读取文件失败:" + e.Error() + ",path=" + path)
	}
	parse(bytes)
	httpttt()
}

func Test_ttttttt(t *testing.T) {
	dir := "/Users/joker/Java/ebidsun/bidsun-tool/bidsun-tool-doe-server-client/infos"
	infos, e := ioutil.ReadDir(dir)
	if nil != e {
		log.Fatal(e)
	}
	for _, info := range infos {
		path := filepath.Join(dir, info.Name())
		initByPath(path)
	}

}
