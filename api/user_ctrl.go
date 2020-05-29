package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

// @Summary 获取人员信息
// @Description 获取人员信息
// @Tags user
// @Accept json
// @Produce json
// @param Authorization header string true " "
// @Success 200 "ok"
// @Failure 400 "fail"
// @Router /user [get]
func GetUserInfo(context *gin.Context) {
	name, _ := context.Get("userName")

	port:= viper.Get("port")
	var ss = viper.Get("jwt.secret")
	log.Print(port)
	log.Print(ss)

	context.JSON(http.StatusOK,
		gin.H{
			"message": fmt.Sprintf("my name is %s", name),
		})
}
