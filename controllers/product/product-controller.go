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
	logger := logger.GetLogInstance("", "")
	item := &dbmodels.Product{}
	db := ctx.Get("db").(*gorm.DB)

	if jerr := jsoniter.Unmarshal(ctx.Request.Body(), &item); jerr != nil {
		logger.Error(jerr)
		return jerr
	}

	// findItem := &dbmodels.Product{}
	// dbErr := db.Where(&dbmodels.Product{Description: item.Description}).First(&findItem).Error

	// if dbErr != gorm.ErrRecordNotFound {
	// 	ctx.Response.SetStatusCode(400)
	// 	r := models.NewResponse(false, nil, "this product already exist.")

	// 	return ctx.WriteData(r.MustMarshal())
	// }
	db.Create(&item)
	r := models.NewResponse(true, nil, "OK")
	return ctx.WriteData(r.MustMarshal())
}

// Update update product with id
func Update(ctx *routing.Context) error {
	logger := logger.GetLogInstance("", "")
	db := ctx.Get("db").(*gorm.DB)
	productID := ctx.Param("id")

	product := &dbmodels.Product{}

	if r := jsoniter.Unmarshal(ctx.Request.Body(), &product); r != nil {
		logger.Error(r)
		return r
	}

	if err := db.Where("id = ?", productID).First(&product).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		r := models.NewResponse(false, nil, "product not found")
		return ctx.WriteData(r.MustMarshal())
	}

	if err := jsoniter.Unmarshal(ctx.Request.Body(), &product); err != nil {
		ctx.Response.SetStatusCode(400)
		logger.Error(err)
		r := models.NewResponse(false, nil, "unexpected error")

		return ctx.WriteData(r.MustMarshal())

	}

	db.Save(&product)
	r := models.NewResponse(true, product, "OK")
	return ctx.WriteData(r.MustMarshal())
}

// GetByID get product by id
func GetByID(ctx *routing.Context) error {
	logger := logger.GetLogInstance("", "")
	db := ctx.Get("db").(*gorm.DB)

	product := models.ProductModel{}

	if err := db.Model(&dbmodels.Product{}).Where("id = ?", ctx.Param("id")).Scan(&product).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		res := models.NewResponse(false, nil, "not found")
		return ctx.WriteData(res.MustMarshal())
	}
	res := models.NewResponse(true, product, "OK")
	return ctx.WriteData(res.MustMarshal())
}

// GetAll get all product
func GetAll(ctx *routing.Context) error {
	db := ctx.Get("db").(*gorm.DB)
	data := []models.ProductModel{}
	db.Model(&dbmodels.Product{}).Scan(&data)

	res := models.NewResponse(true, data, "OK")

	return ctx.WriteData(res.MustMarshal())
}

// GetCampaigns s
func GetCampaigns(ctx *routing.Context) error {
	db := ctx.Get("db").(*gorm.DB)
	cmp := []models.CampaingModel{}
	db.Model(&dbmodels.Product{}).Where("product_id = ?", ctx.Param("id")).Related(&dbmodels.Campaign{}).Scan(&cmp)
	res := models.NewResponse(true, cmp, "OK")
	return ctx.WriteData(res.MustMarshal())
}
