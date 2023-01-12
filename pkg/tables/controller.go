package tables

import (
	"github.com/getground/tech-tasks/backend/definitions/tables"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Controller struct {
	handler Handler
	service tables.Service
}

func NewController(handler Handler, service tables.Service) Controller {
	return Controller{
		handler: handler,
		service: service,
	}
}

func (ctrl Controller) Create(c *gin.Context) {
	req, err := ctrl.handler.Create(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	res, err := ctrl.service.Create(req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (ctrl Controller) CountEmptySeats(c *gin.Context) {
	//	call service
}
