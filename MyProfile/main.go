package main

import (
	"log"
	"net/http"
	"ggithub.com/cricarba/goexercices/certificate"
)

const basePath = "/api"

func main() {
	certificate.SetupRoutes(basePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
