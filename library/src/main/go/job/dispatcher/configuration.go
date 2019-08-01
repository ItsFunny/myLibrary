/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-31 17:46 
# @File : configuration.go
# @Description : 
*/
package dispatcher

import "myLibrary/library/src/main/go/job/worker"

type DispatcherConfiguration struct {
	// 有多少个worker同时运行
	Workers []worker.Worker
	// job缓存池多大
	JobCount int
}
