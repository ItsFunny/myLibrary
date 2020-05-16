/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-26 14:13 
# @File : http.go
# @Description : 
*/
package utils

import (
	"context"
	"fmt"
	"github.com/valyala/fasthttp"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func ConvtMap2FastHttpArgs(params map[string]interface{}) *fasthttp.Args {
	args := &fasthttp.Args{}
	for key, value := range params {
		switch value.(type) {
		case string:
			args.Set(key, value.(string))
		case int:
			args.SetUint(key, value.(int))
		}
	}
	return args
}

// post表单的方式
func DoPostForm(ctx context.Context, postUrl string, keys []string, postValues []string) (*http.Response, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		values := url.Values{}
		l := len(keys)
		for i := 0; i < l; i++ {
			values.Add(keys[i], postValues[i])
		}
		return http.PostForm(postUrl, values)
	}
}

type StaticDynamicHtmlReq struct {
	// 模板文件路径
	TemplateFilePath string
	// 基础存储路径
	StaticHtmlBaseStorePath string
	// 新的html名称
	StaticHtmlName string
	// 对象
	Data map[string]interface{}
}
type StaticDynamicHtmlResp struct {
	// 静态页面的存储路径
	StaticHtmlStorePath string
}

// 动态生成静态页面
// 生成静态文件的方法
// filePath 为模板路径
// htmlOutPath 为持久化后的静态文件路径
func StaticDynamicHtml(req StaticDynamicHtmlReq) (StaticDynamicHtmlResp, error) {
	// 1.获取模版
	contenstTmp, err := template.ParseFiles(req.TemplateFilePath)
	if err != nil {
		return StaticDynamicHtmlResp{}, err
	}
	// 2.获取html生成路径

	// 4.生成静态文件
	return generateStaticHtml(contenstTmp, req, req.Data)
}

// 生成静态文件
func generateStaticHtml(template *template.Template, req StaticDynamicHtmlReq, data map[string]interface{}) (StaticDynamicHtmlResp, error) {
	var (
		result StaticDynamicHtmlResp
	)
	fileName := req.StaticHtmlBaseStorePath + string(filepath.Separator) + req.StaticHtmlName + ".html"
	// 1.判断静态文件是否存在
	if exist(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			fmt.Println("移除文件失败")
			return result, nil
		}
	}
	// 2.生成静态文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("打开文件失败")
	}
	defer file.Close()
	result.StaticHtmlStorePath = fileName
	return result, template.Execute(file, &data)
}

// 判断文件是否存在
func exist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}
