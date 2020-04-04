package router

import (
	"math"
	usercontrollers "snappy-api/controllers/user"
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

	api.Post("user", usercontrollers.Create)
	api.Put("user", usercontrollers.Update)
	api.Get("user/<id>", usercontrollers.GetByID)
	api.Get("user", usercontrollers.GetAll)
	api.Put("user/<id>", usercontrollers.Update)

	return withCors.CorsMiddleware(router.HandleRequest)
}
