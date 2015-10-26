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
