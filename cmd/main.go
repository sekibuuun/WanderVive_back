package main

import (
	// "WanderVive_back/pkg/routes"
	"WanderVive_back/pkg/handlers"
	"fmt"
	"log"
	"net/http"
)

// CORSを許可するラッパー関数
func allowCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ここでCORSに必要なヘッダーを設定
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Preflightリクエストに対する応答
		if r.Method == "OPTIONS" {
			return
		}

		// 次のハンドラーまたはミドルウェアに処理を移行
		next.ServeHTTP(w, r)
	})
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, CORS-enabled world!")
}

func main() {
	mux := http.NewServeMux()
	handler := allowCORS(mux)
	mux.Handle("/", allowCORS(http.HandlerFunc(mainHandler)))
	mux.HandleFunc("/band", handlers.BandHandler)
	mux.HandleFunc("/event", handlers.EventHandler)
	mux.HandleFunc("/livehouse", handlers.LivehouseHandler)
	mux.HandleFunc("/nearbyEvent", handlers.NearbyEventHandler)
	fmt.Println("Server Start Up........")
	// routes.Router(mux)

	log.Fatal(http.ListenAndServe("localhost:8080", handler))
}
