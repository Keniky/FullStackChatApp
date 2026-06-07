package handlers

import (
	"chatApp/repository"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var sessions = make(map[string]repository.User)

func Spa(w http.ResponseWriter, r *http.Request) {

	requestedData := filepath.Join("static/app", r.URL.Path)

	_, err := os.Stat(requestedData)

	//there is no file such that
	if err != nil {
		http.ServeFile(w, r, filepath.Join("static/app/index.html"))
		return
	}

	// if there is a file return it
	http.ServeFile(w, r, requestedData)

}

func Auth(w http.ResponseWriter, r *http.Request) {

	fmt.Println("authentication started")
	cookie, err := r.Cookie("session_id")

	fmt.Println("cookie is ", cookie)
	fmt.Println("user sessions are ", sessions)
	if err != nil {
		fmt.Println("authentication failed")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if _, ok := sessions[cookie.Value]; ok {
		fmt.Println("authentication succeded")
		w.WriteHeader(http.StatusAccepted)
		return
	}

	fmt.Println("authentication failed")
	w.WriteHeader(http.StatusNotAcceptable)

}

func UserData(w http.ResponseWriter, r *http.Request) {

	fmt.Println("user has just demanded his data")
	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user, ok := sessions[cookie.Value]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name":"` + user.Name + `"}`))
	fmt.Println("user has recieved his user name ", user.Name)
}

func LogIn(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	var jsonBody map[string]string
	json.Unmarshal(body, &jsonBody)

	newUser := repository.CreateUser()
	newUser.Name = jsonBody["name"]

	session_id := repository.CreateSession()

	sessions[session_id] = *newUser

	fmt.Println("new user has been added ", newUser.Name)

	//we got the name
	cookie := &http.Cookie{
		Name:  "session_id",
		Value: session_id,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
