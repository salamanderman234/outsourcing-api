package helpers

import (
	"crypto/rand"
	"fmt"
)

type stringHelper struct{}

func (stringHelper) GenerateRandomString(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return s
}

var String = stringHelper{}
