package logic

import (
	"culture/entity"
	"culture/model"
	"time"
)

type Event struct {
}

var _m model.Event

func (event Event) GetEvent(id int) interface{} {
	eventEntity, _ := _m.GetEvent(id)
	if eventEntity.ID <= 0 {
		return make([]interface{}, 0)
	}
	return eventEntity
}

func (event Event) GetEvents() interface{} {
	eventEntitys, _ := _m.GetEvents()
	return eventEntitys
}

func (event Event) EditEvent(param map[string]string) int {
	var eventEntity entity.Event
	eventEntity.Title = param["title"]
	eventEntity.Content = param["content"]

	happenedAt, _ := time.Parse("2006-01-02 15:04:05", param["happened_at"])
	eventEntity.HappenedAt = entity.LocalTime(happenedAt)
	res := _m.EditEvent(eventEntity)
	return res
}
