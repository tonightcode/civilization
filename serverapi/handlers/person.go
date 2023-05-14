package handlers

import (
	"culture/logic"
	"strconv"

	"github.com/gin-gonic/gin"
)

var _lPerson logic.Person

type Person struct {
	Base
}

// get one
func (person Person) GetPerson(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data := _lPerson.GetPerson(id)
	Success(c, data)
}

// get all
func (person Person) GetPersons(c *gin.Context) {
	data := _lPerson.GetPersons()
	Success(c, data)
}

// add or update
func (person Person) EditPerson(c *gin.Context) {
	name := c.Query("name")
	desc := c.Query("desc")
	live_start := c.Query("live_start")
	live_end := c.Query("live_end")
	id, _ := strconv.Atoi(c.Query("id"))
	param := map[string]string{
		"name":       name,
		"desc":       desc,
		"live_start": live_start,
		"live_end":   live_end,
	}
	data := _lPerson.EditPerson(id, param)
	Success(c, data)
}
