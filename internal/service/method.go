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

// Setstring()函数用于将数据存入redis数据库

func Setstring() {
	Connect()                  // 先调用连接函数将redis和go连接
	var Giftdescription string // 礼品描述
	var remain string          //剩余可领取次数
	var date int               //有效期
	var giftcontent string     // 礼品内容
	var name string            // 创建人姓名
	var createtime string      // 创建时间
	fmt.Printf("请输入礼品描述")
	fmt.Scan(&Giftdescription) // 获取管理员的礼品描述
	fmt.Printf("请输入可领取次数")
	fmt.Scan(&remain) // 获取领取次数
	fmt.Printf("请输入有效期（小时）")
	fmt.Scan(&date) // 获取礼品内容
	fmt.Printf("请输入礼品内容")
	fmt.Scan(&giftcontent) // 获取创建人姓名
	fmt.Printf("请输入创建人的姓名")
	fmt.Scan(&name) // 获取创建人姓名
	fmt.Printf("请输入创建时间")
	fmt.Scan(&createtime) // 获取创建时间
	var date2 string
	date2 = string(date) // 将date 从int类型转换为string 类型。

	// Randomstr()生成一个随机码  生成一个随机码赋值给a变量
	a := RandomStr(8)
	// 将生成的随机码a做为这个数据的key  和其他礼品描述，领取次数等一起存入redis
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
	// 将礼品码输出给客户
	fmt.Println("请记好您的礼品码" + a)

}

// // 接口2：管理后台调用-查询礼品码信息
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
	// 在str串中随机取值
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
