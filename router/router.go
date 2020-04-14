package router

import (
	campaign "snappy-api/controllers/campaign"
	image "snappy-api/controllers/image"
	login "snappy-api/controllers/login"
	"snappy-api/models"

	resUser "snappy-api/controllers/restaurant.user"

	product "snappy-api/controllers/product"
	restaurant "snappy-api/controllers/restaurant"

	user "snappy-api/controllers/user"

	"snappy-api/core/database"

	sjwt "snappy-api/core/jwt"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

// Route ss
func Route() fasthttp.RequestHandler {

	router := routing.New()

	db := database.InitDB()
	router.Use(func(c *routing.Context) error {
		c.Set("db", db)
		c.Response.Header.Set("Content-Type", "application/json")

		return c.Next()
	})

	router.Post("/token/user", login.UserLogin)
	router.Post("/token/restaurant", login.RestaurantLogin)
	router.Post("user", user.Create)

	api := router.Group("/api/")

	api.Use(func(c *routing.Context) error {
		tkn := string(c.Request.Header.Peek("Authorization"))
		tknValidate := sjwt.ValidateJWT(tkn)
		if tknValidate == false {
			c.SetStatusCode(401)
			r := models.NewResponse(false, nil, "token not valid")
			c.Abort()
			return c.WriteData(r.MustMarshal())
		}
		return c.Next()
	})
	// user
	api.Get("user/<id>", user.GetByID)
	api.Get("user", user.GetAll)
	api.Put("user/<id>", user.Update)
	// restaurant
	api.Get("restaurant/<id>", restaurant.GetByID)
	api.Get("restaurant", restaurant.GetAll)
	api.Get("restaurant/<id>/products", restaurant.GetProducts)
	api.Get("restaurant/<id>/images", restaurant.GetImages)
	// api.Get("restaurant/<id>/campaigns", restaurant.GetCampaigns)

	api.Put("restaurant/<id>", restaurant.Update)
	api.Post("restaurant", restaurant.Create)

	// product
	api.Post("product", product.Create)
	api.Get("product/<id>", product.GetByID)
	api.Get("product/<id>/campaigns", product.GetCampaigns)
	api.Get("product", product.GetAll)
	api.Put("product/<id>", product.Update)

	// image
	api.Post("image", image.Create)
	api.Get("image/<id>", image.GetByID)
	api.Get("image", image.GetAll)
	api.Put("image/<id>", image.Update)

	// campaign
	api.Post("campaign", campaign.Create)
	api.Get("campaign/<id>", campaign.GetByID)
	api.Get("campaign/<id>/products", campaign.GetProducts)
	api.Get("campaign", campaign.GetAll)
	api.Put("campaign/<id>", campaign.Update)

	// restaurant-user
	api.Post("restaurant-user", resUser.Get)

	api.Post("restaurant-user/<id>", resUser.Create)

	return router.HandleRequest
}
