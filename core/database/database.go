package database

import (
	"fmt"
	"snappy-api/core/config"
	"snappy-api/core/logger"
	dbmodels "snappy-api/models/dbmodels"

	"github.com/jinzhu/gorm"
)

var log = logger.GetLogInstance("", "")

// InitDB s
func InitDB() *gorm.DB {
	dbConf := config.DBConfigs()
	appConf := config.GetAppConfig()
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", dbConf[config.DBHOST], dbConf[config.DBUSER], dbConf[config.DBNAME], dbConf[config.SSLMODE], dbConf[config.DBPASS]) //Build connection string
	log.TextInfo(dbURI)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	log.TextInfo("db connected.")
	db.AutoMigrate(&dbmodels.User{}, &dbmodels.Campaign{}, &dbmodels.ClaimEvent{},
		&dbmodels.Image{}, &dbmodels.Product{}, &dbmodels.Restaurant{})
	addForeignKeys(db)
	if appConf[config.DEBUGMODE] == "true" {
		db = db.Debug()
	}
	return db
}

func addForeignKeys(db *gorm.DB) {
	db.Model(&dbmodels.Product{}).AddForeignKey("restaurant_id", "restaurants(id)", "RESTRICT", "RESTRICT")
	db.Model(&dbmodels.Image{}).AddForeignKey("restaurant_id", "restaurants(id)", "RESTRICT", "RESTRICT")
	db.Model(&dbmodels.ClaimEvent{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&dbmodels.Campaign{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")

}
