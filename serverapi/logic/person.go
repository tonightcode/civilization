package logic

import (
	"culture/entity"
	"culture/model"
	"time"
)

type Person struct {
}

var _mPerson model.Person

func (person Person) GetPerson(id int) interface{} {
	personEntity, _ := _mPerson.GetPerson(id)
	if personEntity.ID <= 0 {
		return make([]interface{}, 0)
	}
	return personEntity
}

func (person Person) GetPersons() interface{} {
	personEntitys, _ := _mPerson.GetPersons()
	return personEntitys
}

func (person Person) EditPerson(id int, param map[string]string) int {
	var personEntity entity.Person
	personEntity.Name = param["name"]
	personEntity.Desc = param["desc"]
	live_start, _ := time.Parse("2006-01-02 15:04:05", param["live_start"])
	live_end, _ := time.Parse("2006-01-02 15:04:05", param["live_end"])
	personEntity.Live_start = entity.LocalTime(live_start)
	personEntity.Live_end = entity.LocalTime(live_end)
	res := _mPerson.EditPerson(personEntity)
	return res
}
