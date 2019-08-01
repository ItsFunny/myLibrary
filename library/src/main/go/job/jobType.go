/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-05-01 09:53
# @File : jobType.go
# @Description :
*/
package job

type JobTypeNode struct {
	Type JobType
	Next *JobTypeNode
}

type JobTypeLinkedList struct {
	Head *JobTypeNode
}

func (receiver *JobTypeLinkedList) Push(jobType JobType) {
}
