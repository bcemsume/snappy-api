package database

import (
	"fmt"
	"snappy-api/core/config"
	"snappy-api/core/logger"
	dbmodels "snappy-api/models/db-models"

	"github.com/jinzhu/gorm"
)

var log = logger.GetLogInstance("", "")

func InitDB() *gorm.DB {
	dbConf := config.DBConfigs()
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", dbConf[config.DBHOST], dbConf[config.DBUSER], dbConf[config.DBNAME], dbConf[config.SSLMODE], dbConf[config.DBPASS]) //Build connection string
	log.TextInfo(dbURI)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	log.TextInfo("db connected.")
	db.AutoMigrate(&dbmodels.User{})
	return db
}
