package database

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"omaha/util"
)

func createSpeakerTable(db *sql.DB) {
	_, err := db.Exec(`
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
		var speakerId int
		var x int
		var y int
		if err := rows.Scan(&speakerId, &x, &y); err != nil {
			log.Fatal(err)
		}
		speaker := &ControllerStatus{}
		speaker.X = x
		speaker.Y = y
		speaker.ID = int8(speakerId)
		speakers = append(speakers, speaker)
	}
	return speakers
}

func addSpeaker(loc speakerLocation) {
	_, err := DB.Exec(`INSERT INTO speaker 
		(x, y)
		VALUES (?, ?)
		`, loc.X, loc.Y)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Created speaker at (%d, %d)\n", loc.X, loc.Y)
	}
}

func populateSpeakerTable() {
	speakerLocations := getSpeakerLocations()
	for _, speakerLoc := range speakerLocations {
		addSpeaker(speakerLoc)
	}
}

type speakerLocationJson struct {
	SpeakerLocations []speakerLocation `json:"circleLocations"`
}

type speakerLocation struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func getSpeakerLocations() []speakerLocation {
	rawJson, _ := ioutil.ReadFile(util.GetOmahaPath() + "/templates/css/speakerLocations.txt")
	speakerLocs := &speakerLocationJson{}
	json.Unmarshal(rawJson, speakerLocs)
	return speakerLocs.SpeakerLocations
}
