package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"omaha/util"
	"os"
)

var (
	DB                    			*sql.DB
	insertSpeakerStmt     			*sql.Stmt
	addSpeakerToZonesStmt 			*sql.Stmt
	addSpeakerToPagingZonesStmt *sql.Stmt
	addSpeakerToSchedulingStmt	*sql.Stmt
	addZoneToSchedulingStmt			*sql.Stmt
)

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", util.GetOmahaPath()+"/omaha.db")
	if err != nil {
		log.Fatal(err)
	}

	var name string
	err = DB.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='account';").Scan(&name)
	switch {
		case err == sql.ErrNoRows:

			// Create DB tables
			createAccountToSpeakersTable()
			createAccountToMaskingZonesTable()
			createAccountToPagingZonesTable()
			createAccountTable()
			createSpeakerTable()
		
			createZoneTable()
			createPagingZoneTable()

			createTargetsTable()
			createZoneToSpeakerTable()
			createPagingZoneToSpeakerTable()
			
			createEqualizerPresetsTable()
			createMusicEqualizerPresetsTable()
			createPagingEqualizerPresetsTable()


			createTargetsTableZone()
			createEqualizerPresetsTableZone()
			createMusicEqualizerPresetsTableZone()
			createPagingEqualizerPresetsTableZone()

			createSchedulingTable()
			createSchedulingTableZone()

			prepareStatements()
			populateZoneTable()
			populateSpeakerTable()

			CreateControllerTable()

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
	addSpeakerToSchedulingStmt, err = getAddSpeakerToSchedulingStmt()
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
