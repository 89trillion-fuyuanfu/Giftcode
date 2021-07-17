package ctrl

import (
	"Giftcode/internal/service"
	"github.com/gin-gonic/gin"
)

func Getstring(c *gin.Context) {
	giftcode := c.PostForm("giftcode")
	//fmt.Println(giftcode)
	//c.String(http.StatusOK,"%s",service.Calculate(name))
	//c.JSON(200, service.Calculate(service.Convert(name)))
	//c.String(http.StatusOK,"%s",service.Getstring(giftcode))

	c.String(200, service.Getstring(giftcode))
}
