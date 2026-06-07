package repository

import (
	"math"
	"math/rand/v2"
	"strconv"
)

func CreateSession() string {
	var id string = strconv.FormatInt(rand.Int64N(math.MaxInt64), 10)
	return id
}
