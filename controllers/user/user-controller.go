package controllers

import (
	sjwt "snappy-api/core/jwt"
	"snappy-api/core/logger"
	"snappy-api/models"
	"snappy-api/models/dbmodels"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	jsoniter "github.com/json-iterator/go"
	routing "github.com/qiangxue/fasthttp-routing"
)

// Create s
func Create(ctx *routing.Context) error {
	item := &models.UserCreateModel{}
	db := ctx.Get("db").(*gorm.DB)

	if jerr := jsoniter.Unmarshal(ctx.Request.Body(), &item); jerr != nil {
		return jerr
	}

	findItem := &dbmodels.User{}
	dbErr := db.Where(&dbmodels.User{UserName: item.UserName}).Or(&dbmodels.User{PhoneNumber: item.PhoneNumber}).First(&findItem).Error

	if dbErr != gorm.ErrRecordNotFound {
		ctx.Response.SetStatusCode(400)
		r := models.NewResponse(false, nil, "this user name or phone number exist.")

		return ctx.WriteData(r.MustMarshal())
	}
	findItem = &dbmodels.User{
		IsActive:    true,
		UserName:    item.UserName,
		PhoneNumber: item.PhoneNumber,
		Password:    item.Password,
		DeviceID:    item.DeviceID,
		FCMToken:    item.FCMToken,
	}

	db.Create(&item)
	r := models.NewResponse(true, nil, "OK")
	return ctx.WriteData(r.MustMarshal())
}

// UserCheckPhoneNumber s
func UserCheckPhoneNumber(ctx *routing.Context) error {
	item := &models.CheckUserPhoneNumberModel{}
	db := ctx.Get("db").(*gorm.DB)

	if jerr := jsoniter.Unmarshal(ctx.Request.Body(), &item); jerr != nil {
		return jerr
	}

	findItem := &dbmodels.User{}
	dbErr := db.Where(&dbmodels.User{PhoneNumber: item.PhoneNumber}).First(&findItem).Error

	if dbErr != gorm.ErrRecordNotFound {
		ctx.Response.SetStatusCode(400)
		r := models.NewResponse(false, nil, "this user name or phone number exist.")

		return ctx.WriteData(r.MustMarshal())
	}
	r := models.NewResponse(true, nil, "OK")
	return ctx.WriteData(r.MustMarshal())
}

// Update update user with id
func Update(ctx *routing.Context) error {
	logger := logger.GetLogInstance("", "")
	db := ctx.Get("db").(*gorm.DB)
	userID := ctx.Param("id")

	body := &models.UserProfileModel{}
	user := &dbmodels.User{}

	if r := jsoniter.Unmarshal(ctx.Request.Body(), &body); r != nil {
		logger.Error(r)
		return r
	}
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		r := models.NewResponse(false, nil, "user not found")
		return ctx.WriteData(r.MustMarshal())
	}
	layout := "2006-01-02 15:04:05 -0700 MST"
	t, _ := time.Parse(layout, body.BirthDay.String())
	user.BirthDay = t
	user.Name = body.Name
	user.LastName = body.LastName
	user.Email = body.Email
	user.Gender = body.Gender

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

func GetUserDetail(ctx *routing.Context) error {

	logger := logger.GetLogInstance("", "")
	db := ctx.Get("db").(*gorm.DB)

	user := dbmodels.User{}
	tkn := string(ctx.Request.Header.Peek("Authorization"))

	_, tokn := sjwt.ValidateJWT(tkn)

	if err := db.Where("id = ?", tokn.(*sjwt.Claims).UserID).First(&user).Error; err != nil {
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
