package main

import (
	"log"
	"net/http"

	"github.com/cricarba/profilebackend/certificate"
)

const basePath = "/api"

func main() {
	certificate.SetupRoutes(basePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
