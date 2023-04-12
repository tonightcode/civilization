package handlers

import (
	"culture/logic"
	"culture/model"
)

var _l logic.Celebrity
var _m model.Celebrity

type Celebrity struct {
	Base
}

func (celebrity Celebrity) GetEvent() ApiResponse {
	celebrity.Auth()
	data := _l.GetEvent()
	return Success(data)
}
