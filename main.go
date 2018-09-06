package main

import (
	"flag"
	"fmt"
	"golang-serverless-and-server/hello"
	"net/http"

	"github.com/apex/gateway"
)

func main() {
	server := flag.Bool("server", false, "run as server")
	flag.Parse()

	hello.RegisterRoutes()
	if *server {
		fmt.Println("Running as server on port 3000")
		if err := http.ListenAndServe(":3000", nil); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Running as lambda on port 3000")
		if err := gateway.ListenAndServe(":3000", nil); err != nil {
			panic(err)
		}
	}
}
