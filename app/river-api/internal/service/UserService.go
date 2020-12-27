package service

import (
	"fmt"
	"github.com/jinzhu/copier"
	"go-gin-web-demo/app/river-api/dto"
	"go-gin-web-demo/app/river-api/internal/model"
)

type UserServiceImpl struct {
}

var UserService *UserServiceImpl

func (userService *UserServiceImpl) Create(userDTO *dto.UserCreateDTO) model.UserModel {
	var userModel model.UserModel
	err := copier.Copy(&userModel, &userDTO)
	if err != nil {
		fmt.Println("Should not raise error")
	}
	userModel.Insert()
	return userModel
}
