package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ApiResponse struct {
	Response
	Data interface{}
}

func Success(c *gin.Context, data interface{}) {
	apiResponse := ApiResponse{
		Response{
			Code: 1,
			Msg:  "Success",
		},
		data,
	}
	c.JSON(http.StatusOK, apiResponse)
}

func Fail(c *gin.Context, msg string) {
	apiResponse := ApiResponse{
		Response{
			Code: 0,
			Msg:  msg,
		},
		make([]interface{}, 0),
	}
	c.AbortWithStatusJSON(http.StatusOK, apiResponse)
}
