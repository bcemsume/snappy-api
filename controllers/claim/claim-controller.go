package controllers

import (
	"snappy-api/models"
	"snappy-api/models/dbmodels"
	"time"

	routing "github.com/qiangxue/fasthttp-routing"

	"snappy-api/core/jwt"
	sjwt "snappy-api/core/jwt"

	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"
)

// AddClaim s
func AddClaim(ctx *routing.Context) error {
	body := &models.ClaimModel{}

	db := ctx.Get("db").(*gorm.DB)

	tkn := string(ctx.Request.Header.Peek("Authorization"))

	_, tokn := sjwt.ValidateJWT(tkn)

	if jerr := jsoniter.Unmarshal(ctx.Request.Body(), &body); jerr != nil {
		return jerr
	}

	cmp := &dbmodels.Campaign{}

	if e := db.Model(&dbmodels.Claim{QRID: body.QRID}).Error; e != gorm.ErrRecordNotFound {
		ctx.Response.SetStatusCode(400)
		r := models.NewResponse(true, nil, "QR Kod daha önceden kullanılmış.")

		return ctx.WriteData(r.MustMarshal())
	}

	if e := db.Model(&dbmodels.Campaign{}).Where("id = ?", body.CampaignID).First(cmp).Error; e == gorm.ErrRecordNotFound {
		ctx.Response.SetStatusCode(400)
		r := models.NewResponse(true, nil, "Kampanya bulunamadi")

		return ctx.WriteData(r.MustMarshal())
	}

	c := &dbmodels.Claim{}
	dbErr := db.Where(&dbmodels.Claim{UserID: tokn.(*sjwt.Claims).UserID, CampaingID: body.CampaignID, IsUsed: false}).Where("claim != ?", cmp.Claim).First(&c).Error

	if dbErr == gorm.ErrRecordNotFound {

		c.CampaingID = body.CampaignID
		c.UserID = tokn.(*sjwt.Claims).UserID
		c.Claim = 1
		c.IsUsed = false

		db.Save(c)

		ctx.Response.SetStatusCode(200)
		r := models.NewResponse(true, nil, "OK")
		return ctx.WriteData(r.MustMarshal())

	} else if cmp.Claim == c.Claim {

		db.Save(&dbmodels.Claim{
			CampaingID: body.CampaignID,
			UserID:     tokn.(*sjwt.Claims).UserID,
			IsUsed:     false,
			Claim:      1,
		})
		ctx.Response.SetStatusCode(200)
		r := models.NewResponse(true, nil, "OK")
		return ctx.WriteData(r.MustMarshal())
	}
	c.QRID = body.QRID
	c.Claim++
	db.Save(c)
	cl := &dbmodels.ClaimEvent{
		CampaignID: c.CampaingID,
		UserID:     c.UserID,
		ScanDate:   time.Now(),
	}

	db.Save(cl)
	r := models.NewResponse(true, nil, "OK")
	return ctx.WriteData(r.MustMarshal())
}

// GetRewards s
func GetRewards(ctx *routing.Context) error {

	db := ctx.Get("db").(*gorm.DB)
	tkn := string(ctx.Request.Header.Peek("Authorization"))
	_, tokn := sjwt.ValidateJWT(tkn)

	c := &[]models.RewardModel{}
	if e := db.Model(&dbmodels.Claim{}).Select("claims.claim as user_claim, campaigns.claim as campaign_claim, products.description as product_name, restaurants.title as restaurant_title").Joins("join users on users.id = claims.user_id").Joins("join campaigns on campaigns.id = claims.campaing_id").Joins("join products on products.id = campaigns.product_id").Joins("join restaurants on restaurants.id = products.restaurant_id").Where("claims.user_id = ?", tokn.(*jwt.Claims).UserID).Scan(c); e != nil {

	}
	r := models.NewResponse(true, c, "OK")
	return ctx.WriteData(r.MustMarshal())
}
