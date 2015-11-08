package database

import (
	"crypto/rand"
	"database/sql"
	"errors"
	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/big"
)

func LoginAccount(username, password string) (string, error) {
	var accountHash string
	err := DB.QueryRow(`
		SELECT hash 
		FROM account
		WHERE username=?;
		`, username).Scan(&accountHash)
	switch {
	case err == sql.ErrNoRows:
		return "", errors.New("Invalid account name")
	}
	err = bcrypt.CompareHashAndPassword([]byte(accountHash), []byte(password))
	if err != nil {
		return "", errors.New("Invalid password")
	}
	sessionHash := generateSessionHash()
	_, err = DB.Exec(`
		INSERT INTO accountSession
		(sessionKey, userID)
		VALUES (?, ?)
		`, sessionHash, 0)
	if err != nil {
		log.Fatal(err)
		return "", err
	} else {
		log.Printf("%s logged in\n", username)
		return sessionHash, nil
	}
}

func CreateAccount(username, password, name string) error {
	hashByte, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hash := string(hashByte)
	_, err := DB.Exec(`INSERT INTO account 
		(username, name, hash)
		VALUES (?, ?, ?)
		`, username, name, hash)
	if err == nil {
		log.Printf("Created account %s\n", username)
		return nil
	}
	switch err.(sqlite3.Error).Code {
	case 19:
		log.Println("Account already exists")
		return errors.New("An account with that username already exists")
	default:
		log.Fatal(err)
		return nil
	}
	return nil
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

		CREATE INDEX usernameIndex ON account(username);

		CREATE TABLE accountSession (
			sessionKey text primary key,
			userID int not null
		);
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created account table")
	}

	CreateAccount("admin", "password", "admin")
	CreateAccount("andy", "andy", "andy")
	CreateAccount("danny", "danny", "danny")
	CreateAccount("bruce", "bruce", "bruce")
	CreateAccount("scott", "scott", "scott")
	CreateAccount("eric", "eric", "eric")
}

func IsSessionHashValid(hash string) bool {
	var count int
	err := DB.QueryRow(`
		SELECT count(1) from accountSession where sessionKey = ?;
		`, hash).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count == 1
}

func generateSessionHash() string {
	var chars string = "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsSize := big.NewInt(int64(len(chars)))
	hashSize := 30
	hash := make([]byte, hashSize)
	for i := 0; i < hashSize; i++ {
		index, err := rand.Int(rand.Reader, charsSize)
		if err != nil {
			log.Fatal(err)
		}
		hash[i] = chars[index.Int64()]
	}
	hashString := string(hash)
	return hashString
}
