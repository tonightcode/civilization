package handlers

import (
	"culture/logic"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var _l logic.Event

type Event struct {
	Base
}

// get one
func (event Event) GetEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data := _l.GetEvent(id)
	c.JSON(http.StatusOK, Success(data))
}

// get all
func (event Event) GetEvents(c *gin.Context) {
	data := _l.GetEvents()
	c.JSON(http.StatusOK, Success(data))
}

// add or update
func (event Event) EditEvent(c *gin.Context) {
	title := c.Query("title")
	content := c.Query("content")
	happened_at := c.Query("happened_at")
	id, _ := strconv.Atoi(c.Query("id"))
	param := map[string]string{
		"title":       title,
		"content":     content,
		"happened_at": happened_at,
	}
	data := _l.EditEvent(id, param)
	c.JSON(http.StatusOK, Success(data))
}
