package handlers

import (
	"culture/logic"
	"strconv"

	"github.com/gin-gonic/gin"
)

var _l logic.Event

type Event struct {
	Base
}

type eventParams struct {
	Id          string `form:"title"`
	Title       string `form:"title" json:"title" binding:"required"`
	Content     string `form:"content" json:"content" binding:"required"`
	Happened_at string `form:"happened_at" json:"happened_at" binding:"required"`
}

// get one
func (event Event) GetEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data := _l.GetEvent(id)
	Success(c, data)
}

// get all
func (event Event) GetEvents(c *gin.Context) {
	data := _l.GetEvents()
	Success(c, data)
}

// add or update
func (event Event) EditEvent(c *gin.Context) {
	title := c.Query("title")
	content := c.Query("content")
	happened_at := c.Query("happened_at")
	param := map[string]string{
		"title":       title,
		"content":     content,
		"happened_at": happened_at,
	}
	data := _l.EditEvent(param)
	Success(c, data)
}
