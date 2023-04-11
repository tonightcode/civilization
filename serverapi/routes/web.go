package routes

import (
	"culture/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义一个 WebRoute 结构体用于存放单个路由
type WebRoute struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc gin.HandlerFunc
}

// 声明 WebRoutes 切片存放所有 Web 路由
type WebRoutes []WebRoute

// 定义所有 Web 路由
var webRoutes = WebRoutes{
	WebRoute{
		"Home",
		"GET",
		"/celebrities/getEvent",
		func(ctx *gin.Context) {
			// defer func() {
			// 	if err := recover(); err != nil {
			// 		ctx.JSON(http.StatusOK, "Error")
			// 	}
			// }()
			ctx.JSON(http.StatusOK, handlers.Celebrity{}.GetEvent())
		},
	},
}
