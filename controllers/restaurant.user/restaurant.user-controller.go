package controllers

import (
	"snappy-api/models"
	dbmodels "snappy-api/models/dbmodels"

	sjwt "snappy-api/core/jwt"

	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"
	routing "github.com/qiangxue/fasthttp-routing"
)

// Create s
func Create(ctx *routing.Context) error {
	item := dbmodels.RestaurantUser{}

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
	if dbErr := db.Preload("Users").Where("id = ?", ctx.Param("id")).Find(&rest).Error; dbErr == gorm.ErrRecordNotFound {
		ctx.Response.SetStatusCode(400)
		r := models.NewResponse(false, nil, "restaurant not found")

		return ctx.WriteData(r.MustMarshal())
	}

	rest.Users = []*dbmodels.RestaurantUser{&item}
	db.Save(rest)
	r := models.NewResponse(true, rest.Users, "OK")
	return ctx.WriteData(r.MustMarshal())
}

// Get s
func Get(ctx *routing.Context) error {
	tkn := string(ctx.Request.Header.Peek("Authorization"))
	_, tValue := sjwt.ValidateJWT(tkn)
	db := ctx.Get("db").(*gorm.DB)

	restUsr := &dbmodels.RestaurantUser{}
	usr := &models.RestaurantUserModel{}
	dbErr := db.Preload("Restaurants").First(&restUsr, tValue.(*sjwt.Claims).UserID).Scan(&usr).Error

	usr.RestaurantID = restUsr.Restaurants[0].ID
	if dbErr == gorm.ErrRecordNotFound {
		ctx.Response.SetStatusCode(400)
		r := models.NewResponse(false, nil, "user not found")

		return ctx.WriteData(r.MustMarshal())
	}

	r := models.NewResponse(true, usr, "OK")
	return ctx.WriteData(r.MustMarshal())
}
