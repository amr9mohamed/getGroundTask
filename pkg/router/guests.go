package router

import (
	"github.com/getground/tech-tasks/backend/pkg/modules/guests"
	"github.com/gin-gonic/gin"
)

func GuestsInitRoute(router *gin.Engine, ctrl guests.Controller) {
	router.POST("/guest_list/:name", ctrl.Create)
	router.GET("/guest_list", ctrl.GetGuestList)
	router.GET("/guests", ctrl.GetGuests)
}
