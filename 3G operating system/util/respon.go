package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type resTemplate struct {
	status int    `json:"status"`
	Infor  string `json:"infor"`
}

// register//
var Ok = resTemplate{
	status: 200,
	Infor:  "register successfully",
}
var OKlogin = resTemplate{
	status: 200,
	Infor:  "login successfully",
}

var paramError = resTemplate{
	status: 300,
	Infor:  "param error",
}

// 连接报错//
var Internalerr = resTemplate{
	status: 500,
	Infor:  "internal err",
}

func RespOk(c *gin.Context) {
	c.JSON(http.StatusOK, Ok)
}
func OkLogin(c *gin.Context) {
	c.JSON(http.StatusOK, OKlogin)
}

func RespParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, paramError)
}

func ResInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Internalerr)

}

func NormErr(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
}
