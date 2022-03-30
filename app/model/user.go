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
	OpenId   string
	Source   string
	NickName string
	Avatar   string
	Age      int
	Gender   int
}

type UserTokenResp struct {
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}

type UserInfo struct {
	UserId   uint
	NickName string
	Avatar   string
	Age      int
	Gender   int
}
