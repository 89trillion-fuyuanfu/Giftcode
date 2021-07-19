package router

import (
	"Giftcode/internal/ctrl"
	"github.com/gin-gonic/gin"
)

func Getroute() {
	// 路由配置
	router := gin.Default()
	port := "8080"
	//匹配calculate?name=xxx

	// 收到控制层传来的参数"ctrl.Getstring" 将其返回给浏览器前台页面。
	router.POST("/giftcode", ctrl.Getstring)

	router.Run(":" + port)

}
