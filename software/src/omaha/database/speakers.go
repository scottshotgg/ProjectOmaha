package database

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"omaha/util"
	"strconv"
)

// createSpeakerTable creates the speaker table in the database
func createSpeakerTable() {
	_, err := DB.Exec(`
		CREATE TABLE speaker (
			speakerID INTEGER PRIMARY KEY,
			name varchar(50),
			y INTEGER,
			x INTEGER,
			volumeLevel INTEGER,
			band0 INTEGER,
			band1 INTEGER,
			band2 INTEGER,
			band3 INTEGER,
			band4 INTEGER,
			band5 INTEGER,
			band6 INTEGER,
			band7 INTEGER,
			band8 INTEGER,
			band9 INTEGER,
			band10 INTEGER,
			band11 INTEGER,
			band12 INTEGER,
			band13 INTEGER,
			band14 INTEGER,
			band15 INTEGER,
			band16 INTEGER,
			band17 INTEGER,
			band18 INTEGER,
			band19 INTEGER,
			band20 INTEGER
		)
	`) // needs to be an equalizer thing in here
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created speaker table")
	}
}

// GetAllSpeakers gets all speakers from the database and returns them as a slice of ControllerStatus objects
func GetAllSpeakers() []*ControllerStatus {
	rows, err := DB.Query(`
		SELECT speakerID, x, y, 
			volumeLevel,
			band0,
			band1,
			band2,
			band3,
			band4,
			band5,
			band6,
			band7,
			band8,
			band9,
			band10,
			band11,
			band12,
			band13,
			band14,
			band15,
			band16,
			band17,
			band18,
			band19,			
			band20

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
		var volumeLevel int
		var band0 int
		var band1 int
		var band2 int
		var band3 int
		var band4 int
		var band5 int
		var band6 int
		var band7 int
		var band8 int
		var band9 int
		var band10 int
		var band11 int
		var band12 int
		var band13 int
		var band14 int
		var band15 int
		var band16 int
		var band17 int
		var band18 int
		var band19 int
		var band20 int

		if err := rows.Scan(&speakerID, &x, &y, &volumeLevel, &band0,
			&band1,
			&band2,
			&band3,
			&band4,
			&band5,
			&band6,
			&band7,
			&band8,
			&band9,
			&band10,
			&band11,
			&band12,
			&band13,
			&band14,
			&band15,
			&band16,
			&band17,
			&band18,
			&band19,
			&band20)
		err != nil {
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
	var band0 int
	var band1 int
	var band2 int
	var band3 int
	var band4 int
	var band5 int
	var band6 int
	var band7 int
	var band8 int
	var band9 int
	var band10 int
	var band11 int
	var band12 int
	var band13 int
	var band14 int
	var band15 int
	var band16 int
	var band17 int
	var band18 int
	var band19 int
	var band20 int

	err := DB.QueryRow(`
		SELECT x, y, volumeLevel, 
			band0,
			band1,
			band2,
			band3,
			band4,
			band5,
			band6,
			band7,
			band8,
			band9,
			band10,
			band11,
			band12,
			band13,
			band14,
			band15,
			band16,
			band17,
			band18,
			band19,
			band20
		FROM speaker
		WHERE speakerID=?;
		`, speakerID).Scan(&x, &y, &volumeLevel, &band0,
			&band1,
			&band2,
			&band3,
			&band4,
			&band5,
			&band6,
			&band7,
			&band8,
			&band9,
			&band10,
			&band11,
			&band12,
			&band13,
			&band14,
			&band15,
			&band16,
			&band17,
			&band18,
			&band19,
			&band20)
	if err != nil {
		log.Fatal(err)
	}
	speaker := &ControllerStatus{X: x, Y: y, VolumeLevel: volumeLevel, ID: speakerID, Equalizer: [21]int{band0, band1, band2, band3, band4, band5, band6, band7, band8, band9, band10, band11, band12, band13, band14, band15, band16, band17, band18, band19, band20}}
	return speaker
}

// SaveSpeaker saves the speaker provided to the database
// make this a general thing that will insert some stuff, hack for now 
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

// make a command to save the band

func SaveBand(speaker *ControllerStatus, band int, level int) {

	var stringOfStatement string = "UPDATE speaker SET band" + strconv.Itoa(band) + " = ? WHERE speakerID = ?"

	//log.Println(stringOfStatement)

	//statement, err := DB.Prepare(stringOfStatement)

	_, errr := DB.Exec(stringOfStatement, level, speaker.ID)		// may not even need the speaker object to be passed in, idk why this shit is uneccessarily abstracted
	if errr != nil {				// also remember that the level that we are passing in is not offset, it is the true level, as it should be
		log.Fatal(errr)			
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
		log.Printf("Created speaker %d at (%d, %d)\n", id, loc.X, loc.Y)
		return int8(id)
	}
	return 0
}

func getInsertSpeakerStatement() (*sql.Stmt, error) {
	stmt, err := DB.Prepare(`INSERT INTO speaker 
		(x, y, volumeLevel, 
			band0,
			band1,
			band2,
			band3,
			band4,
			band5,
			band6,
			band7,
			band8,
			band9,
			band10,
			band11,
			band12,
			band13,
			band14,
			band15,
			band16,
			band17,
			band18,
			band19,			
			band20)
		VALUES (?, ?, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
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
	defer func(speakerStmt, zoneStmt *sql.Stmt) {
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
