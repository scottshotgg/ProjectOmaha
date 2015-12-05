package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"omaha/util"
	"database/sql"
)

// createSpeakerTable creates the speaker table in the database
func createSpeakerTable() {
	_, err := DB.Exec(`
		CREATE TABLE speaker (
			speakerID INTEGER PRIMARY KEY,
			name varchar(50),
			y INTEGER,
			x INTEGER,
			volumeLevel INTEGER
		)
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created speaker table")
	}
}

// GetAllSpeakers gets all speakers from the database and returns them as a slice of ControllerStatus objects
func GetAllSpeakers() []*ControllerStatus {
	rows, err := DB.Query(`
		SELECT speakerID, x, y
		FROM speaker
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	speakers := []*ControllerStatus{}
	for rows.Next() {
		var speakerID int
		var x int
		var y int
		if err := rows.Scan(&speakerID, &x, &y); err != nil {
			log.Fatal(err)
		}
		speaker := &ControllerStatus{}
		speaker.X = x
		speaker.Y = y
		speaker.ID = int8(speakerID)
		speakers = append(speakers, speaker)
	}
	return speakers
}

// GetSpeaker gets a speaker from the database
func GetSpeaker(speakerID int8) *ControllerStatus {
	var volumeLevel int8
	var x int
	var y int
	err := DB.QueryRow(`
		SELECT x, y, volumeLevel
		FROM speaker
		WHERE speakerID=?;
		`, speakerID).Scan(&x, &y, &volumeLevel)
	if err != nil {
		log.Fatal(err)
	}
	speaker := &ControllerStatus{X: x, Y: y, VolumeLevel: volumeLevel, ID: speakerID}
	return speaker
}

// SaveSpeaker saves the speaker provided to the database
func SaveSpeaker(speaker *ControllerStatus) {
	_, err := DB.Exec(`
		UPDATE speaker
		SET
			volumeLevel = ?
		WHERE speakerID = ?
	`, speaker.VolumeLevel, speaker.ID)
	if err != nil {
		log.Fatal(err)
	}
}

// addSpeaker adds a speaker to the database at the specified location
func addSpeaker(loc speakerLocation) int8 {
	// add to speaker table
	result, err := insertSpeakerStmt.Exec(loc.X, loc.Y)
	if err != nil {
		log.Fatal(err)
	} else {
		id, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		// add to default zone
		//addSpeakerToDefaultZone(int8(id), zoneStmt2)
		log.Printf("Created speaker %d at (%d, %d)\n", id,loc.X, loc.Y)
		return int8(id)
	}
	return 0
}

func getInsertSpeakerStatement() (*sql.Stmt, error) {
	stmt, err := DB.Prepare(`INSERT INTO speaker 
		(x, y, volumeLevel)
		VALUES (?, ?, 0)
		`)
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

// populateSpeaker table populates the speaker table with the locations from the speaker locations file
func populateSpeakerTable() {
	tx, err := DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// reset the statements back to their original values
	defer func(speakerStmt, zoneStmt *sql.Stmt){
		insertSpeakerStmt = speakerStmt
		addSpeakerToZonesStmt = zoneStmt
	}(insertSpeakerStmt, addSpeakerToZonesStmt)
	// change the statements to use this transaction
	insertSpeakerStmt = tx.Stmt(insertSpeakerStmt)
	addSpeakerToZonesStmt = tx.Stmt(addSpeakerToZonesStmt)
	speakerLocations := getSpeakerLocations()
	for _, speakerLoc := range speakerLocations {
		id := addSpeaker(speakerLoc)
		addSpeakerToDefaultZone(id)
	}
	tx.Commit()
}

type speakerLocation struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type speakerLocationsJSON struct {
	SpeakerLocations []speakerLocation `json:"circleLocations"`
}

// getSpeakerLocations parses the speaker locations file and returns a list of speakerLocations
func getSpeakerLocations() []speakerLocation {
	rawJSON, _ := ioutil.ReadFile(util.GetOmahaPath() + "/templates/css/speakerLocations.txt")
	speakerLocs := &speakerLocationsJSON{}
	json.Unmarshal(rawJSON, speakerLocs)
	return speakerLocs.SpeakerLocations
}
