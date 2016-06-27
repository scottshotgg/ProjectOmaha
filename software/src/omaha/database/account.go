/*
	
	Account.go focuses on the account aspects of the software system; verifying logins, who can access which speakers and zones, account access levels, as well as creating accounts. 
	It also contains the authentication routines used. Currently there is not way to delete an account of change an account after making one.

*/
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

	// Take the username that was submitted and find the UID for that user.
	err = DB.QueryRow(`
		SELECT uid
		FROM account
		WHERE username=?`, username).Scan(&userID)
	
	log.Println(userID)

	// Create an account session based off the UID and the session hash.
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

func LogoutAccount(hash string) (error) {
	// Take the username that was submitted and find the UID for that user.
	_, err := DB.Exec(`
		DELETE
		FROM accountSession
		WHERE sessionKey=?`, hash)
	
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

/*
	GetLevelOfAccount retrieves the level of an account based on the username of an account and returns the level of the account in the form of an integer
	Return values: 0, 1, 2 
*/
func GetLevelOfAccount(username string) int {
	var level = 0
	err := DB.QueryRow(`SELECT level FROM account WHERE username=?`, username).Scan(&level)
	if(err != nil) {
		log.Println("error getting level of account in GetLevelOfAccount")
	}

	log.Println(level)
	return level
}

/*
	GetSpeakerForAccount checks which speaker it has access to and returns the speaker ID that it has access to.
*/
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
		log.Println("error getting speaker from account in GetSpeakerForAccount")
	}

	return speakerID
}

/*
	GetZoneForAccount checks which zones the account has access to and return the zone ID that it has access to.
*/
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
		log.Println("error getting zone from account in GetZoneForAccount")
	}

	return zoneID
}

/*
	CreateAccount inserts a row in the account table with the given configuration.
*/
func CreateAccount(level int, username string, password string, name string, email string, phone string, speakerID int, zoneID int) error {
	hashByte, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hash := string(hashByte)
	// Insert the values provided into the account table
	_, err := DB.Exec(`INSERT INTO account 
		(level, username, name, email, phone, hash)
		VALUES (?, ?, ?, ?, ?, ?)
		`, level, username, name, email, phone, hash)			// later on make the email a net/mail type and use the AddressParser to verify
	if err == nil {
		log.Printf("Created account %s\n", username)

	var uid int

	// Get the UID
	DB.QueryRow(`
		SELECT uid 
		FROM account
		where username=? 
		`, username).Scan(&uid)

	// If the speaker is provided, use the UID to link the account to a speaker
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

	// If the zone is provided, use the UID to link the account to a zone
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

/*
	getCreateAccountStmt is a private function used to create the account insertion statement.
*/
func getCreateAccountStmt() (*sql.Stmt, error) {
	return DB.Prepare(`INSERT INTO account 
		(username, name, hash)
		VALUES (?, ?, ?)
		`)
}

/*
	createAccountTable is a private function used to create the account table.
*/
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

/*
	createAccountToSpeakersTable creates the AccountToSpeakers table that is used to link the accounts to the speakers that they can control.
*/ 
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

/*
	createAccountToMaskingZonesTable creates the AccountToMaskingZones table that is used to link the accounts to the making zones that they can control. 
*/
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

/*
	creatAccountToPagingZonesTable creates the AccountToPagingZones table that is used to link the accounts to the paging zones that they can control.
*/
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

/* 
	IsSessionHashValid checks if the given hash is present in the database. 
	The return value is true if count is equal to 1, and false otherwise.
*/
func IsSessionHashValid(hash string) bool {
	var count int
	// Ideally this would have a time that it would expire at and you would check that time, but this has not been implemented yet so sessions are infinite.
	// Along with this time, the browser would need to put something like "...; expires=Tue, 21 Jun, 2016 12:00:00 UTC;" into the cookie that would correlate with the backend session time.
	err := DB.QueryRow(`
		SELECT count(1) from accountSession where sessionKey = ?;
		`, hash).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(count)
	if count == 1 {
		return true
	} else {
		return false
	}
}

/*
	AuthenticatePermissionFromHash takes a session hash string uses a nested select statment to retrieve the account access level based on the submitted hash by
	using the accountSession table to acquire a userID from the hash and then using that userID to then retrieve the access level from the account table.
*/
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
		log.Println("error getting account hash in AuthenticatePermissionFromHash")
	}
	log.Println(permission)
	return permission
}

/*
	AuthenticateSpeakerFromHash takes a session hash string and uses a nested select statment to retrieve the account access level based on the submitted hash by
	using the accountSeession table to acquire a userID from the hash and then using that userID to retrieve the speaker from the AccountToSpeaker table.
*/
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
		log.Println("error retrieving speaker in AuthenticateSpeakerFromHash")
	}

	log.Println(speaker)
	return speaker
}

/*
	AuthenticateZonesFromHash taks a session hash string and uses a nested select statment to retrieve the account access level based on the submitted hash by
	using the accountSession table to acquire a userID from the hash and then using then using the userID to retrieve the zone from the AccountsToZone table.
*/
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
		log.Println("error retrieving zone in AuthenticateZoneFromHash")
	}

	log.Println(zone)
	return zone
}

/*
	generateSessionHash is a private function that takes a string of the English alphabet capital and non-capital along with the first ten digits and creates a
	unique hash from based on the random function.
*/
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
