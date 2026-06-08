package repository

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand/v2"
)

type Room struct {
	Id         int64
	Join       chan *User
	Leave      chan *User
	Users      map[*User]bool
	NewMessage chan []byte
}

func CreateRoom() *Room {
	var roomId int64 = rand.Int64N(math.MaxInt64)

	var room *Room = &Room{
		Id:         roomId,
		Join:       make(chan *User),
		Leave:      make(chan *User),
		Users:      make(map[*User]bool),
		NewMessage: make(chan []byte),
	}
	go room.run()
	return room
}

func (r *Room) run() {
	fmt.Println(r.Id, " has started running")

	for {
		select {
		case user := <-r.Join:
			//add user to users
			r.Users[user] = true
		case user := <-r.Leave:
			//remove user from room
			delete(r.Users, user)
			close(user.Recieve)
		case msg := <-r.NewMessage:
			for user := range r.Users {
				user.Recieve <- msg
			}
		}
	}
}
func (r *Room) UsersToJSON() []byte {

	users := make([]*User, 0, len(r.Users))
	fmt.Println("getting users from ", r.Id)
	for user := range r.Users {
		users = append(users, user)
		fmt.Println(user.Id)
	}
	answer, _ := json.Marshal(users)
	return answer
}
