package database

import (
	"database/sql"
	"fmt"
	"github.com/getground/tech-tasks/backend/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(cfg config.Database) (gormDB *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port,
		cfg.Name,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	gormDB, err = gorm.Open(
		mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}),
		&gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Info),
			SkipDefaultTransaction: true,
		},
	)
	return gormDB.Session(&gorm.Session{}), err
}

func getCfg() gorm.Config {
	return gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
}

func NewDatabaseForTests(conn gorm.Dialector) (*gorm.DB, error) {
	gCfg := getCfg()
	connection, err := gorm.Open(conn, &gCfg)
	if err != nil {
		return nil, err
	}
	connection = connection.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")

	return connection.Session(&gorm.Session{}), nil
}
