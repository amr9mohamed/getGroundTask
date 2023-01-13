package guests

import (
	"github.com/getground/tech-tasks/backend/definitions/guests"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Controller struct {
	handler Handler
	service guests.Service
}

func NewController(handler Handler, service guests.Service) Controller {
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
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ctrl Controller) GetGuestList(c *gin.Context) {
	res, err := ctrl.service.GetGuestList()
	if err != nil {
		log.Error(err)
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"error": err,
			},
		)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ctrl Controller) GetGuests(c *gin.Context) {
	res, err := ctrl.service.GetGuests()
	if err != nil {
		log.Error(err)
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"error": err,
			},
		)
		return
	}

	c.JSON(http.StatusOK, res)
}
