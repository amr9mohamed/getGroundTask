package router

import (
	"github.com/getground/tech-tasks/backend/pkg/tables"
	"github.com/gin-gonic/gin"
)

func TablesInitRouter(router *gin.Engine, ctrl tables.Controller) {
	router.POST("/tables", ctrl.Create)
	router.GET("/seats_empty", ctrl.CountEmptySeats)
}
