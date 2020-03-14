package controllers

import (
	"snappy-api/core/database"
	"snappy-api/models"
	dbmodels "snappy-api/models/db-models"

	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"
	routing "github.com/qiangxue/fasthttp-routing"
)

func Create(ctx *routing.Context) error {
	item := &dbmodels.User{}

	var db = database.InitDB()
	defer db.Close()

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
		ctx.WriteData(res)
		return nil
	}
	db.Create(&item)
	res, _ := jsoniter.Marshal(r)
	ctx.WriteData(res)

	return nil
}
