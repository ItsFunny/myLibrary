/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-08-19 09:19 
# @File : video.go
# @Description : 
*/
package video

import (
	bytes2 "bytes"
	"github.com/corona10/goimagehash"
	"github.com/mfonda/simhash"
	"gocv.io/x/gocv"
	"image/jpeg"
	"myLibrary/library/src/main/go/utils"
	"strconv"
	"sync"
)

const (
	// 只需要捕获最少的帧
	// 适用于很小的视频
	VIDEO_FRAME_LEVEL_LESS  = 0
	VIDEO_FRAME_LEVEL_MID   = 1
	VIDEO_FRAME_LEVEL_UPMID = 2
	VIDEO_FRAME_LEVEL_MOST  = 3

	// 10M
	VIDEO_SIZE_10M = 10 * 1 << 10
	// 100M
	VIDEO_SIZE_100M = 10 * VIDEO_SIZE_10M
	// 1G
	VIDEO_SIZE_1G = 10 * VIDEO_SIZE_100M
)

var (
	VIDEO_LEVEL_COUNT_ARRAY []int8
)

func init() {
	VIDEO_LEVEL_COUNT_ARRAY = make([]int8, 4)
	VIDEO_LEVEL_COUNT_ARRAY[VIDEO_FRAME_LEVEL_LESS] = 10
	VIDEO_LEVEL_COUNT_ARRAY[VIDEO_FRAME_LEVEL_MID] = 15
	VIDEO_LEVEL_COUNT_ARRAY[VIDEO_FRAME_LEVEL_UPMID] = 30
	VIDEO_LEVEL_COUNT_ARRAY[VIDEO_FRAME_LEVEL_MOST] = 100
}

func CompareVideosWithImg(filePath1, filePath2 string, funcName string, threshold int) (float64, error) {
	hashes, e := GetVideoFramesWithImg(filePath1, funcName)
	if nil != e {
		return 0.0, e
	}
	hashes2, e := GetVideoFramesWithImg(filePath2, funcName)
	if nil != e {
		return 0.0, e
	}

	return CompareImgHashes(hashes, hashes2, threshold)
}
func CompareImgHashes(hashes1, hashes2 []*goimagehash.ImageHash, threshold int) (float64, error) {
	l1 := len(hashes1)
	l2 := len(hashes2)
	count := 0
	less := l1
	lessesHashes := hashes1
	bigHashes := hashes2
	big := l2
	if l1 > l2 {
		less = l2
		lessesHashes = hashes2
		big = l1
		bigHashes = hashes1
	}
	for i := 0; i < less; i++ {
		for j := 0; j < big; j++ {
			distance, e := lessesHashes[i].Distance(bigHashes[j])
			if nil != e {
				return 0.0, e
			}
			if distance <= threshold {
				count++
				break
			}
		}
	}

	return float64(count) / float64(big), nil
}

func CompareVideos(filePath1, filePath2 string) (float64, error) {
	uint64s, e := GetVideoFramesWithSimHash(filePath1)
	if nil != e {
		return 0, e
	}
	chanuint64s, e := GetVideoFramesWithSimHash(filePath2)
	if nil != e {
		return 0, e
	}
	ffs := compareUints(uint64s, chanuint64s)
	return ffs, nil
	//
}

// 判断这个切片A在切片B中相似的百分比
// 相似的个数/总个数 即可
func CompareUints(l1, l2 []uint64, threshold uint8) float64 {
	count := 0
	for _, i := range l1 {
		for _, j := range l2 {
			hash1 := goimagehash.NewImageHash(i, goimagehash.DHash)
			hash2 := goimagehash.NewImageHash(j, goimagehash.DHash)
			distance, e := hash1.Distance(hash2)
			if nil != e {
				panic(e)
				return 0.0
			} else if distance <= int(threshold) {
				count++
				break
			}
			// compare := simhash.Compare(i, j)
			// if compare <= threshold {
			// 	count++
			// 	break
			// }
		}
	}

	return float64(count) / float64(len(l2))
}
func compareUints(chan1, chan2 chan uint64) float64 {
	l1 := make([]uint64, 0)
	for i1 := range chan1 {
		l1 = append(l1, i1)
	}
	l2 := make([]uint64, 0)
	for i2 := range chan2 {
		l2 = append(l2, i2)
	}
	return CompareUints(l1, l2, 6)
}

// func GetVideoFramesWithImg(filePath string, funName string) {
//
// }
func GetVideoFramesWithSimHash(filePath string) (chan uint64, error) {
	bytes, e := GetVideoFrames(filePath)
	if nil != e {
		return nil, e
	}
	wg := sync.WaitGroup{}
	l := len(bytes)
	result := make(chan uint64, l)
	defer close(result)
	wg.Add(l)
	for i := 0; i < l; i++ {
		go func(index int) {
			// image, e := getImgHash("a", bytes[index])
			// image, e := jpeg.Decode(bytes2.NewBuffer(bytes[index]))
			// if nil != e {
			// 	result <- 0
			// 	fmt.Println("发生了异常:", e.Error())
			// 	wg.Done()
			// } else {
			// 	result <- image.GetHash()
			// 	wg.Done()
			// }
			hash := simhash.Simhash(simhash.NewWordFeatureSet(bytes[index]))
			result <- hash
			wg.Done()
		}(i)
	}
	wg.Wait()

	return result, nil
}

