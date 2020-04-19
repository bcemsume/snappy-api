package main

import (
	"fmt"
	"math"
	"os"
	"snappy-api/models"
	"snappy-api/router"

	"github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/valyala/fasthttp"
)

func main() {

	withCors := fasthttpcors.NewCorsHandler(fasthttpcors.Options{
		AllowMaxAge:    math.MaxInt32,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Printf("server listen on :" + port)
	server := &fasthttp.Server{
		Name:         "snappy-api",
		Handler:      withCors.CorsMiddleware(router.Route()),
		ErrorHandler: errorHandler,
	}
	panic(server.ListenAndServe(":" + port))
}

func errorHandler(ctx *fasthttp.RequestCtx, err error) {
	fmt.Printf("error requested resource: %v %v\n", string(ctx.Method()), string(ctx.Path()))
	fmt.Println("error handler", err.Error())
	res := models.NewResponse(false, nil, "error")

	ctx.Write(res.MustMarshal())
}
