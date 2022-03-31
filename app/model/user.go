package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	OpenId   string `gorm:"not null; uniqueIndex:open_id_source_uniq,priority:1"` // 第三方登录openId
	Source   string `gorm:"not null; uniqueIndex:open_id_source_uniq,priority:2"` // 哪个平台登录
	NickName string `gorm:"not null"`
	Avatar   string
	Age      int
	Gender   int // 0=unknown 1=male 2=female
}

type CreateUserTokenReq struct {
	OpenId   string `json:"openId" binding:"required"`
	Source   string `json:"source" binding:"required"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Age      int    `json:"age"`
	Gender   int    `json:"gender"`
}

type UserTokenResp struct {
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}

type UserInfo struct {
	UserId   uint   `json:"userId"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Age      int    `json:"age"`
	Gender   int    `json:"gender"`
}
