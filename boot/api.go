package boot

import (
	"github.com/getground/tech-tasks/backend/config"
	"github.com/getground/tech-tasks/backend/pkg/database"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func API(cfg config.API) *gin.Engine {
	dbConn, err := database.New(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	db, err := dbConn.DB()
	err = db.Ping()
	if err != nil {
		log.Fatalf("error pinging ")
	}
	engine := gin.New()
	return engine
}
