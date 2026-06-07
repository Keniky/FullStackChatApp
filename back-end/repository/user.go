package repository

import (
	"math"
	"math/rand/v2"
)

type User struct {
	Id     int64
	Pfp    string
	Read   chan string
	RoomId int64
	Name   string
}

func CreateUser() *User {
	return &User{
		Id:   rand.Int64N(math.MaxInt64),
		Read: make(chan string),
	}
}
