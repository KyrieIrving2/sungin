package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sungin/contract"
	"sungin/jwt"
)

// @description 获取token
// @tags auth
// @version 1.0
// @accept application/x-json-stream
// @Produce json
// @Success 200 "ok"
// @router /auth [get]
func GetAccessToken(context *gin.Context) {
	context.JSON(http.StatusOK,
		gin.H{
			"Token": "Bearer jiodsfjJopjop.joisjofjos.jisojdofjs",
		})
}

// @description 获取token
// @tags auth
// @version 1.0
// @accept application/x-json-stream
// @Produce json
// @Param LoginInputModel body contract.LoginInputModel true "LoginInputModel"
// @Success 200 "ok"
// @router /auth [post]
func AccessToken(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user contract.LoginInputModel
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	// 校验用户名和密码是否正确
	if user.Username == "q1mi" && user.Password == "q1mi123" {
		// 生成Token
		tokenString, _ := jwt.GenToken(user.Username)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return
}

