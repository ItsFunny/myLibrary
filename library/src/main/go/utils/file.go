/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-04 14:28 
# @File : file.go
# @Description : 
*/
package utils

import (
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// 判断所给路径文件/文件夹是否存在
func IsFileOrDirExists(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func CreateMultiFileDirs(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 获取当前可执行文件的路径
func GetCurrentExecPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	// fmt.Println("path111:", path)
	if runtime.GOOS == "windows" {
		path = strings.Replace(path, "\\", "/", -1)
	}
	// fmt.Println("path222:", path)
	i := strings.LastIndex(path, "/")
	if i < 0 {
		return "", errors.Errorf(`Can't find "/" or "\".`)
	}
	// fmt.Println("path333:", path)
	return string(path[0 : i+1]), nil
}
