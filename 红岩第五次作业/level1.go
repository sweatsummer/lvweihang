// Lv0
// 复习上课所用到的代码 自己手敲一遍
// Lv1
// 通过cookie 实现用户的状态保持 通过登录接口实现登录 登录用户与未登录用户访问同一接口所获取的信息不同就行
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

var u user

func login(c *gin.Context) {

	fmt.Println("登录：")
	fmt.Println("请输入帐号:")
	fmt.Scan(&u.Username)
	fmt.Println("请输入密码:")
	fmt.Scan(&u.Password)
	if u.Username == "彭彭" && u.Password == "114514" {
		c.Next()
	} else {
		c.Abort()
	}
}

func handercookei(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		cookie = "null"
		c.SetCookie("gin_cookie",
			"val's gin",
			3600,
			"/",
			"localhost",
			false, false)
	}
	if cookie == "val's gin" {
		c.Next()
	}
}

func main() {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.GET("/info", login, handercookei, func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message":  "login success!",
				"username": u.Username,
				"password": u.Password,
			})
		})
		user.GET("/message", func(c *gin.Context) {
			cookie1, _ := c.Cookie("gin_cookie")
			if cookie1 != "val's gin" {
				c.Abort()
			}
		}, func(c *gin.Context) {
			c.JSON(200, gin.H{
				"secret": "这是个秘密",
			})
		})
	}
	r.Run()
}
