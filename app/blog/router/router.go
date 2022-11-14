package router

import (
	"github.com/Crush-s/ginProject/app/blog/handler"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.POST("/123", handler.GoodsHandler)
	e.GET("/234", handler.CheckoutHandler)
	e.GET("/someJSON", handler.SomeJSON)
	e.GET("/someStruct", handler.SomeStruct)
	e.GET("/someXML", handler.SomeXML)
	e.GET("/someYAML", handler.SomeYAML)
	e.GET("/someProtoBuf", handler.SomeProtoBuf)
	e.GET("/index", handler.Index)
	e.GET("/ce", handler.Ce)
}
