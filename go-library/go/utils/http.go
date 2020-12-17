/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-26 14:13 
# @File : http.go
# @Description : 
*/
package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
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
	NewName string
	Suffix  string
	// 对象
	Data map[string]interface{}
}
type StaticDynamicHtmlResp struct {
	// 静态页面的存储路径
	StaticHtmlStorePath string
}

func NewStaticDynamicHtmlReq(templatePath, staticBaseStorePath, newName string, data map[string]interface{}) StaticDynamicHtmlReq {
	r := StaticDynamicHtmlReq{
		TemplateFilePath:        templatePath,
		StaticHtmlBaseStorePath: staticBaseStorePath,
		NewName:                 newName,
		Data:                    data,
	}
	return r
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
	fileName := req.StaticHtmlBaseStorePath + string(filepath.Separator)

	if !IsFileOrDirExists(fileName) {
		if err := CreateMultiFileDirs(fileName); nil != err {
			return result, errors.New("创建文件夹失败:" + err.Error())
		}
	}
	if len(req.Suffix) == 0 {
		req.Suffix = ".html"
	}
	fileName += req.NewName + req.Suffix
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
		return result, errors.New("打开文件失败:" + err.Error())
	}
	defer file.Close()
	result.StaticHtmlStorePath = fileName
	return result, template.Execute(file, &data)
}

type HttpResult struct {
	Data interface{} `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func BuildHttpRequest(url string, data []byte, heads map[string]string) (*http.Request, error) {
	request, e := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if nil != e {
		return nil, errors.New("创建httpRequest失败")
	}
	for k, v := range heads {
		request.Header.Set(k, v)
	}
	return request, nil

}

type HttpReq struct {
	TimeOutSeconds int
	MaxFailCount   int
	HttpUrl        string
	BytesData      []byte
	Headers        map[string]string
	// res,bool: 是否停止重试,error: 错误信息
	HandlerResponse func(response *http.Response) (*HttpResult, bool, error)
}

func PostPerClient(req HttpReq) (interface{}, error) {
	result, e := httpForAes(req)
	if nil != e {
		return "", e
	}
	return result.Data, nil
}

func httpForAes(req HttpReq) (*HttpResult, error) {

	client := &http.Client{}
	defer client.CloseIdleConnections()

	// failCount := 0
	// var lastErr error
	// for failCount <= req.MaxFailCount {
	// 	httpReq, e := BuildHttpRequest(req.HttpUrl, req.BytesData, req.Headers)
	// 	if nil != e {
	// 		panic("构建http请求参数失败:" + e.Error())
	// 	}
	// 	// httpReq.Header.Set("Content-Type", "application/json;charset=UTF-8")
	// 	httpResp, e := client.Do(httpReq)
	// 	if nil != e {
	// 		failCount++
	// 		log.Println("http调用失败:" + e.Error())
	// 		lastErr = e
	// 		continue
	// 	}
	// 	if res, keepRetry, e := req.HandlerResponse(httpResp); nil != e {
	// 		log.Println("http调用失败,状态为:" + httpResp.Status + ",错误信息为:"+e.Error())
	// 		if keepRetry {
	// 			failCount++
	// 			log.Println("开始失败重试,failCount=" + strconv.Itoa(failCount))
	// 			continue
	// 		} else {
	// 			log.Println("异常停止,停止失败重试")
	// 			lastErr = e
	// 			break
	// 		}
	// 	} else{
	// 		return res,nil
	// 	}
	// }
	// return nil,lastErr

	resultChan := make(chan *HttpResult)
	errChan := make(chan error)
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*time.Duration(req.TimeOutSeconds))

	go func() {
		defer func() {
			if ePanic := recover(); nil != ePanic {
				if err, ok := ePanic.(error); ok {
					errChan <- errors.New("http时发生了panic,错误为:" + err.Error())
				} else {
					marshal, _ := json.Marshal(ePanic)
					if len(marshal) == 0 {
						marshal = []byte("unknown")
					}
					errChan <- errors.New("未知panic:" + string(marshal))
				}
			}
		}()
		failCount := 0
		var lastErr error
		for failCount <= req.MaxFailCount {
			httpReq, e := BuildHttpRequest(req.HttpUrl, req.BytesData, req.Headers)
			if nil != e {
				panic("构建http请求参数失败:" + e.Error())
			}
			// httpReq.Header.Set("Content-Type", "application/json;charset=UTF-8")
			httpResp, e := client.Do(httpReq)
			if nil != e {
				failCount++
				log.Println("http调用失败:" + e.Error())
				lastErr = e
				continue
			}
			if res, keepRetry, e := req.HandlerResponse(httpResp); nil != e {
				log.Println("http调用失败,状态为:" + httpResp.Status + ",错误信息为:")
				if keepRetry {
					failCount++
					log.Println("开始失败重试,failCount=" + strconv.Itoa(failCount))
					continue
				} else {
					log.Println("异常停止,停止失败重试")
					lastErr = e
					break
				}
			} else {
				resultChan <- res
			}
		}
		errChan <- lastErr
	}()

	select {
	case res := <-resultChan:
		cancelFunc()
		return res, nil
	case err := <-errChan:
		return nil, err
	case <-timeout.Done():
		return nil, errors.New("http请求达到超时也无法获取到aes密钥,检查Java服务器是否启动着")
	}
}
