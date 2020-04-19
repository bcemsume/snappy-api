package controllers

import (
	"snappy-api/core/logger"
	"snappy-api/models"
	dbmodels "snappy-api/models/dbmodels"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	jsoniter "github.com/json-iterator/go"
	routing "github.com/qiangxue/fasthttp-routing"
)

// Create s
func Create(ctx *routing.Context) error {
	item := &dbmodels.Restaurant{}

	db := ctx.Get("db").(*gorm.DB)

	if jerr := jsoniter.Unmarshal(ctx.Request.Body(), &item); jerr != nil {
		return jerr
	}

	findItem := &dbmodels.Restaurant{}
	dbErr := db.Where(&dbmodels.Restaurant{Email: item.Email}).First(&findItem).Error

	if dbErr != gorm.ErrRecordNotFound {
		ctx.Response.SetStatusCode(400)
		r := models.NewResponse(false, nil, "this email already exist.")

		return ctx.WriteData(r.MustMarshal())
	}
	item.IsActive = true
	db.Create(&item)
	r := models.NewResponse(true, nil, "OK")
	return ctx.WriteData(r.MustMarshal())
}

// Update update restaurant with id
func Update(ctx *routing.Context) error {
	logger := logger.GetLogInstance("", "")
	db := ctx.Get("db").(*gorm.DB)
	restID := ctx.Param("id")

	rest := &dbmodels.Restaurant{}

	if r := jsoniter.Unmarshal(ctx.Request.Body(), &rest); r != nil {
		logger.Error(r)
		return r
	}

	if err := db.Where("id = ?", restID).First(&rest).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		r := models.NewResponse(false, nil, "restaurant not found")
		return ctx.WriteData(r.MustMarshal())
	}

	if err := jsoniter.Unmarshal(ctx.Request.Body(), &rest); err != nil {
		ctx.Response.SetStatusCode(400)
		logger.Error(err)
		r := models.NewResponse(false, nil, "unexpected error")

		return ctx.WriteData(r.MustMarshal())
	}

	db.Save(&rest)
	r := models.NewResponse(true, rest, "OK")
	return ctx.WriteData(r.MustMarshal())
}

// GetByID get restaurant by id
func GetByID(ctx *routing.Context) error {
	logger := logger.GetLogInstance("", "")
	db := ctx.Get("db").(*gorm.DB)

	rest := dbmodels.Restaurant{}

	if err := db.Where("id = ?", ctx.Param("id")).First(&rest).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		res := models.NewResponse(false, nil, "not found")
		return ctx.WriteData(res.MustMarshal())
	}
	res := models.NewResponse(true, rest, "OK")
	return ctx.WriteData(res.MustMarshal())
}

// GetAll get all restaurant
func GetAll(ctx *routing.Context) error {
	db := ctx.Get("db").(*gorm.DB)
	data := []dbmodels.Restaurant{}
	db.Find(&data)
	res := models.NewResponse(true, data, "OK")

	return ctx.WriteData(res.MustMarshal())
}

// GetProducts s
func GetProducts(ctx *routing.Context) error {
	db := ctx.Get("db").(*gorm.DB)
	prd := []models.ProductModel{}
	db.Table("products").Where("deleted_at is null and restaurant_id = ?", ctx.Param("id")).Select("products.description, products.price, products.id, products.finish_date, products.restaurant_id").Scan(&prd)

	// db.Model(&dbmodels.Restaurant{}).Where("restaurant_id = ?", ctx.Param("id")).Preload("Campaigns").Related(&prd).Find(&prd)
	res := models.NewResponse(true, prd, "OK")
	return ctx.WriteData(res.MustMarshal())
}

// GetCampaigns s
func GetCampaigns(ctx *routing.Context) error {
	db := ctx.Get("db").(*gorm.DB)
	prd := []models.CampaingModel{}
	db.Model(&dbmodels.Campaign{}).Select("campaigns.id, campaigns.claim,campaigns.product_id, campaigns.finish_date, products.description ").Joins("join products on campaigns.product_id = products.id").Where("products.restaurant_id = ?", ctx.Param("id")).Scan(&prd)

	// db.Model(&dbmodels.Restaurant{}).Where("restaurant_id = ?", ctx.Param("id")).Preload("Campaigns").Related(&prd).Find(&prd)
	res := models.NewResponse(true, prd, "OK")
	return ctx.WriteData(res.MustMarshal())
}

// GetImages s
func GetImages(ctx *routing.Context) error {

	db := ctx.Get("db").(*gorm.DB)
	data := []models.ImageModel{}
	db.Model(&dbmodels.Image{}).Where("restaurant_id = ?", ctx.Param("id")).Scan(&data)
	res := models.NewResponse(true, data, "OK")
	return ctx.WriteData(res.MustMarshal())
}
