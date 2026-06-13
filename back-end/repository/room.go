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
	NewMembers chan bool
}

func CreateRoom() *Room {
	var roomId int64 = rand.Int64N(math.MaxInt64)

	var room *Room = &Room{
		Id:         roomId,
		Join:       make(chan *User),
		Leave:      make(chan *User),
		Users:      make(map[*User]bool),
		NewMessage: make(chan []byte),
		NewMembers: make(chan bool, 1),
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
			r.NewMembers <- true
		case user := <-r.Leave:
			//remove user from room
			r.NewMembers <- true
			delete(r.Users, user)
			close(user.Recieve)
		case msg := <-r.NewMessage:
			for user := range r.Users {
				user.Recieve <- msg
			}
		}
	}
}

type jsonResponse struct {
	Members []*User `json:"members"`
}

func (r *Room) UsersToJSON() []byte {

	users := make([]*User, 0, len(r.Users))
	fmt.Println("getting users from ", r.Id)
	for user := range r.Users {
		users = append(users, user)
	}
	answer, err := json.Marshal(jsonResponse{Members: users})
	if err != nil {
		fmt.Println("UsersToJSON marshal failed:", err)
		return []byte(`{"members":[]}`)
	}
	return answer
}

func (r *Room) AlertUsers() {
	for msg := range r.NewMembers {
		for user := range r.Users {
			fmt.Println("alerting ", user.Name)
			user.NewMember <- msg
		}
	}
}
