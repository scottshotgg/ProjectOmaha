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
/*
	LoginAccount checks the given information against what is in the account table.
	If the information is correct, a new session hash is returned.
*/
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

	var userID int

	err = DB.QueryRow(`
		SELECT uid
		FROM account
		WHERE username=?`, username).Scan(&userID)
	
	log.Println(userID)

	_, err = DB.Exec(`
		INSERT INTO accountSession
		(sessionKey, userID)
		VALUES (?, ?)
		`, sessionHash, userID)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	log.Printf("%s logged in\n", username)
	return sessionHash, nil
}

func GetLevelOfAccount(username string) int {
	var level = 0
	err := DB.QueryRow(`SELECT level FROM account WHERE username=?`, username).Scan(&level)
	if(err != nil) {
		log.Println("we got le error on GetLevelOfAccount")
	}

	log.Println(level)
	return level
}

func GetSpeakerForAccount(username string) int {
	var speakerID = -1
	err := DB.QueryRow(`
		SELECT speaker 
		FROM AccountToSpeakers 
		WHERE uid=(
			SELECT uid 
			FROM account 
			WHERE username=?)
			`, username).Scan(&speakerID)
	if(err != nil) {
		log.Println("we got le error on GetSpeakerForAccount")
	}

	log.Println("this is the speaker", speakerID)
	return speakerID
}

func GetZoneForAccount(username string) int {
	var zoneID = -1
	err := DB.QueryRow(`
		SELECT zone 
		FROM AccountToMaskingZones 
		WHERE uid=(
			SELECT uid 
			FROM account 
			WHERE username=?)
			`, username).Scan(&zoneID)
	if(err != nil) {
		log.Println("we got le error on GetZoneForAccount")
	}

	log.Println("this is the zone", zoneID)

	return zoneID
}

// CreateAccount inserts a row in the account table with the given configuration.
// later on make the email a net/mail type and use the AddressParser to verify
func CreateAccount(level int, username string, password string, name string, email string, phone string, speakerID int, zoneID int) error {
	hashByte, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hash := string(hashByte)
	_, err := DB.Exec(`INSERT INTO account 
		(level, username, name, email, phone, hash)
		VALUES (?, ?, ?, ?, ?, ?)
		`, level, username, name, email, phone, hash)
	if err == nil {
		log.Printf("Created account %s\n", username)

	var uid int

	DB.QueryRow(`
		SELECT uid 
		FROM account
		where username=? 
		`, username).Scan(&uid)

	if(speakerID > 0) {
		_, err := DB.Exec(`
			INSERT INTO AccountToSpeakers 
			VALUES (?, ?)
			`, uid, speakerID)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Created speaker link\n")
		}
	}

	if(zoneID > 0) {
		_, err := DB.Exec(`
			INSERT INTO AccountToMaskingZones 
			VALUES (?, ?)
			`, uid, zoneID)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Created zone link\n")
		}
	}

		return nil
	}
	switch err.(sqlite3.Error).Code {
	case sqlite3.ErrConstraint:
		// If a constraint is violated, this username is already in use
		return errors.New("An account with that username already exists")
	default:
		log.Fatal(err)
		return nil
	}

	return nil
}

func getCreateAccountStmt() (*sql.Stmt, error) {
	return DB.Prepare(`INSERT INTO account 
		(username, name, hash)
		VALUES (?, ?, ?)
		`)
}

func createAccountTable() {
	_, err := DB.Exec(`
		CREATE TABLE account (
			uid INTEGER PRIMARY KEY AUTOINCREMENT,
			level INTEGER NOT NULL,
			username VARCHAR(20) NOT NULL UNIQUE,
			name VARCHAR(40),
			email VARCHAR(40),
			phone VARCHAR(40),
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
		log.Println("Created table account")
	}
	
	CreateAccount(2, "super", "super", "superuser", "", "", -1, -1)
}
 
func createAccountToSpeakersTable() {
	_, err := DB.Exec(`
		CREATE TABLE  AccountToSpeakers (
			uid INTEGER REFERENCES account(uid),
			speaker INTEGER REFERENCES Speaker(SpeakerID),
			PRIMARY KEY (uid, speaker)
		);
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created table AccountToSpeakers")
	}
}

func createAccountToMaskingZonesTable() {
	_, err := DB.Exec(`
		CREATE TABLE  AccountToMaskingZones (
			uid INTEGER REFERENCES account(uid),
			zone INTEGER REFERENCES zone(zoneID),
			PRIMARY KEY (uid, zone)
		);
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created table AccountToMaskingZones")
	}
}

func createAccountToPagingZonesTable() {
	_, err := DB.Exec(`
		CREATE TABLE  AccountToPagingZones (
			uid INTEGER REFERENCES account(uid),
			zone INTEGER REFERENCES zone(zoneID),
			PRIMARY KEY (uid, zone)
		);
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created table AccountToPagingZones")
	}
}

// IsSessionHashValid checks if the given hash is present in the database
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

func AuthenticatePermissionFromHash(hash string) int {
	var permission = 0
	log.Println(hash)
	err := DB.QueryRow(`
		SELECT level 
		FROM account 
		WHERE uid=(
			SELECT userID 
			FROM accountSession 
			WHERE sessionkey=?)
		`, hash).Scan(&permission)
	if(err != nil) {
		log.Println("we got le error bro")
	}
	log.Println(permission)
	return permission
}

func AuthenticateSpeakerFromHash(hash string) int {
	var speaker = 0
	err := DB.QueryRow(`
		SELECT speaker
		FROM AccountToSpeakers
		WHERE uid=(
			SELECT uid
			FROM accountSession
			WHERE sessionKey=?)
		`, hash).Scan(&speaker)

	if(err != nil) {
		log.Println("got an error retrieving speaker")
	}

	log.Println(speaker)
	return speaker
}

func AuthenticateZoneFromHash(hash string) int {
	var zone = 0
	err := DB.QueryRow(`
		SELECT zone
		FROM AccountToMaskingZones
		WHERE uid=(
			SELECT uid
			FROM accountSession
			WHERE sessionKey=?)
		`, hash).Scan(&zone)

	if(err != nil) {
		log.Println("got an error retrieving zone")
	}

	log.Println(zone)
	return zone
}

func generateSessionHash() string {
	var chars = "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
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
