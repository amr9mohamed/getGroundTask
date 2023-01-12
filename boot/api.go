package boot

import (
	"github.com/getground/tech-tasks/backend/config"
	"github.com/getground/tech-tasks/backend/pkg/database"
	"github.com/getground/tech-tasks/backend/pkg/router"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func API(cfg config.API) *gin.Engine {
	dbConn, err := database.New(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dbConn)
	engine := gin.New()
	engine.Use(
		gin.LoggerWithWriter(
			gin.DefaultWriter, "/ping",
		),
		gin.Recovery(),
	)

	// init repositories
	// init services
	// init controllers
	// init routers
	router.HealthCheckInitRoute(engine)
	return engine
}
