package main

import (
	"fmt"
	"github.com/Crush-s/ginProject/app/blog/router"
	router2 "github.com/Crush-s/ginProject/app/shop/router"
	"github.com/Crush-s/ginProject/routers"
)

func main() {
	// 加载多个APP的路由配置
	routers.Include(router2.Routers, router.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}

func ui(a, b int) int {
	c := 8
	return c*a + 2*b
}
