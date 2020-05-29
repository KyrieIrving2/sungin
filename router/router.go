package router

import (
	"github.com/gin-gonic/gin"
	"sungin/api"
	"sungin/jwt"
)

func InitRouter(r *gin.Engine) {
	baseGroup := r.Group("api/v1/")

	//需要鉴权路由组
	baseAuthGroup:=baseGroup.Group("",jwt.JWTAuthMiddleware())

	userGroup := baseAuthGroup.Group("user")
	{
		userGroup.GET("", api.GetUserInfo)
	}

	authGroup := baseGroup.Group("auth")
	{
		authGroup.GET("", api.GetAccessToken)
		authGroup.POST("", api.AccessToken)
	}
}

