/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/4/10 11:54 上午
# @File : status.go
# @Description :
# @Attention :
*/
package constants

const (
	NONE    = 0
	STARTED = 1 << 0
	READY   = 1<<1 | STARTED

	STOP  = 1
	FLUSH = 1<<1 | STOP
)
