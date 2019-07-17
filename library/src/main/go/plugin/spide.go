/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-04 13:37 
# @File : spide.go
# @Description :    爬虫
*/
package plugin

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"myLibrary/library/src/main/go/utils"
	"net/http"
	"path/filepath"
	"strings"
)

// TODO 错误日志统一写到文件中

const (
	SPIDE_TYPE_ARTICLE = iota
	SPIDE_TYPE_VIDEO
	SPIDE_TYPE_IMAGE
)

var (
	IMAGES_SPIDE_VALIDER = func(spide SpideNode) (string, error) {
		index := strings.LastIndex(spide.Url, ".")
		if index == -1 {
			return "", errors.New("如果是图片,结尾必须是 xxx.png/img等常规格式")
		}
		l := len(spide.Url)
		if index == l-1 {
			return "", errors.New("结尾不可以.结尾")
		}
		suffix := utils.SubStringBetween(spide.Url, index+1, l)
		suffix = strings.ToLower(suffix)
		if suffix == "png" || suffix == "jpg" || suffix == "jpeg" {
			return suffix, nil
		} else {
			return suffix, errors.New("结尾必须是png,jpg,jpeg等")
		}
	}
)

type SpideReqValid = func(spide Spide) (string, error)

// FIXME 更改名字
type Spide struct {
	IsDebug bool `yaml:"is_debug"`
	//
	// DefaultPath string `yaml:"default_path"`
	// 存储路径,通常而言是一个绝对路径
	StorePath string `yaml:"store_path"`
	// 资源映射路径 供外界访问的
	MappingPath string
}

func NewSpider(storePath, mappingPath string) *Spide {
	s := &Spide{
		IsDebug:     false,
		StorePath:   storePath,
		MappingPath: mappingPath,
	}
	return s
}

// 爬虫爬取对象
type SpideNode struct {
	Type int    `form:"type" json:"type"`
	Url  string `form:"url" json:"url"`
	// StorePath string // 存储地址
	NewName string `form:"new_name" json:"new_name"` // 新的名字

	// 能唯一代表一个用户的
	SpecialUUID string `form:"special_uuid" json:"special_uuid"`
}

// 爬虫爬取集合
type SpideList struct {
	Items []SpideNode `form:"items" json:"items"`
}
type SpideResult struct {
	StorePath   string
	MappingPath string
	Url         string
}

// path 存储的位置
// err 错误位置
func (this *Spide) SpideOne(spide SpideNode) (result SpideResult, err error) {
	var storePath, mappingPath string
	basePath := this.StorePath + spide.SpecialUUID
	mappingBasePath := this.MappingPath + spide.SpecialUUID
	if !utils.IsFileOrDirExists(basePath) {
		// 创建
		err = utils.CreateMultiFileDirs(basePath)
		if nil != err {
			log.Println("无法创建多级目录:", err)
			return
		}
	}
	if spide.Type == SPIDE_TYPE_ARTICLE {
		storePath, mappingPath, err = this.downloadArticle(basePath, mappingBasePath, spide)
	} else {
		storePath, mappingPath, err = this.downOthers(basePath, mappingBasePath, spide)
	}
	result.StorePath = storePath
	result.MappingPath = mappingPath
	result.Url = spide.Url

	return
}

// 批量爬取
func (this *Spide) BatchSpide(list SpideList) (result []SpideResult, err error) {
	// FIXME 协程处理
	for _, sn := range list.Items {
		res, e := this.SpideOne(sn)
		if nil != e {
			log.Printf("爬取[%s]的时候出错:%s", sn.Url, e.Error())
			return
		}
		result = append(result, res)
	}
	return
}

func (this *Spide) downloadArticle(basePath string, baseMappingPath string, spide SpideNode) (storePath string, mappingPath string, err error) {
	resp, err := http.Get(spide.Url)
	if nil != err {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New("请求失败")
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return
	}
	suffix := ".html"
	// 存到文件中
	if spide.NewName == "" {
		spide.NewName = utils.GenerateUUID()
	}

	storePath = basePath + string(filepath.Separator) + "articles"
	mappingPath = baseMappingPath + string(filepath.Separator) + "articles"

	if !utils.IsFileOrDirExists(storePath) {
		err = utils.CreateMultiFileDirs(storePath)
		if nil != err {
			log.Println("无法创建多级目录:", err)
			return
		}
	}
	tPath := storePath + string(filepath.Separator) + spide.NewName + suffix
	tmpath := mappingPath + string(filepath.Separator) + spide.NewName + suffix
	if utils.IsFileOrDirExists(tPath) {
		log.Println("文件已经存在,重新生成uuid")
		spide.NewName = utils.GenerateUUID()
		storePath = storePath + string(filepath.Separator) + spide.NewName + suffix
		mappingPath = mappingPath + string(filepath.Separator) + spide.NewName + suffix
	} else {
		storePath = tPath
		mappingPath = tmpath
	}

	err = ioutil.WriteFile(storePath, bytes, 0777)
	if nil != err {
		// TODO 日志
		str := fmt.Sprintf("写入到文件失败:%s,结构体为{%+v}", err.Error(), spide)
		log.Println(str)
		return
	}
	return
}

func (this *Spide) downOthers(basePath, baseMappingPath string, spide SpideNode) (path, mappingPath string, err error) {
	cmds := "you-get "
	if this.IsDebug {
		cmds += " -d "
	}
	typeName := "videos"
	suffix := ""
	if spide.Type == SPIDE_TYPE_IMAGE {
		typeName = "images"
		s, err := utils.GetLowerSuffixFromUrl(spide.Url)
		if nil != err {
			return "", "", err
		}
		suffix = s

	}
	path = basePath + string(filepath.Separator) + typeName + string(filepath.Separator)
	mappingPath = baseMappingPath + string(filepath.Separator) + typeName + string(filepath.Separator)
	if !utils.IsFileOrDirExists(path) {
		err = utils.CreateMultiFileDirs(path)
		if nil != err {
			log.Println("无法创建多级目录:", err)
			return
		}
	}
	if spide.NewName == "" {
		spide.NewName = utils.GenerateUUID()
	} else if utils.IsFileOrDirExists(path + spide.NewName + suffix) {
		log.Printf("文件名称:%s 已经存在,重新生成为uuid", spide.NewName)
		spide.NewName = utils.GenerateUUID()
	}

	// path = path + string(filepath.Separator) + spide.NewName

	cmds += "-o " + path
	cmds += " -O " + spide.NewName
	cmds += " " + spide.Url

	path = path + spide.NewName + suffix
	mappingPath = mappingPath + spide.NewName + suffix

	err = ExecCmdWithLog(cmds)

	return
}
