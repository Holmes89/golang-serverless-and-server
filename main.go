package main

import (
	"flag"
	"golang-serverless-and-server/hello"
	"net/http"

	"github.com/apex/gateway"
)

func main() {
	server := flag.Bool("server", false, "run as server")
	hello.RegisterRoutes()
	if *server {
		if err := http.ListenAndServe(":3000", nil); err != nil {
			panic(err)
		}
	} else {
		if err := gateway.ListenAndServe(":3000", nil); err != nil {
			panic(err)
		}
	}
}
