package utils

import (
	"crypto/md5"
	"fmt"
)

func Hash(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

func CheckHash(password string, hash string) bool {
	if Hash(password) != hash {
		return false
	}
	return true
}
