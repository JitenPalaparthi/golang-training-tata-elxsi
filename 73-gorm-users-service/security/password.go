package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(plain string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	} else {
		return nil
	}
}

// package main

// import (
// "fmt"
// "golang.org/x/crypto/bcrypt"
// )

// func main() {
// password := "myS3cureP@ssword"

// // Hash the password with default cost
// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// if err != nil {
// fmt.Println("Error hashing password:", err)
// return
// }
// fmt.Println("Hashed Password:", string(hashedPassword))

// // Verify the password
// err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
// if err == nil {
// fmt.Println("Password match!")
// } else {
// fmt.Println("Invalid password.")
// }
// }
