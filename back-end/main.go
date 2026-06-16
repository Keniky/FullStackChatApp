package main

import (
	"chatApp/database"
	"chatApp/handlers"
	"net/http"

	_ "github.com/lib/pq"
)

// go: embed static/app

func main() {
	conn := database.GetDBConnection()
	defer conn.Close()

	mux := http.NewServeMux()
	//serve static files directly
	mux.Handle("/static/app/", http.StripPrefix("/static/app/", http.FileServer(http.Dir("./static/app"))))
	//if request is to website return the files
	mux.HandleFunc("/", handlers.Spa)
	//if request is to the api to send

	//api to log in and create new user
	mux.HandleFunc("/api/v1/login", handlers.LogIn)

	//handler to make sure user is logged in
	mux.HandleFunc("/api/v1/auth", handlers.Auth)

	//handler to return user data GET
	mux.HandleFunc("/api/v1/user", handlers.UserData)

	//handler to create new room and return room id
	mux.HandleFunc("/api/v1/room", handlers.CreateRoom)

	//handler to join a room
	mux.HandleFunc("/api/v1/chat", handlers.JoinRoom)
	//chat users GET METHOD
	mux.HandleFunc("/api/v1/chat/members", handlers.GetUsers)

	//new SignIn api
	mux.HandleFunc("POST /api/v2/signin", handlers.NewSignIn)
	//new LogIn api
	mux.HandleFunc("POST /api/v2/login", handlers.NewLogIn)

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
