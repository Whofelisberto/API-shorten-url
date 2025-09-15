package store

import (
	"math/rand"
	"time"
)
const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func init() {
	rand.Seed(time.Now().UnixNano())
}

func genCode() string {
	const n = 8
	bytes := make([]byte, n)

	for i := 0; i < n; i++ {
		bytes[i] = characters[rand.Intn(len(characters))]
	}
	return string(bytes)
}
