/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-02 10:46 
# @File : chinese.go
# @Description : 
*/
package sort

import "xinglan.com/core-system/platform/platform-rpc/common/utils"

var (
	DEFAULT_SORTED_BY_CHINESE = func(strI, strJ string, desc bool) bool {
		hashI, _ := utils.UTF82GBK(strI)
		hashJ, _ := utils.UTF82GBK(strJ)
		bLen := len(hashI)
		for idx, chr := range hashI {
			if idx > bLen-1 {
				if desc {
					return false
				} else {
					return true
				}

			}
			if chr != hashJ[idx] {
				if desc {
					return chr < hashJ[idx]
				} else {
					return chr > hashJ[idx]
				}
			}
		}
		return true
	}

	DEFAULT_SORTED_BY_INT = func(int1, int2 int, desc bool) bool {
		if desc {
			return int1 < int2
		} else {
			return int1 > int2
		}
	}
)
