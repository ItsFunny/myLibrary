/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-20 13:50 
# @File : select.go
# @Description : 
*/
package db


// 拼接sql语句
func SpliceInSQL(colunm string, idSlice []int) (string, []interface{}) {
	var (
		judgeSQL string
		params   []interface{}
	)
	idList := make([]int, 0)
	hashSet := make(map[int]struct{})
	for _, id := range idSlice {
		if _, exist := hashSet[id]; !exist && id > 0 {
			hashSet[id] = struct{}{}
			idList = append(idList, id)
		}
	}

	judgeSQL += " "
	l := len(idList)
	if l == 0 {
		return "", params
	} else if l == 1 {
		judgeSQL = colunm + "=?"
		params = append(params, idList[0])
	} else {
		judgeSQL += colunm + " IN (  "
		for i := 0; i < len(idList)-1; i++ {
			judgeSQL += "?,"
			params = append(params, idList[i])
		}
		judgeSQL += " ? )"
		params = append(params, idList[len(idList)-1])
	}
	return judgeSQL, params
}
