package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Id    int `json:"id"`
	Email int `json:"email"`
}

type Users []User
