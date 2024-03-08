package handlers

import (
	"WanderVive_back/pkg/databases"
	"WanderVive_back/pkg/models"
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func BandHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=user01 dbname=wondervive_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	svc := databases.Service{
		DB: db,
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	res := models.BandResponse{
		Contents: svc.GetBand(),
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
