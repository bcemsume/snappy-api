package controllers

import (
	"snappy-api/models"
	dbmodels "snappy-api/models/dbmodels"

	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"
	routing "github.com/qiangxue/fasthttp-routing"
)

// Create s
func Create(ctx *routing.Context) error {
	item := &dbmodels.RestaurantUser{}

	db := ctx.Get("db").(*gorm.DB)

	if jerr := jsoniter.Unmarshal(ctx.Request.Body(), &item); jerr != nil {
		return jerr
	}

	rest := &dbmodels.Restaurant{}

	findItem := &dbmodels.RestaurantUser{}
	dbErr := db.Where(&dbmodels.RestaurantUser{Email: item.Email}).Or(&dbmodels.RestaurantUser{UserName: item.UserName}).First(&findItem).Error

	if dbErr != gorm.ErrRecordNotFound {
		ctx.Response.SetStatusCode(400)
		r := models.NewResponse(false, nil, "this email or username already exist.")

		return ctx.WriteData(r.MustMarshal())
	}
	item.IsActive = true
	db.Preload("Users").Where("id = ?", ctx.Param("id")).First(&rest)

	r := models.NewResponse(true, rest.Users, "OK")
	return ctx.WriteData(r.MustMarshal())
}
