package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func NewUserID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic("can't generate new user")
	}

	return hex.EncodeToString(b)
}
