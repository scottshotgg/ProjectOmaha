package database

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreateAccount(username, password, name string) {
	hashByte, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hash := string(hashByte)
	log.Println(hash)
	_, err := DB.Exec(`INSERT INTO account 
		(username, name, hash)
		VALUES (?, ?, ?)
		`, username, name, hash)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Created account %s", username)
	}
}

// func GetAccountByUsername(username string) omaha.Account {
// 	return nil
// }
