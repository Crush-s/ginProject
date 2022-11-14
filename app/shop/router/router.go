package router

import (
	"github.com/Crush-s/ginProject/app/shop/handler"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("345", handler.PostHandler)
	e.GET("456", handler.CommentHandler)
}
