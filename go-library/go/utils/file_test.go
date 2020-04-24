/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-24 09:26 
# @File : file_test.go
# @Description : 
*/
package utils

import (
	"fmt"
	"testing"
)

func TestGetFileSize(t *testing.T) {
	path := "/Users/joker/Downloads/fengkuangwaixingren.mp4"
	size, e := GetFileSize(path)
	fmt.Println(e)
	fmt.Println(size)
}
func TestIsFileOrDirExists(t *testing.T) {
	// path := "./hash.go"
	// fmt.Println(IsFileOrDirExists(path))
	var a [2]int
	fmt.Println(a[0])
}

func TestGetFilesBelownDirFilterBySuffix(t *testing.T) {
	dir := "/Users/joker"
	strings, e := GetFilesBelownDirFilterBySuffix(dir, ".gz")
	if nil != e {
		panic(e)
	}
	for _, s := range strings {
		fmt.Println(s)
	}
}

func TestGetAllFilesUnderDir(t *testing.T) {
	dirPath := "."
	files, e := GetAllFilesUnderDir(dirPath)
	if nil != e {
		panic(e)
	} else {
		for _, f := range files {
			fmt.Println(f.Name())
			f.Close()
		}
	}
}

