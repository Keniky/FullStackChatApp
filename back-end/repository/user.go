package repository

import (
	"chatApp/database"
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

// store user in the database
func SaveUserInDB(name, url string) bool {
	conn := database.GetDBConnection()

	transaction, err := conn.Begin()

	if err != nil {
		return false
	}

	var id int
	err = transaction.QueryRow(
		"INSERT INTO profiles(profile_url) VALUES($1) RETURNING profile_id",
		url,
	).Scan(&id)

	if err != nil {
		fmt.Println("failed to insert new profile")
		return false
	}

	//now insert user it self

	_, err = transaction.Exec(
		"INSERT INTO users(user_name , profile_id) VALUES($1, $2)",
		name,
		id,
	)

	if err != nil {
		fmt.Println("failed to insert new user")
		return false
	}
	transaction.Commit()

	return true
}

type UserNameAndPfp struct {
	Name string
	Pfp  string
}

func VerifyIfUserIsSignedUp(name string) (userData *UserNameAndPfp) {
	conn := database.GetDBConnection()
	rows, err := conn.Query("SELECT users.user_name, profiles.profile_url "+
		"FROM users "+
		"JOIN profiles ON users.profile_id = profiles.profile_id "+
		"WHERE users.user_name = $1 ",
		name,
	)
	if err != nil {
		fmt.Println("failed to retrieve data from database")
		return nil
	}

	defer rows.Close()

	if rows.Next() {
		var temp UserNameAndPfp
		err = rows.Scan(&temp.Name, &temp.Pfp)
		if err != nil {
			fmt.Println("error while reading data for user")
			return nil
		}
		return &temp

	} else {
		return nil
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
