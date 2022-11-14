package handler

import (
	"fmt"
	"github.com/Crush-s/ginProject/app/blog/edu"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func GoodsHandler(c *gin.Context) {
	// 声明接收的变量
	var login edu.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if login.User != "root" || login.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

// SomeJSON 1.json
func SomeJSON(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "someJSON", "status": 200})
}

// SomeStruct 2. 结构体响应
func SomeStruct(c *gin.Context) {
	var msg = struct {
		Name    string
		Message string
		Number  int
	}{
		Name:    "张三",
		Message: "成功",
		Number:  23,
	}
	c.JSON(http.StatusOK, msg)
}

// SomeXML 3.XML
func SomeXML(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"message": "someJSON", "status": 200})
}

// SomeYAML 4.YAML响应
func SomeYAML(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"message": "someJSON", "status": 200})
}

// SomeProtoBuf 5.protobuf格式,谷歌开发的高效存储读取的工具
// 数组？切片？如果自己构建一个传输格式，应该是什么格式？
func SomeProtoBuf(c *gin.Context) {
	reps := []int64{int64(1), int64(2)}
	// 定义数据
	label := "label"
	// 传protobuf格式数据
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	c.ProtoBuf(200, data)
}

// Index html
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})
}

// Ce 中间件
func Ce(c *gin.Context) {
	// 取值
	req, _ := c.Get("request")
	fmt.Println("request:", req)
	// 页面接收
	c.JSON(200, gin.H{"request": req})
}

func CheckoutHandler(c *gin.Context) {

}
