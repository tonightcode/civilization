package routes

import (
	"culture/handlers"

	"github.com/gin-gonic/gin"
)

// 定义一个 WebRoute 结构体用于存放单个路由
type WebRoute struct {
	Name    string
	Method  string
	Pattern string
	Fn      handlers.ApiResponse
}

var c gin.Context

// 声明 WebRoutes 切片存放所有 Web 路由
type WebRoutes []WebRoute

// 定义所有 Web 路由
var webRoutes = WebRoutes{
	WebRoute{
		"Home",
		"GET",
		"/celebrities/getEvent",
		handlers.Celebrity{}.GetEvent(),
	},
}
