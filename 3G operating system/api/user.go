package api

import (
	"3G_operating_system/model"
	"3G_operating_system/service"
	"3G_operating_system/util"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
)

// register 注册判断//
func register(c *gin.Context) {
	UserName := c.PostForm("username")
	Password := c.PostForm("password")
	if (UserName == " " || Password == " ") || (UserName == "" || Password == "") {
		util.RespParamErr(c)
		return
	}
	u, err := service.IsExistUserByName(UserName)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RespParamErr(c)
		return
	}
	if u.UserName != "" {
		util.NormErr(c, 300, "账户已存在...")
		return
	}
	err = service.CreateUser(model.User{
		UserName: UserName,
		Password: Password,
	})
	if err != nil {
		util.RespParamErr(c)
		return
	}
	util.RespOk(c)
}

func Login(c *gin.Context) {
	UserName := c.PostForm("username")
	Password := c.PostForm("password")
	if (UserName == " " || Password == " ") || (UserName == "" || Password == "") {
		util.RespParamErr(c)
		return
	}
	u, err := service.IsExistUserByName(UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user err %v", err)
			util.ResInternalErr(c)
			return
		}
		return
	}
	if u.Password != Password {
		util.NormErr(c, 200, "密码错误")
		return
	}
	util.OkLogin(c)
	c.SetCookie("name", Password, 3600, "", "/", false, false)
}

func recharge(c *gin.Context) {
	Username := c.PostForm("username")
	Number := c.PostForm("pay")
	if Username == "" || Username == " " {
		util.RespParamErr(c)
		return
	}
	_, err := service.InputData(Number)
	if err != nil {
		util.RespParamErr(c)
		return
	}
	util.RespOk(c)
}
