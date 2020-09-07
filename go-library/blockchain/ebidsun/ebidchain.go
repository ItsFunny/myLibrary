/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-05 19:20 
# @File : ebidchain.go
# @Description : 
# @Attention : 
*/
package main

import (
	"myLibrary/go-library/blockchain"
)

// func main() {
// 	t, e := template.ParseFiles("/Users/joker/go/src/myLibrary/go-library/blockchain/ebidsun/test.yaml")
// 	if nil != e {
// 		panic(e)
// 	}
//
// 	// if e := utils.CreateMultiFileDirs("/Users/joker/go/src/myLibrary/go-library/blockchain/ebidsun/test2.yaml"); nil != e {
// 	// 	panic(e)
// 	// }
// 	file, e := os.OpenFile("/Users/joker/go/src/myLibrary/go-library/blockchain/ebidsun/test.yaml",os.O_CREATE|os.O_WRONLY, os.ModePerm)
// 	if nil != e {
// 		panic(e)
// 	}
//
// 	if e = t.Execute(file, &map[string]string{"test": "joker"}); nil != e {
// 		panic(e)
// 	}
//
// }
var (
	eConfiguraiton EBidChainConfiguration
)

func init() {
	eConfiguraiton.BlockChainConfiguration = config.NewBlockChainConfiguration()
}

type EBidChainConfiguration struct {
	*config.BlockChainConfiguration
}

func main() {
	// str := os.Environ()
	// m:=make(map[string]string)
	// for _, s := range str {
	// 	fmt.Println(s)
	// }
	path := "/Users/joker/go/src/myLibrary/go-library/blockchain/ebidsun/ebidsun-application-blockchain-local-test.json"
	// bytes, e := ioutil.ReadFile("/Users/joker/go/src/myLibrary/go-library/blockchain/ebidsun/ebidsun-application-blockchain-local-test.json")
	// if nil != e {
	// 	panic(e)
	// }
	// if e = json.Unmarshal(bytes, &eConfiguraiton); nil != e {
	// 	panic(e)
	// }
	// fmt.Println(eConfiguraiton)
	e := eConfiguraiton.Config(path, config.ConfigWrapper{})
	if nil != e {
		panic(e)
	}

}
