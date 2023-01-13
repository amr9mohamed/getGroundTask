package boot

import (
	"github.com/getground/tech-tasks/backend/config"
	"github.com/getground/tech-tasks/backend/pkg/database"
	"github.com/getground/tech-tasks/backend/pkg/modules/guests"
	"github.com/getground/tech-tasks/backend/pkg/modules/tables"
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

	// inti handlers
	tablesHdl := tables.NewHandler()
	guestsHdl := guests.NewHandler()

	// init repositories
	tablesRepo := tables.NewRepository(dbConn)
	guestsRepo := guests.NewRepository(dbConn)

	// init services
	tablesSrv := tables.NewService(tablesRepo)
	guestsSrv := guests.NewService(guestsRepo, tablesSrv)

	// init controllers
	tablesCtrl := tables.NewController(tablesHdl, tablesSrv)
	guestsCtrl := guests.NewController(guestsHdl, guestsSrv)

	// init routers
	router.HealthCheckInitRoute(engine)
	router.TablesInitRouter(engine, tablesCtrl)
	router.GuestsInitRoute(engine, guestsCtrl)

	return engine
}
