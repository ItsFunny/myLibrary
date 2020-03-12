/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-14 10:08 
# @File : dbBaseModel.go
# @Description : 
*/
package db

import "time"

type DBBaseModel struct {
	ID            int    `gorm:"column:ID" json:"id"`
	CreatedDate   time.Time  `gorm:"column:CREATE_DATE" json:"created_date"`
	CreatedUser   string `gorm:"column:CREATE_USER" json:"created_user"`
	CreatedUserId int    `gorm:"column:CREATE_USER_ID" json:"created_user_id"`

	UpdatedDate   time.Time  `gorm:"column:UPDATE_DATE" json:"updated_date"`
	UpdatedUser   string `gorm:"column:UPDATE_USER" json:"updated_user"`
	UpdatedUserId int    `gorm:"column:UPDATE_USER_ID" json:"updated_user_id"`
}
