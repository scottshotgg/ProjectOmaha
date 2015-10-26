package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"omaha/util"
)

func InitDB() {
	_, err := sql.Open("sqlite3", util.GetOmahaPath()+"/omaha.db")
	if err != nil {
		fmt.Println("DB Initialization Failed")
	} else {
		fmt.Println("DB Initialized")
	}

}
