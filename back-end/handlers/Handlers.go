package handlers

import (
	"chatApp/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/websocket"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: messageBufferSize}

// sessions
var sessions = make(map[string]*repository.User)

// rooms
var rooms = make(map[string]*repository.Room)

// serve website files
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

// verify if user is logged int
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

func RoomHandlerCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user who should be authentificated is requesting to create a room ")
	var currentRoom *repository.Room = repository.CreateRoom()
	var currentRoomId string = strconv.FormatInt(currentRoom.Id, 10)

	rooms[currentRoomId] = currentRoom

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(`{"room_id":"` + currentRoomId + `"}`))
	fmt.Println("room has been served ")
}

// return user data
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

// create user
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

	sessions[session_id] = newUser

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

func JoinRoom(w http.ResponseWriter, r *http.Request) {
	//user should be authorized or else they will be redirected to the main room
	//without it they cant join no user lol
	cookie, err := r.Cookie("session_id")

	if err != nil {
		fmt.Println("no cookie user failed to join room")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user, ok := sessions[cookie.Value]

	if !ok {
		fmt.Println("user not available fake cookie")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	//now get the room id from user
	roomId := r.URL.Query().Get("room_id")
	//now we get the room id

	if roomId == "" {
		fmt.Println("user tried sneak and put an empty id lol  ")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	//check if there is a room
	currentRoom, ok := rooms[roomId]

	if !ok {
		fmt.Println("room not available fake code room is ", roomId)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// whenever somebody logs into the room create a websocket with them
	// add user to room
	// let new users write their messages
	// and let users read what is in the channel
	//user added to room

	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		//bad kills system
		// log.Fatal("while serving http : ", err)
		log.Println("while serving http : ", err)
		return
	}

	user.Socket = socket
	user.URoom = currentRoom
	currentRoom.Join <- user

	//when user disconnects just kick him out of the room lol
	defer func() { currentRoom.Leave <- user }()

	//let user check his email and write messages in his window
	go user.Write()

	//read messages by user and send them to room
	user.Read()

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	//no need to authentify
	//now get the room id from user
	roomId := r.URL.Query().Get("room_id")
	//now we get the room id

	if roomId == "" {
		fmt.Println("user tried sneak and put an empty id lol  ")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	//check if there is a room
	currentRoom, ok := rooms[roomId]

	if !ok {
		fmt.Println("room not available fake code room is ", roomId)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	//if there is room
	w.Header().Set("Content-Type", "application/json")
	w.Write(currentRoom.UsersToJSON())

}
