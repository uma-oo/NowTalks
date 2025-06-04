package utils

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// edit the cost to become 10 instead of 14
func CheckPasswordHash(password, hash string) bool {
	start := time.Now()
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	elapsed := time.Since(start)
	fmt.Println("bcrypt comparison took:", elapsed)
	return err == nil
}
