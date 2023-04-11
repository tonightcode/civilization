package handlers

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ApiResponse struct {
	Response
	Data struct {
		Total string `json:"total"`
		List  string `json:"list"`
	} `json:"data"`
}

func Success(data interface{}) ApiResponse {
	apiResponse := ApiResponse{}
	apiResponse.Code = 0
	apiResponse.Msg = "Success"
	return apiResponse
}
