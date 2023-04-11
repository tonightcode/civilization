package handlers

import (
	"zongheng/logic"
)

type Celebrity struct {
	Base
}

func (celebrity Celebrity) GetAll() ApiResponse {
	celebrity.Auth()
	_l := logic.Celebrity{}
	data := _l.GetAll()
	return Success(data)
}
