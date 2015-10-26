package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"omaha/util"
	"os"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", util.GetOmahaPath()+"/omaha.db")
	if err != nil {
		log.Fatal(err)
	}
	// check if user table exists
	var name string
	err = DB.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='account';").Scan(&name)
	switch {
	case err == sql.ErrNoRows:
		// table doesn't exist
		createAccountTable(DB)
	case err != nil:
		log.Fatal(err)
	}
}

func createAccountTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE  account (
		uid INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(20) NOT NULL,
		name VARCHAR(20),
		hash CHAR(60) NOT NULL
		);`) // yes, pw shouldn't be plain text
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created account table")
	}
}

func createDBFile() {
	file, err := os.Create(util.GetOmahaPath() + "/omaha.db")
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created DB file")
}
