package loginController

import (
	sjwt "snappy-api/core/jwt"
	"snappy-api/core/logger"
	"snappy-api/models"
	"snappy-api/models/dbmodels"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"

	routing "github.com/qiangxue/fasthttp-routing"
)

// UserLogin s
func UserLogin(ctx *routing.Context) error {

	logger := logger.GetLogInstance("", "")

	body := &models.LoginModel{}
	if r := jsoniter.Unmarshal(ctx.Request.Body(), &body); r != nil {
		logger.Error(r)
		return r
	}
	expirationTime := time.Now().Add(9999 * time.Minute)

	db := ctx.Get("db").(*gorm.DB)

	user := dbmodels.User{}

	if err := db.Where(&dbmodels.User{UserName: body.UserName}).Or(&dbmodels.User{Password: body.Password}).First(&user).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		res := models.NewResponse(false, nil, "user not found")
		return ctx.WriteData(res.MustMarshal())
	}

	claims := &sjwt.UserClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	tokenString := sjwt.CreateJWT(claims)
	l := models.TokenModel{AccessKey: tokenString}
	r := models.NewResponse(false, l, "OK")
	return ctx.WriteData(r.MustMarshal())
}

func RestaurantLogin(ctx *routing.Context) error {

	logger := logger.GetLogInstance("", "")

	body := &models.LoginModel{}
	if r := jsoniter.Unmarshal(ctx.Request.Body(), &body); r != nil {
		logger.Error(r)
		return r
	}
	expirationTime := time.Now().Add(9999 * time.Minute)

	db := ctx.Get("db").(*gorm.DB)

	user := dbmodels.RestaurantUser{}

	if err := db.Where(&dbmodels.RestaurantUser{UserName: body.UserName}).Or(&dbmodels.RestaurantUser{Password: body.Password}).First(&user).Error; err != nil {
		logger.Error(err)
		ctx.Response.SetStatusCode(404)
		res := models.NewResponse(false, nil, "user not found")
		return ctx.WriteData(res.MustMarshal())
	}

	claims := &sjwt.RestaurantClaims{
		RestaurantID: user.ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	tokenString := sjwt.CreateJWT(claims)
	l := models.TokenModel{AccessKey: tokenString}
	r := models.NewResponse(true, l, "OK")
	return ctx.WriteData(r.MustMarshal())
}