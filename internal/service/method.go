package service

import (
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"time"
)

// 定义一个*redis.Client类型的全局变量
var Client *redis.Client

//连接数据库
func Connect() {
	// 连接数据库
	// 传入一个指定Redis集群服务器信息的结构体类型的参数，返回一个redis集群的客户端
	Client = redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		PoolSize:     1000,
		ReadTimeout:  time.Millisecond * time.Duration(100),
		WriteTimeout: time.Millisecond * time.Duration(100),
		IdleTimeout:  time.Second * time.Duration(60),
	})
	_, err := Client.Ping().Result()
	if err != nil {
		panic("init redis error")
	} else {
		fmt.Println("init redis ok")
	}
}
func Setstring() {
	Connect()
	var Giftdescription string // 礼品描述
	var remain string          //剩余可领取次数
	var date int               //有效期
	var giftcontent string     // 礼品内容
	var name string            // 创建人姓名
	var createtime string
	fmt.Printf("请输入礼品描述")
	fmt.Scan(&Giftdescription)
	fmt.Printf("请输入可领取次数")
	fmt.Scan(&remain)
	fmt.Printf("请输入有效期（小时）")
	fmt.Scan(&date)
	fmt.Printf("请输入礼品内容")
	fmt.Scan(&giftcontent)
	fmt.Printf("请输入创建人的姓名")
	fmt.Scan(&name)
	fmt.Printf("请输入创建时间")
	fmt.Scan(&createtime)
	var date2 string
	date2 = string(date)

	a := RandomStr(8)
	err := Client.Set(a, "礼品描述："+Giftdescription+" 剩余可领取次数："+remain+" 有效期："+date2+" 礼品内容："+giftcontent+" 创建人姓名："+name+" 创建时间："+createtime, time.Duration(date)*time.Hour).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = Client.Set(a, "礼品描述："+Giftdescription+" 剩余可领取次数："+remain+" 有效期："+date2+" 礼品内容："+giftcontent+" 创建人姓名："+name+" 创建时间："+createtime, 0*time.Second).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("请记好您的礼品码" + a)

}

// 得到存储的字符串
func Getstring(giftcode string) string {
	Connect()
	result, err := Client.Get(giftcode).Result()
	if err != nil {
		fmt.Println(err)
		return "没有该激活码"
	}
	//fmt.Println(result)
	return result

}

//RandomStr 随机生成字符串
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 接口2：管理后台调用-查询礼品码信息
func Inquire(redisdb *redis.Client, giftcode string) {
	result, err := redisdb.Get(giftcode).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
