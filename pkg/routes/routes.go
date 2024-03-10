package routes

import (
	"WanderVive_back/pkg/handlers"
	"fmt"
	"log"
	"net/http"
)

func Router(mux *http.ServeMux) {
	http.HandleFunc("/band", handlers.BandHandler)
	http.HandleFunc("/event", handlers.EventHandler)
	http.HandleFunc("/livehouse", handlers.LivehouseHandler)
	http.HandleFunc("/nearbyEvent", handlers.NearbyEventHandler)
	fmt.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
