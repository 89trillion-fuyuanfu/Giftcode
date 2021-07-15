
# 第三题礼品码技术文档

### 1，整体框架

礼品码，主要思路是创建map类型数据来描述礼品码和礼品描述。使用key来描述礼品码，使用value来作为礼品内容，value为struct类型。

### 2，目录结构

![image-20210715170147351](/Users/fuyuanfu/Library/Application Support/typora-user-images/image-20210715170147351.png)

Main.go  --代码入口

Connection.go -- 数据库连接文件

Returnroute.go --得到网页路由的参数后，在网页进行渲染。

Interface.go --三个接口文件

Getroute.go--接受网页的url

Method.go--实现三个接口的方法

README.md--项目技术文档

### 3，代码逻辑分层

| 层        | 文件夹                          | 主要职责                       | 调用关系                     | 其他说明         |
| --------- | ------------------------------- | ------------------------------ | ---------------------------- | ---------------- |
| 应用层    | App/main.go                     | 进程启动                       | 调用路由层                   | 不可同层互相调用 |
| 控制层    | ctrl/Returnroute.go             | 返回路由信息，将数据渲染在网页 | 被路由层调用                 | 不可同层互相调用 |
| 模型层    | Model/connection.goInterface.go | 连接数据库和实现三个接口       | 被service调用                | 不可同层互相调用 |
| service层 | Service.method.go               | 实现三个接口中的方法           | 可以调用模型层，被控制层调用 | 不可同层互相调用 |

### 4.存储设计

创建map类型来存储礼品码和礼品内容。礼品码为key，礼品内容为value，value为结构体struct类型，礼品内容为切片内容，key为string类型。

### 5.接口设计

请求方法：http post

接口地址：https://localhost:8080/getroute

请求参数：

礼品码：string类型 字符串

请求响应：返回用户的奖励内容，一个字符数组

响应状态码：服务端未启动

|          | 礼品码key | 值value | 次数remain | 有效期date | 礼品内容gift |
| -------- | --------- | ------- | ---------- | ---------- | ------------ |
| 数据类型 | string    | struct  | int        | int        | String[]     |



### 6，第三方库

"github.com/garyburd/redigo/redis"  导入redis数据库相对应的方法，使用redis数据库。

### 7.编译执行

在app包里运行main.go文件

### 

### 8，逻辑图

<img width="847" alt="第三题流程图" src="https://user-images.githubusercontent.com/87068277/125781793-d3d8b2d7-a3c1-4d14-b5c3-8b522ccd7707.png">

