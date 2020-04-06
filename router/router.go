package router

import (
	"math"
	campaign "snappy-api/controllers/campaign"
	image "snappy-api/controllers/image"

	product "snappy-api/controllers/product"
	restaurant "snappy-api/controllers/restaurant"

	user "snappy-api/controllers/user"

	"snappy-api/core/database"

	"github.com/AdhityaRamadhanus/fasthttpcors"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

// Route ss
func Route() fasthttp.RequestHandler {

	withCors := fasthttpcors.NewCorsHandler(fasthttpcors.Options{
		AllowMaxAge: math.MaxInt32,
	})

	router := routing.New()
	db := database.InitDB()
	router.Use(func(c *routing.Context) error {
		c.Set("db", db)
		c.Response.Header.Set("Content-Type", "application/json")

		return c.Next()
	})

	api := router.Group("/api/")
	// user
	api.Post("user", user.Create)
	api.Get("user/<id>", user.GetByID)
	api.Get("user", user.GetAll)
	api.Put("user/<id>", user.Update)
	// restaurant
	api.Get("restaurant/<id>", restaurant.GetByID)
	api.Get("restaurant", restaurant.GetAll)
	api.Get("restaurant/<id>/products", restaurant.GetProducts)
	api.Get("restaurant/<id>/images", restaurant.GetImages)
	api.Get("restaurant/<id>/campaigns", restaurant.GetCampaigns)

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

	return withCors.CorsMiddleware(router.HandleRequest)
}
