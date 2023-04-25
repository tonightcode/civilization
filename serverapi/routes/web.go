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
	Fn      gin.HandlerFunc
}

var c gin.Context

// 声明 WebRoutes 切片存放所有 Web 路由
type WebRoutes []WebRoute

// 定义所有 Web 路由
var webRoutes = WebRoutes{
	WebRoute{
		"getEvent",
		"GET",
		"/celebrities/getEvent",
		handlers.Event{}.GetEvent,
	},
	WebRoute{
		"getEvents",
		"GET",
		"/celebrities/getEvents",
		handlers.Event{}.GetEvents,
	},
	WebRoute{
		"editEvent",
		"GET",
		"/celebrities/editEvent",
		handlers.Event{}.EditEvent,
	},
}
