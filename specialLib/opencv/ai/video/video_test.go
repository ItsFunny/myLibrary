/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-08-19 10:07 
# @File : video_test.go
# @Description : 
*/
package video

import (
	"fmt"
	"testing"
)

func TestGetVideoFrames(t *testing.T) {
	filePath := "/Users/joker/Desktop/temp/Jay-test.flv"
	ints, e := GetVideoFramesWithSimHash(filePath)
	if nil != e {
		panic(e)
	}
	for v := range ints {
		fmt.Println(v)
	}
}
func TestCompareVideos(t *testing.T) {
	// filePath1 := "/Users/joker/Desktop/temp/a.mp4"
	filePath1 := "/Users/joker/Desktop/图片/测试/视频相似度测试/jay.flv"
	filePath2 := "/Users/joker/Desktop/图片/测试/视频相似度测试/jay2.flv"
	f, e := CompareVideos(filePath1, filePath2)
	if nil != e {
		panic(e)
	}
	fmt.Println(f)
}
func TestCompareVideosWithImg(t *testing.T) {
	filePath1 := "/Users/joker/Desktop/图片/测试/视频相似度测试/jay.flv"
	filePath2 := "/Users/joker/Desktop/图片/测试/视频相似度测试/jay.flv"
	f, e := CompareVideosWithImg(filePath1, filePath2,"a",6)
	if nil != e {
		panic(e)
	}
	fmt.Println(f)
}