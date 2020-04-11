package main

import (
	"fmt"
	"math"
	"os"
	"snappy-api/router"

	"github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/valyala/fasthttp"
)

func main() {

	withCors := fasthttpcors.NewCorsHandler(fasthttpcors.Options{
		AllowMaxAge: math.MaxInt32,
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Printf("server listen on :" + port)
	panic(fasthttp.ListenAndServe(":"+port, withCors.CorsMiddleware(router.Route())))

}
