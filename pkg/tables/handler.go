package tables

import (
	"github.com/getground/tech-tasks/backend/definitions/tables"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) Create(c *gin.Context) (req tables.CreateRequest, err error) {
	err = c.ShouldBindJSON(&req)
	return
}
