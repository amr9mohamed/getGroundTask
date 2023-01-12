package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheckInitRoute(router *gin.Engine) {
	router.GET(
		"/ping", func(c *gin.Context) {
			c.JSON(
				http.StatusOK, gin.H{
					"message":  "pong",
					"success":  true,
					"response": "https://www.youtube.com/watch?v=t2NgsJrrAyM",
				},
			)
		},
	)
}
