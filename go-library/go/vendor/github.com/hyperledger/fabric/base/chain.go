/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-12-11 13:18
# @File : chain.go
# @Description :
# @Attention :
*/
package base

import (
	"errors"
	"reflect"
)

var (
	STORE_ITERATOR_ERROR = errors.New("STOP_ITERATOR")
)

type ListHook func(node ILinkedList) error
type ListCreator func() ILinkedList
type ListCopy func(from, to ILinkedList) (ILinkedList, error)

type ILinkedList interface {
	GetNext() ILinkedList
	SetNext(linkInterface ILinkedList)
}

func NewBaseAlgorithmLinkedList(alg BaseType) *BaseAlgorithmLinkedList {
	return &BaseAlgorithmLinkedList{
		AlgorithmType: alg,
	}
}

type BaseAlgorithmLinkedList struct {
	AlgorithmType BaseType
	Next          ILinkedList
}

// func (b *BaseAlgorithmLinkedList) ToJson() ([]byte,error) {
// 	value, err := helper.ToValue(b)
// 	if nil!=err{
// 		return nil,err
// 	}
// 	return json.Marshal(value)
// }

func (b *BaseAlgorithmLinkedList) GetNext() ILinkedList {
	return b.Next
}

func (b *BaseAlgorithmLinkedList) SetNext(linkInterface ILinkedList) {
	b.Next = linkInterface
}

func IteratorFill(from, to ILinkedList, creator ListCreator, hookCopy ListCopy) (ILinkedList, error) {
	dumy := to
	for tmp := from; !IsNil(tmp); tmp = tmp.GetNext() {
		if IsNil(to) {
			to = creator()
			dumy = to
		}
		copy, e := hookCopy(tmp, dumy)
		if nil != e {
			return nil, e
		}
		dumy = copy
		if IsNil(tmp.GetNext()) {
			break
		}
		dumy.SetNext(creator())
		dumy = dumy.GetNext()
	}

	return to, nil
}

func IteratorFillWithDefault(from ILinkedList, creator ListCreator, hookCopy ListCopy) (ILinkedList, error) {
	to := creator()
	dumy := to
	for tmp := from; !IsNil(tmp); tmp = tmp.GetNext() {
		if IsNil(to) {
			to = creator()
			dumy = to
		}
		copy, e := hookCopy(tmp, dumy)
		if nil != e {
			return nil, e
		}
		dumy = copy
		if IsNil(tmp.GetNext()) {
			break
		}
		dumy.SetNext(creator())
		dumy = dumy.GetNext()
	}

	return to, nil
}

func Iterator(list ILinkedList, hook ListHook) error {
	for tmp := list; !IsNil(tmp); tmp = tmp.GetNext() {
		if err := hook(tmp); nil != err {
			if err == STORE_ITERATOR_ERROR {
				break
			}
			return err
		}
	}
	return nil
}

func LinkLast(firer, new ILinkedList) ILinkedList {
	if IsNil(firer) {
		return new
	}

	temp := firer
	for !IsNil(temp.GetNext()) {
		temp = temp.GetNext()
	}
	temp.SetNext(new)
	return firer
}
func IsNil(firer interface{}) bool {
	return firer == nil || (reflect.ValueOf(firer).Kind() == reflect.Ptr && reflect.ValueOf(firer).IsNil())
}
