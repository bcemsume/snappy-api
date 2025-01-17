package router

import (
	campaign "snappy-api/controllers/campaign"
	claim "snappy-api/controllers/claim"
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

		if err := db.DB().Ping(); err != nil {
			db = database.InitDB()
		}
		c.Set("db", db)
		c.Response.Header.Set("Content-Type", "application/json")

		return c.Next()
	})

	router.Post("/token/user", login.UserLogin)
	router.Post("/token/restaurant", login.RestaurantLogin)
	router.Post("/user", user.Create)
	router.Post("/user/check-phone-number", user.UserCheckPhoneNumber)

	api := router.Group("/api/")

	api.Use(func(c *routing.Context) error {
		tkn := string(c.Request.Header.Peek("Authorization"))
		tknValidate, _ := sjwt.ValidateJWT(tkn)
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
	api.Get("user-profile", user.GetUserDetail)

	// restaurant
	api.Get("restaurants", restaurant.GetAll)
	api.Get("restaurant", restaurant.GetByID)
	api.Get("restaurants/<id>", restaurant.GetByIDUser)

	api.Get("restaurant/<id>/products", restaurant.GetProducts)
	api.Get("restaurant/<id>/images", restaurant.GetImages)
	api.Get("restaurant/<id>/campaigns", restaurant.GetCampaigns)
	api.Post("restaurant/image", restaurant.AddImages)

	api.Put("restaurant/<id>", restaurant.Update)
	api.Post("restaurant", restaurant.Create)

	// product
	api.Post("product", product.Create)
	api.Get("product/<id>", product.GetByID)
	api.Get("product/<id>/campaigns", product.GetCampaigns)
	api.Get("product", product.GetAll)
	api.Put("product/<id>", product.Update)
	api.Delete("product/<id>", product.Delete)

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
	api.Delete("campaign/<id>", campaign.Delete)

	// restaurant-user
	api.Get("restaurant-user", resUser.Get)
	api.Post("restaurant-user/<id>", resUser.Create)

	// claim
	api.Post("claim", claim.AddClaim)
	api.Get("claim/rewards", claim.GetRewards)

	return router.HandleRequest
}
