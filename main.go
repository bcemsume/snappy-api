package main

import (
	"fmt"
	"os"
	"snappy-api/router"

	"github.com/valyala/fasthttp"
)

func main() {

	mapper.InitMapper()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Printf("server listen on :" + port)
	panic(fasthttp.ListenAndServe(":"+port, router.Route()))

}
