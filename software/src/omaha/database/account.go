package database

import (
	"database/sql"
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

func createAccountTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE  account (
			uid INTEGER PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(20) NOT NULL UNIQUE,
			name VARCHAR(20),
			hash CHAR(60) NOT NULL
		);

		CREATE INDEX usernameIndex ON account(username);`) // yes, pw shouldn't be plain text
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created account table")
	}
}
