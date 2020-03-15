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

	r := models.ResponseMessage{Message: "OK", IsSucceeded: true}

	if jerr := jsoniter.Unmarshal(ctx.Request.Body(), &item); jerr != nil {
		return jerr
	}

	findItem := &dbmodels.User{}
	dbErr := db.Where(&dbmodels.User{UserName: item.UserName}).First(&findItem).Error

	if dbErr != gorm.ErrRecordNotFound {
		r.IsSucceeded = false
		r.Message = "this user name already exist."
		ctx.Response.SetStatusCode(400)
		res, _ := jsoniter.Marshal(r)
		return ctx.WriteData(res)
	}
	item.IsActive = true
	db.Create(&item)
	res, _ := jsoniter.Marshal(r)

	return ctx.WriteData(res)
}

// Update update user with id
func Update(ctx *routing.Context) error {
	logger := logger.GetLogInstance("", "")
	db := ctx.Get("db").(*gorm.DB)

	user := &dbmodels.User{}
	r := models.ResponseMessage{Message: "OK", IsSucceeded: true}

	if r := jsoniter.Unmarshal(ctx.Request.Body(), &user); r != nil {
		logger.Error(r)
		return r
	}

	if err := db.Where("id = ?", user.ID).First(&user).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		res, _ := jsoniter.Marshal(r)

		return ctx.WriteData(res)
	}

	if err := jsoniter.Unmarshal(ctx.Request.Body(), &user); err != nil {
		ctx.Response.SetStatusCode(400)
		logger.Error(err)
		r.IsSucceeded = false
		r.Message = err.Error()
		res, _ := jsoniter.Marshal(r)

		return ctx.WriteData(res)

	}

	db.Save(&user)
	res, _ := jsoniter.Marshal(r)
	return ctx.WriteData(res)
}

// GetByID get user by id
func GetByID(ctx *routing.Context) error {
	logger := logger.GetLogInstance("", "")
	db := ctx.Get("db").(*gorm.DB)

	user := dbmodels.User{}
	r := models.ResponseMessage{Message: "OK", IsSucceeded: true}

	if err := db.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		r.IsSucceeded = false
		r.Message = "user not found."
		res, err := jsoniter.Marshal(r)
		if err != nil {
			return err
		}
		return ctx.WriteData(res)
	}
	res, err := jsoniter.Marshal(user)
	if err != nil {
		return err
	}
	return ctx.WriteData(res)
}

// GetAll get all user
func GetAll(ctx *routing.Context) error {

	return nil
}
