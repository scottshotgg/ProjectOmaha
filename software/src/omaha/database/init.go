package database

import (
	"database/sql"
	// import the sqlite implementation
	_ "github.com/mattn/go-sqlite3"
	"log"
	"omaha/util"
	"os"
)

// DB is the object used to access the database
var (
	DB                    *sql.DB
	insertSpeakerStmt     *sql.Stmt
	addSpeakerToZonesStmt *sql.Stmt
	addSpeakerToPagingZonesStmt *sql.Stmt
)

// InitDB creates the DB variable. If the database hasn't been created yet, it will be created.
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
	// user table doesn't exist
	case err == sql.ErrNoRows:
		// create all tables
		createAccountTable()
		createSpeakerTable()
		createZoneTable()
		createPagingZoneTable()
		createEqualizerPresetsTable()
		createMusicEqualizerPresetsTable()
		createTargetsTable()
		createZoneToSpeakerTable()
		createPagingZoneToSpeakerTable()
		prepareStatements()
		populateSpeakerTable()
	case err != nil:
		log.Fatal(err)
	}
}

func prepareStatements() {
	var err error
	insertSpeakerStmt, err = getInsertSpeakerStatement()
	if err != nil {
		log.Fatal(err)
	}
	addSpeakerToZonesStmt, err = getAddSpeakerToZonesStmt()
	if err != nil {
		log.Fatal(err)
	}
	addSpeakerToPagingZonesStmt, err = getAddSpeakerToPagingZonesStmt()
	if err != nil {
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
