package main

import (
	"flag"
	"golang-serverless-and-server/hello"
	"net/http"

	"github.com/apex/gateway"
)

func main() {
	lambda := flag.Bool("lambda", false, "run as lambda")
	hello.RegisterRoutes()
	if *lambda {
		if err := gateway.ListenAndServe(":3000", nil); err != nil {
			panic(err)
		}
	} else {
		if err := http.ListenAndServe(":3000", nil); err != nil {
			panic(err)
		}
	}
}
