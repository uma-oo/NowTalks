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

func CheckPasswordHash(password, hash string) bool {
	fmt.Println("start comparing")
	start := time.Now()
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	elapsed := time.Since(start)
	fmt.Println("bcrypt comparison took:", elapsed)
    fmt.Println("length", len(hash))
	fmt.Println("done comparing")
	fmt.Println("err inside the hash", err)
	return err == nil
}
