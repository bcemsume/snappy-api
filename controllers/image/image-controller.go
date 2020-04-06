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
	item := &dbmodels.Image{}
	db := ctx.Get("db").(*gorm.DB)

	if jerr := jsoniter.Unmarshal(ctx.Request.Body(), &item); jerr != nil {
		logger.Error(jerr)
		return jerr
	}

	db.Create(&item)
	r := models.NewResponse(true, nil, "OK")
	return ctx.WriteData(r.MustMarshal())
}

// Update update image with id
func Update(ctx *routing.Context) error {
	logger := logger.GetLogInstance("", "")
	db := ctx.Get("db").(*gorm.DB)
	imageID := ctx.Param("id")

	img := &dbmodels.Image{}

	if r := jsoniter.Unmarshal(ctx.Request.Body(), &img); r != nil {
		logger.Error(r)
		return r
	}

	if err := db.Where("id = ?", imageID).First(&img).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		r := models.NewResponse(false, nil, "image not found")
		return ctx.WriteData(r.MustMarshal())
	}

	if err := jsoniter.Unmarshal(ctx.Request.Body(), &img); err != nil {
		ctx.Response.SetStatusCode(400)
		logger.Error(err)
		r := models.NewResponse(false, nil, "unexpected error")

		return ctx.WriteData(r.MustMarshal())

	}

	db.Save(&img)
	r := models.NewResponse(true, img, "OK")
	return ctx.WriteData(r.MustMarshal())
}

// GetByID get image by id
func GetByID(ctx *routing.Context) error {
	logger := logger.GetLogInstance("", "")
	db := ctx.Get("db").(*gorm.DB)

	img := dbmodels.Image{}

	if err := db.Where("id = ?", ctx.Param("id")).First(&img).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		res := models.NewResponse(false, nil, "not found")
		return ctx.WriteData(res.MustMarshal())
	}
	res := models.NewResponse(true, img, "OK")
	return ctx.WriteData(res.MustMarshal())
}

// GetAll get all image
func GetAll(ctx *routing.Context) error {
	db := ctx.Get("db").(*gorm.DB)
	data := []dbmodels.Image{}
	db.Find(&data)

	res := models.NewResponse(true, data, "OK")

	return ctx.WriteData(res.MustMarshal())
}
