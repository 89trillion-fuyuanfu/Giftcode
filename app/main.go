package main

import "Giftcode/app/http"

func main() {
	//使用简介：1， 后台管理员设置奖品方法为：service.Setstring()  需要设置奖品则运行service.Setstring()
	// 使用简介：在postman中输入url:http://localhost:8080/giftcode  key设置为giftcode 在value中输入礼品码即可获得礼品详细信息
	// jrTn5htC 9Xf9GAx5   这两个激活码的时间都设置为500小时，7月17日创建的，您也可以自己利用service2.Setstring()来创建：
	// 1，和redis数据库建立连接
	//service2.Connect()

	http.Start()

}
