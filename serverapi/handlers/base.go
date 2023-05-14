package handlers

import (
	"github.com/gin-gonic/gin"
)

type Base struct {
}

func (base Base) Auth() string {
	return "success"
}

func (base Base) GetParam(c *gin.Context, param string) string {
	if param == "" {
		Fail(c, param+"不能为空")
		return ""
	}
	return c.Query(param)
}
