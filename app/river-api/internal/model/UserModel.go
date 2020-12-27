package model

import (
	"fmt"
	"go-gin-web-demo/app/river-api/internal/dao"
)

type UserModel struct {
	Id       int32  `gorm:"column:id;" json:"id"`
	Email    string `gorm:"column:email;" json:"email"`
	Name     string `gorm:"column:name;" json:"name"`
	Password string `gorm:"column:password;" json:"password"`
}

func (UserModel) TableName() string {
	return "user"
}

func (user *UserModel) Insert() {
	fmt.Println(dao.DB)
	dao.DB.Create(user)
}
