package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// 加载配置文件
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	// 遍历 web.go 中定义的所有 webRoutes
	for _, route := range webRoutes {
		switch route.Method {
		case "GET":
			router.GET(route.Pattern, route.HandleFunc)
		}
	}
	return router
}

// 记录请求日志信息中间件
func loggingRequestInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//打印url明细
		fmt.Printf("Request URL: %s\n", r.URL)
		fmt.Printf("User Agent: %s\n", r.Header.Get("User-Agent"))
		fmt.Printf("Request Header: %v\n", r.Header)
		next.ServeHTTP(w, r)
	})
}
