package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"omaha/util"
)

func InitDB() {

	db, err := sql.Open("sqlite3", util.GetOmahaPath()+"/omaha.db")
	defer db.Close()
	if err != nil {
		log.Fatalf("DB file doesn't exist\n")
	}
	// check if user table exists
	var name string
	err = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='account';").Scan(&name)
	switch {
	case err == sql.ErrNoRows:
		// table doesn't exist
		createAccountTable(db)
	case err != nil:
		log.Fatal(err)
	}
}

func createAccountTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE  account (
		uid INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(50) NOT NULL,
		name VARCHAR(50),
		password VARCHAR(20) NOT NULL
		);`) // yes, pw shouldn't be plain text
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created account table")
	}
}
