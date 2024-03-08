package handlers

import (
	"WanderVive_back/pkg/databases"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func LivehouseHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(databases.GetLivehouse()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := w.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}
