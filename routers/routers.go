package routers

import (
	"github.com/Crush-s/ginProject/app/handlers"
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options []Option

// Include 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// Init 初始化
func Init() *gin.Engine {
	r := gin.New()
	r.LoadHTMLGlob("tem/**/*")
	// 注册中间件
	r.Use(handlers.MiddleWare())
	for _, opt := range options {
		opt(r)
	}
	return r
}
