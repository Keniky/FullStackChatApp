package config

import "time"

const (
	PORT            = "5000"
	SESSIONDURATION = time.Hour * 24
	Dsn             = "postgres://eneusdev:idkman01@localhost:5432/chatappdb?sslmode=disable"
)
