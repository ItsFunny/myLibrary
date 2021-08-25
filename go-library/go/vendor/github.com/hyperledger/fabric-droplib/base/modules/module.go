/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/4/10 2:40 下午
# @File : module.go
# @Description :
# @Attention :
*/
package modules

import (
	"fmt"
	"strings"
)

type Module interface {
	fmt.Stringer
	Index() uint16
}

type module struct {
	name  string
	index uint16
}

func NewModule(name string, index uint16) module {
	name = strings.ToUpper(name)
	return module{
		index: index,
		name:  name,
	}
}

func (m module) String() string {
	return m.name
}

func (m module) Index() uint16 {
	return m.index
}

