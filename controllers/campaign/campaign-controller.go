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
	item := &dbmodels.Campaign{}

	db := ctx.Get("db").(*gorm.DB)

	if jerr := jsoniter.Unmarshal(ctx.Request.Body(), &item); jerr != nil {
		return jerr
	}

	findItem := &dbmodels.Campaign{}
	dbErr := db.Where(&dbmodels.Campaign{ProductID: item.ProductID}).First(&findItem).Error

	if dbErr != gorm.ErrRecordNotFound {
		ctx.Response.SetStatusCode(400)
		r := models.NewResponse(false, nil, "this cmp already exist.")

		return ctx.WriteData(r.MustMarshal())
	}
	db.Create(&item)
	r := models.NewResponse(true, nil, "OK")
	return ctx.WriteData(r.MustMarshal())
}

// Update update campaign with id
func Update(ctx *routing.Context) error {
	logger := logger.GetLogInstance("", "")
	db := ctx.Get("db").(*gorm.DB)
	id := ctx.Param("id")

	cmp := &dbmodels.Campaign{}

	if r := jsoniter.Unmarshal(ctx.Request.Body(), &cmp); r != nil {
		logger.Error(r)
		return r
	}

	if err := db.Where("id = ?", id).First(&cmp).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		r := models.NewResponse(false, nil, "cmp not found")
		return ctx.WriteData(r.MustMarshal())
	}

	if err := jsoniter.Unmarshal(ctx.Request.Body(), &cmp); err != nil {
		ctx.Response.SetStatusCode(400)
		logger.Error(err)
		r := models.NewResponse(false, nil, "unexpected error")

		return ctx.WriteData(r.MustMarshal())

	}

	db.Save(&cmp)
	r := models.NewResponse(true, cmp, "OK")
	return ctx.WriteData(r.MustMarshal())
}

// GetByID get campaign by id
func GetByID(ctx *routing.Context) error {
	logger := logger.GetLogInstance("", "")
	db := ctx.Get("db").(*gorm.DB)

	cmp := models.CampaingModel{}

	if err := db.Model(&[]dbmodels.Campaign{}).Where("id = ?", ctx.Param("id")).Scan(&cmp).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		res := models.NewResponse(false, nil, "not found")
		return ctx.WriteData(res.MustMarshal())
	}
	res := models.NewResponse(true, cmp, "OK")
	return ctx.WriteData(res.MustMarshal())
}

// GetAll get all campaign
func GetAll(ctx *routing.Context) error {
	db := ctx.Get("db").(*gorm.DB)
	data := []models.CampaingModel{}
	db.Model(&dbmodels.Campaign{}).Select("campaigns.id, campaigns.claim,campaigns.product_id, campaigns.finish_date, products.description ").Joins("join products on campaigns.product_id = products.id").Scan(&data)
	res := models.NewResponse(true, data, "OK")

	return ctx.WriteData(res.MustMarshal())
}

//GetProducts s
func GetProducts(ctx *routing.Context) error {
	db := ctx.Get("db").(*gorm.DB)
	cmp := []models.CampaingModel{}
	db.Model(dbmodels.Product{}).Where("product_id = ?", ctx.Param("id")).Related(&dbmodels.Campaign{}).Scan(&cmp)
	res := models.NewResponse(true, cmp, "OK")
	return ctx.WriteData(res.MustMarshal())
}
