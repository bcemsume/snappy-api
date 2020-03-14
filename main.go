package main

import (
	"log"
	"math"
	"os"

	"github.com/AdhityaRamadhanus/fasthttpcors"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func main() {
	withCors := fasthttpcors.NewCorsHandler(fasthttpcors.Options{
		AllowMaxAge: math.MaxInt32,
	})

	router := routing.New()

	router.Use(func(c *routing.Context) error {
		c.Response.Header.Set("Content-Type", "application/json")
		return nil
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// api := router.Group("/api")

	// api.Post("/user", controller.Create)
	// api.Put("/user", controller.Update)
	// api.Delete("/user/<id>", controller.Delete)
	// api.Get("/user/<id>", controller.GetById)
	// api.Get("/user", controller.GetAll)

	log.Println("server listen " + port)
	panic(fasthttp.ListenAndServe(":"+port, withCors.CorsMiddleware(router.HandleRequest)))

}
