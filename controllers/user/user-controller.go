package controllers

import (
	"snappy-api/core/logger"
	"snappy-api/models"
	dbmodels "snappy-api/models/db-models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	jsoniter "github.com/json-iterator/go"
	routing "github.com/qiangxue/fasthttp-routing"
)

// Create s
func Create(ctx *routing.Context) error {
	item := &dbmodels.User{}

	db := ctx.Get("db").(*gorm.DB)

	if jerr := jsoniter.Unmarshal(ctx.Request.Body(), &item); jerr != nil {
		return jerr
	}

	findItem := &dbmodels.User{}
	dbErr := db.Where(&dbmodels.User{UserName: item.UserName}).First(&findItem).Error

	if dbErr != gorm.ErrRecordNotFound {
		ctx.Response.SetStatusCode(400)
		r := models.NewResponse(false, nil, "this user name already exist.")

		return ctx.WriteData(r.MustMarshal())
	}
	item.IsActive = true
	db.Create(&item)
	r := models.NewResponse(true, nil, "OK")
	return ctx.WriteData(r.MustMarshal())
}

// Update update user with id
func Update(ctx *routing.Context) error {
	logger := logger.GetLogInstance("", "")
	db := ctx.Get("db").(*gorm.DB)
	userID := ctx.Param("id")

	user := &dbmodels.User{}

	if r := jsoniter.Unmarshal(ctx.Request.Body(), &user); r != nil {
		logger.Error(r)
		return r
	}

	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		r := models.NewResponse(false, nil, "user not found")
		return ctx.WriteData(r.MustMarshal())
	}

	if err := jsoniter.Unmarshal(ctx.Request.Body(), &user); err != nil {
		ctx.Response.SetStatusCode(400)
		logger.Error(err)
		r := models.NewResponse(false, nil, "unexpected error")

		return ctx.WriteData(r.MustMarshal())

	}

	db.Save(&user)
	r := models.NewResponse(true, user, "OK")
	return ctx.WriteData(r.MustMarshal())
}

// GetByID get user by id
func GetByID(ctx *routing.Context) error {
	logger := logger.GetLogInstance("", "")
	db := ctx.Get("db").(*gorm.DB)

	user := dbmodels.User{}

	if err := db.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		res := models.NewResponse(false, nil, "not found")
		return ctx.WriteData(res.MustMarshal())
	}
	res := models.NewResponse(true, user, "OK")
	return ctx.WriteData(res.MustMarshal())
}

// GetAll get all user
func GetAll(ctx *routing.Context) error {
	db := ctx.Get("db").(*gorm.DB)
	users := []dbmodels.User{}
	db.Find(&users)

	res := models.NewResponse(true, users, "OK")

	return ctx.WriteData(res.MustMarshal())
}
