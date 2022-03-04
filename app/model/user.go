package model

import (
	"time"
)

type User struct {
	Id        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Salt      string     `json:"salt"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type UserRegisterCommand struct {
	Name       string `form:"name" json:"name" binding:"gte=1,lte=20"`
	Email      string `form:"email" json:"email" binding:"required,email"`
	Password   string `form:"password" json:"password" binding:"required,gte=6"`
	RePassword string `form:"re_password" json:"re_password" binding:"eqfield=Password"`
}

type UserLoginCommand struct {
	Email    string `form:"name" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,gte=6"`
}
