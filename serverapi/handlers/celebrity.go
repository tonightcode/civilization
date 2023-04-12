package handlers

import (
	"culture/logic"
)

var _l logic.Celebrity

type Celebrity struct {
	Base
}

func (celebrity Celebrity) GetEvent() ApiResponse {
	celebrity.Auth()
	data := _l.GetEvent()
	return Success(data)
}
