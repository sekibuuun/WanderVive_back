package main

import (
	"WanderVive_back/pkg/routes"
	"log"
	"net/http"
)

func main() {
	routes.Router()
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
