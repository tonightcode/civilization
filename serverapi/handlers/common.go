package handlers

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

func Fail(msg string) ApiResponse {
	apiResponse := ApiResponse{
		Response{
			Code: 1,
			Msg:  "Success",
		},
		make([]interface{}, 0),
	}

	return apiResponse
}
