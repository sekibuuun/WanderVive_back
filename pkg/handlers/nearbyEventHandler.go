package handlers

import (
	"WanderVive_back/pkg/databases"
	"WanderVive_back/pkg/models"
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

func filterDate(events []models.EventAndLivehouse, date string) []models.EventAndLivehouse {
	var res []models.EventAndLivehouse
	for _, e := range events {
		if e.EventDate == date {
			res = append(res, e)
		}
	}
	return res
}

func convertRad(degree float64) float64 {
	rad := (degree / 180) * math.Pi

	return rad
}

func haversineDistance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	radius := 6371000
	dlat := lat2 - lat1
	dlon := lon2 - lon1
	temp := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	constant := 2 * math.Asin(math.Sqrt(temp))
	return float64(radius) * constant
}

func isNear(currentCoordinate [2]float64, livehouseCoordinate [2]float64, maxDist float64) bool {
	return maxDist >= haversineDistance(convertRad(currentCoordinate[1]), convertRad(currentCoordinate[0]), convertRad(livehouseCoordinate[0]), convertRad(livehouseCoordinate[1]))
}

func filterNear(events []models.EventAndLivehouse, date string, currentCoordinate [2]float64, maxDist float64) []models.EventAndLivehouse {
	res := make([]models.EventAndLivehouse, 0)
	for _, event := range events {
		livehouseCoordinate := [2]float64{event.Longitude, event.Latitude}
		if isNear(currentCoordinate, livehouseCoordinate, maxDist) {
			res = append(res, event)
		}
	}
	return res
}

func NearbyEventHandler(w http.ResponseWriter, r *http.Request) {
	strLongitude := r.FormValue("longitude")
	strLatitude := r.FormValue("latitude")
	longitude, err := strconv.ParseFloat(strLongitude, 64)
	if err != nil {
		log.Println(err)
		return
	}
	latitude, err := strconv.ParseFloat(strLatitude, 64)
	if err != nil {
		log.Println(err)
		return
	}
	currentCoordinate := [2]float64{longitude, latitude}
	date := r.FormValue("date")
	strMaxDist := r.FormValue("maxDist")
	maxDist, err := strconv.ParseFloat(strMaxDist, 64)
	if err != nil {
		log.Println(err)
		return
	}
	db, err := sql.Open("postgres", "host=localhost port=5432 user=user01 dbname=wondervive_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	svc := databases.Service{
		DB: db,
	}
	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)
	events := svc.GetEventAndLivehouse()
	events = filterDate(events, date)
	res := models.NearbyEventResponse{
		Contents: filterNear(events, date, currentCoordinate, maxDist),
	}
	if err := enc.Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}
