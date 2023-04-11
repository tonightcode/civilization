package logic

import (
	"culture/entity"
	"culture/model"
)

type Celebrity struct {
}

var _m model.Celebrity

func (celebrity Celebrity) GetEvent() entity.Event {
	event, _ := _m.GetEvent()
	return event
}
