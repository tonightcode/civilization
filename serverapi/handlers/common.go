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

func Success(data interface{}) ApiResponse {
	apiResponse := ApiResponse{
		Response{
			Code: 1,
			Msg:  "Success",
		},
		data,
	}
	return apiResponse
}

func Fail() {
	response := Response{
		Code: 0,
		Msg:  "Error",
	}

	var ctx gin.Context

	ctx.JSON(http.StatusOK, response)
}