func getImgHash(funcName string, bs []byte) (*goimagehash.ImageHash, error) {
	buffer := bytes2.NewBuffer(bs)
	image, e := jpeg.Decode(buffer)
	imgHash := new(goimagehash.ImageHash)
	if nil != e {
		return nil, e
	}
	switch funcName {
	case "a":
		imgHash, e = goimagehash.AverageHash(image)
	case "d":
		imgHash, e = goimagehash.DifferenceHash(image)
	}
	if nil != e {
		return nil, e
	}
	return imgHash, nil
}

func GetVideoFrames(filePath string) ([][]byte, error) {
	// 当level为less的时候默认为10
	level, e := getVideoFrameLevel(filePath)
	if nil != e {
		return nil, e
	}
	return getVideoFrames(filePath, level)
}

func getVideoFrameLevel(filePath string) (int, error) {
	size, e := utils.GetFileSize(filePath)
	if nil != e {
		return 0, e
	}
	if size < VIDEO_SIZE_10M {
		return VIDEO_FRAME_LEVEL_LESS, nil
	} else if size < VIDEO_SIZE_100M {
		return VIDEO_FRAME_LEVEL_MID, nil
	} else if size < VIDEO_SIZE_1G {
		return VIDEO_FRAME_LEVEL_UPMID, nil
	} else {
		return VIDEO_FRAME_LEVEL_MOST, nil
	}
}

// level 代表捕获帧的数量
// 获取某一个时刻的pic
// picCount: 需要保存的图片张数
func getVideoFrames(filePath string, level int) ([][]byte, error) {

	picCount := int(VIDEO_LEVEL_COUNT_ARRAY[level])
	result := make([][]byte, picCount)
	// load video
	vc, err := gocv.VideoCaptureFile(filePath)
	if err != nil {
		return nil, err
	}

	// fps是帧率，意思是每一秒刷新图片的数量，frames是一整段视频中总的图片数量。
	frames := vc.Get(gocv.VideoCaptureFrameCount)
	total := frames
	fps := vc.Get(gocv.VideoCaptureFPS)
	// 获取时间总长
	duration := frames / fps
	// fmt.Println(duration)
	// 递增的值
	loopAddFrequence := duration / float64(picCount)
	for i, j := 0.0, 0; j < picCount; i += loopAddFrequence {
		// Set Video frames
		// time/duration 获取到那个时间点的百分比
		frames = (i / duration) * total
		vc.Set(gocv.VideoCapturePosFrames, frames)
		img := gocv.NewMat()
		vc.Read(&img)
		gocv.IMWrite("/Users/joker/Desktop/temp/images/"+utils.GenerateUUID()+"----"+strconv.Itoa(j)+".jpg", img)
		result[j] = img.ToBytes()
		j++
	}

	return result, err
}

func GetVideoFramesWithImg(filePath string, funcName string) ([]*goimagehash.ImageHash, error) {
	// 当level为less的时候默认为10
	level, e := getVideoFrameLevel(filePath)
	if nil != e {
		return nil, e
	}
	return getVideoFramesWithGoimage(filePath, funcName, level)
}

// 通过goimage 获取hash
func getVideoFramesWithGoimage(filePath string, funName string, level int) ([]*goimagehash.ImageHash, error) {
	picCount := int(VIDEO_LEVEL_COUNT_ARRAY[level])
	// result := make([][]byte, picCount)
	result := make([]*goimagehash.ImageHash,0)
	// load video
	vc, err := gocv.VideoCaptureFile(filePath)
	if err != nil {
		return nil, err
	}

	// fps是帧率，意思是每一秒刷新图片的数量，frames是一整段视频中总的图片数量。
	frames := vc.Get(gocv.VideoCaptureFrameCount)
	total := frames
	fps := vc.Get(gocv.VideoCaptureFPS)
	// 获取时间总长
	duration := frames / fps
	// fmt.Println(duration)
	// 递增的值
	loopAddFrequence := duration / float64(picCount)
	for i, j := 0.0, 0; j < picCount; i += loopAddFrequence {
		// Set Video frames
		// time/duration 获取到那个时间点的百分比
		frames = (i / duration) * total
		vc.Set(gocv.VideoCapturePosFrames, frames)
		img := gocv.NewMat()
		vc.Read(&img)
		// gocv.IMWrite("/Users/joker/Desktop/temp/images/"+utils.GenerateUUID()+"----"+strconv.Itoa(j)+".jpg", img)
		// result[j] = img.ToBytes()
		image, err := img.ToImage()
		if nil != err {
			return result, err
		}
		imageHash := new(goimagehash.ImageHash)
		switch funName {
		case "a":
			imageHash, err = goimagehash.AverageHash(image)
		}
		if nil != err {
			return result, err
		}
		result = append(result, imageHash)
		j++
	}

	return result, err
}
