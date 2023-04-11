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

func Success(data map[string]string) ApiResponse {
	apiResponse := ApiResponse{}
	apiResponse.Code = 0
	apiResponse.Msg = "Success"
	apiResponse.Data.Total = data["total"]
	apiResponse.Data.List = data["list"]
	return apiResponse
}
