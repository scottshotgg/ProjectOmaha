package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"omaha/util"
)

// createSpeakerTable creates the speaker table in the database
func createSpeakerTable() {
	_, err := DB.Exec(`
		CREATE TABLE speaker (
			speakerID INTEGER PRIMARY KEY,
			name varchar(50),
			y INTEGER,
			x INTEGER
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

// addSpeaker adds a speaker to the database at the specified location
func addSpeaker(loc speakerLocation) {
	// add to speaker table
	result, err := DB.Exec(`INSERT INTO speaker 
		(x, y)
		VALUES (?, ?)
		`, loc.X, loc.Y)
	if err != nil {
		log.Fatal(err)
	} else {
		id, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		// add to default zone
		addSpeakerToDefaultZone(int8(id))
		log.Printf("Created speaker %d at (%d, %d)\n", id,loc.X, loc.Y)
	}
}

// populateSpeaker table populates the speaker table with the locations from the speaker locations file
func populateSpeakerTable() {
	speakerLocations := getSpeakerLocations()
	for _, speakerLoc := range speakerLocations {
		addSpeaker(speakerLoc)
	}
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
