package guests

import (
	"errors"
	"github.com/getground/tech-tasks/backend/definitions/guests"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}
func (h Handler) Create(c *gin.Context) (req guests.CreateRequest, err error) {
	name := c.Param("name")
	if name == "" {
		err = errors.New("name parameter is not sent")
		return
	}
	req.Name = name
	err = c.ShouldBindJSON(&req)
	return
}

func (h Handler) CheckIn(c *gin.Context) (req guests.CheckInRequest, err error) {
	name := c.Param("name")
	if name == "" {
		err = errors.New("name is required")
		return
	}
	req.Name = name
	err = c.ShouldBindJSON(&req)
	return
}

func (h Handler) CheckOut(c *gin.Context) (name string, err error) {
	name = c.Param("name")
	if name == "" {
		err = errors.New("name is required")
	}
	return
}
