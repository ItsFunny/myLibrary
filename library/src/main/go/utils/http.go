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
	"net/http"
	"net/url"
)

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
