package main

import (
	"chatApp/handlers"
	"net/http"
)

// go: embed static/app

func main() {

	mux := http.NewServeMux()
	//if request is to website return the files
	mux.HandleFunc("/", handlers.Spa)
	//if request is to the api to send
	mux.HandleFunc("/api/v1/login", handlers.LogIn)
	mux.HandleFunc("/api/v1/auth", handlers.Auth)
	mux.HandleFunc("/api/v1/user", handlers.UserData)
	mux.HandleFunc("/api/v1/room", handlers.UserData)
	http.ListenAndServe(":80", corsMiddleware(mux))

}
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
