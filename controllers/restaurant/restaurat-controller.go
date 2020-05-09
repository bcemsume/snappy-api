package controllers

import (
	"snappy-api/core/logger"
	"snappy-api/models"
	dbmodels "snappy-api/models/dbmodels"
	"strconv"

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

	body := &models.RestaurantModel{}

	rest := &dbmodels.Restaurant{}

	if r := jsoniter.Unmarshal(ctx.Request.Body(), &body); r != nil {
		logger.Error(r)
		return r
	}

	if err := db.Model(&dbmodels.Restaurant{}).Where("id = ?", restID).First(&rest).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		r := models.NewResponse(false, nil, "restaurant not found")
		return ctx.WriteData(r.MustMarshal())
	}
	if body.Logo != "" {
		img := &dbmodels.Image{}
		if e := db.Where(&dbmodels.Image{RestaurantID: body.ID, Type: 2}).Find(img).Error; e != gorm.ErrRecordNotFound {
			db.Delete(img)
		}
		i := &dbmodels.Image{
			RestaurantID: body.ID,
			ImageURL:     body.Logo,
			Order:        0,
			Type:         2,
		}

		db.Save(i)
	}

	rest.Address = body.Address
	rest.Email = body.Address
	rest.Lang = body.Lang
	rest.Long = body.Long
	rest.PaymentMethods = body.PaymentMethods
	rest.WorkingDays = body.WorkingDays
	rest.Phone = body.Phone
	rest.Title = body.Title

	db.Save(&rest)
	r := models.NewResponse(true, rest, "OK")
	return ctx.WriteData(r.MustMarshal())
}

// GetByID get restaurant by id
func GetByID(ctx *routing.Context) error {
	logger := logger.GetLogInstance("", "")
	db := ctx.Get("db").(*gorm.DB)

	rest := models.RestaurantModel{}

	if err := db.Model(&dbmodels.Restaurant{}).Where("id = ?", ctx.Param("id")).Scan(&rest).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		res := models.NewResponse(false, nil, "not found")
		return ctx.WriteData(res.MustMarshal())
	}

	logo := dbmodels.Image{}
	if err := db.Where(&dbmodels.Image{RestaurantID: rest.ID, Type: 2}).First(&logo).Error; err == nil {
		rest.Logo = logo.ImageURL
	}

	imgs := []dbmodels.Image{}
	if err := db.Where(&dbmodels.Image{RestaurantID: rest.ID, Type: 1}).Find(&imgs).Error; err == nil {

		for _, elem := range imgs {
			i := &models.Image{
				ImageURL: elem.ImageURL,
				Order:    elem.Order,
				Type:     elem.Type,
			}
			rest.Images = append(rest.Images, *i)
		}

	}

	cmp := []models.CampaingModel{}
	if err := db.Model(&dbmodels.Campaign{}).Select("campaigns.id, campaigns.claim,campaigns.product_id, campaigns.finish_date, products.description ").Joins("join products on campaigns.product_id = products.id").Where("products.restaurant_id = ?", ctx.Param("id")).Scan(&cmp).Error; err == nil {
		rest.Campaigns = cmp
	}

	res := models.NewResponse(true, rest, "OK")
	return ctx.WriteData(res.MustMarshal())
}

// GetAll get all restaurant
func GetAll(ctx *routing.Context) error {
	db := ctx.Get("db").(*gorm.DB)
	data := []models.RestaurantListModel{}
	db.Model(&dbmodels.Restaurant{}).Select("restaurants.id, restaurants.title,restaurants.lang, restaurants.long images.image_url as logo").Joins("join images on restaurants.id = images.restaurant_id").Where("images.type = 2 and images.deleted_at is null").Scan(&data)

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
	data := []models.Image{}
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	db.Model(&dbmodels.Image{}).Where("restaurant_id = ? and type = 1", id).Scan(&data)
	res := models.NewResponse(true, &models.ImageModel{RestaurantID: 1 << id, Images: data}, "OK")
	return ctx.WriteData(res.MustMarshal())
}

// AddImages s
func AddImages(ctx *routing.Context) error {
	db := ctx.Get("db").(*gorm.DB)
	logger := logger.GetLogInstance("", "")
	body := models.ImageModel{}

	if r := jsoniter.Unmarshal(ctx.Request.Body(), &body); r != nil {
		logger.Error(r)
		return r
	}

	if e := db.Find(&dbmodels.Restaurant{}, body.RestaurantID).Error; e == gorm.ErrRecordNotFound {
		res := models.NewResponse(false, nil, "restaurant not found")
		return ctx.WriteData(res.MustMarshal())
	}
	img := &[]dbmodels.Image{}
	if e := db.Where(&dbmodels.Image{RestaurantID: body.RestaurantID, Type: 1}).Find(img).Error; e != gorm.ErrRecordNotFound {
		db.Delete(img)
	}

	for _, elem := range body.Images {
		i := &dbmodels.Image{
			RestaurantID: body.RestaurantID,
			ImageURL:     elem.ImageURL,
			Order:        elem.Order,
			Type:         elem.Type,
		}

		db.Save(i)
	}

	res := models.NewResponse(true, nil, "OK")
	return ctx.WriteData(res.MustMarshal())
}
