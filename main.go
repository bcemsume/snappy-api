package main

import (
	"os"
	"snappy-api/router"

	"github.com/valyala/fasthttp"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	panic(fasthttp.ListenAndServe(":"+port, router.Route()))

}
