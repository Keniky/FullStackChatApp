package services

import (
	"chatApp/repository"
	"math"
	"math/rand/v2"
)

type Room struct {
	Users map[repository.User]bool

	Join chan repository.User

	Leave chan repository.User

	Message chan string

	Id int64
}

const (
	max = math.MaxInt64
)

func CreateRoom() (room *Room) {
	room = &Room{
		Users:   make(map[repository.User]bool),
		Join:    make(chan repository.User),
		Leave:   make(chan repository.User),
		Message: make(chan string),
		Id:      rand.Int64N(max),
	}

	return room
}

func (room *Room) Run() {
	for {
		select {
		case user := <-room.Join:
			room.Users[user] = true

		case user := <-room.Leave:
			delete(room.Users, user)
			close(user.Read)
		case msg := <-room.Message:
			for user := range room.Users {
				user.Read <- msg
			}
		}
	}
}
