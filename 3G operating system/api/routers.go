package api

import (
	"3G_operating_system/util"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	u := r.Group("/user")
	{
		//注册用户//
		u.POST("/register", register)
		//登录//
		u.POST("/login", Login)
		//自我充值,使用中间件鉴权//
		u.POST("/recharge", util.MiddleWare(), recharge)
	}
}
