package repository

import (
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/gorilla/websocket"
)

type User struct {
	Id      int64       `json:"id"`
	Name    string      `json:"name"`
	Pfp     string      `json:"pfp"`
	Recieve chan []byte `json:"-"`
	URoom   *Room       `json:"-"`

	//each user a socket
	Socket *websocket.Conn `json:"-"`

	NewMember    chan bool       `json:"-"`
	MemberSocket *websocket.Conn `json:"-"`
}

func CreateUser() *User {
	return &User{
		Id:        rand.Int64N(math.MaxInt64),
		Recieve:   make(chan []byte),
		NewMember: make(chan bool),
	}
}

func (u *User) Read() {

	defer u.Socket.Close()

	for {

		_, msg, err := u.Socket.ReadMessage()
		if err != nil {
			fmt.Println("failed to read message in user ", u.Id, "user name is ", u.Name)
			return
		}
		//get user messages and write then in room
		u.URoom.NewMessage <- msg
	}
}

func (u *User) Write() {
	defer u.Socket.Close()

	//get room messages and write them in the socket
	for msg := range u.Recieve {
		u.Socket.WriteMessage(websocket.TextMessage, msg)
	}

}
func (u *User) RunMemberDetector() {

	for msg := range u.NewMember {
		fmt.Println(msg, "for user ", u.Name)
		u.MemberSocket.WriteMessage(websocket.TextMessage, u.URoom.UsersToJSON())
	}
}
