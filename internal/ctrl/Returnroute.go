package ctrl

import (
	"Giftcode/internal/service"
	"github.com/gin-gonic/gin"
)

func Getstring(c *gin.Context) {
	// 将post表单请求的值赋值给"giftcode"变量
	giftcode := c.PostForm("giftcode")

	//c.String(http.StatusOK,"%s",service.Calculate(name))
	//c.JSON(200, service.Calculate(service.Convert(name)))
	//c.String(http.StatusOK,"%s",service.Getstring(giftcode))

	// 调用service层的getstring方法来处理网页传入的"giftcode"参数，然后将数据传递给路由层（router）
	c.String(200, service.Getstring(giftcode))
}
