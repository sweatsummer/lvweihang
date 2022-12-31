package service

import (
	"3G_operating_system/dao"
	"3G_operating_system/model"
)

// IsExistUserByName 根据用户名查找//
func IsExistUserByName(name string) (u model.User, err error) {
	u, err = dao.SearchUserByUserName(name)
	return
}

// InputData 输入业绩//
func InputData(number string) (u model.User, err error) {
	err = dao.InputRecharge(number)
	return
}

// CreateUser 创建用户//
func CreateUser(u model.User) error {
	err := dao.InsertUser(u)
	return err
}
