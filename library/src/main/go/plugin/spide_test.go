/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-04 14:42 
# @File : spide_test.go
# @Description : 
*/
package plugin

import (
	"fmt"
	"testing"
)

func TestSpide_SpideOne(t *testing.T) {
	spideHelper := &Spide{
		IsDebug:     true,
		DefaultPath: "/Users/joker/Desktop/",
	}
	sn := SpideNode{
		Type:        SPIDE_TYPE_IMAGE,
		Url:         "https://stallman.org/rms.jpg",
		NewName:     "tesss",
		SpecialUUID: "joker",
	}
	path, err := spideHelper.SpideOne(sn)
	if nil != err {
		panic(err)
	}
	fmt.Println(path)
}

func TestSpide_SpideOne_ARTICLE(t *testing.T) {
	spideHelper := &Spide{
		IsDebug:     true,
		DefaultPath: "/Users/joker/Desktop/",
	}
	sn := SpideNode{
		Type:        SPIDE_TYPE_ARTICLE,
		Url:         "https://golangtc.com/t/543fb94f421aa94691000070",
		NewName:     "tesss",
		SpecialUUID: "joker",
	}
	path, err := spideHelper.SpideOne(sn)
	if nil != err {
		panic(err)
	}
	fmt.Println(path)
}
