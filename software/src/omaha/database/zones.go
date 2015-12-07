package database

import (
	"database/sql"
	"log"
)

// createZoneTable creates the zone table in the database
func createZoneTable() {
	_, err := DB.Exec(`
		CREATE TABLE zone (
			zoneID INTEGER PRIMARY KEY AUTOINCREMENT,
			name varchar(50)	
		)
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created zone table")
		AddZone("default")
	}
}

// createZoneToSpeakerTable creates the zoneToSpeaker table in the database
func createZoneToSpeakerTable() {
	_, err := DB.Exec(`
		CREATE TABLE zoneToSpeaker (
			zoneID INTEGER REFERENCES Zone(ZoneID),
			speakerID INTEGER REFERENCES Speaker(SpeakerID),
			PRIMARY KEY (ZoneID, SpeakerID),
			FOREIGN KEY(zoneID) REFERENCES zone,
			FOREIGN KEY(speakerID) REFERENCES speaker
		)
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created zoneToSpeaker table")
	}
}

// AddZone adds a zone to the database with the specified name
func AddZone(name string) {
	_, err := DB.Exec(`INSERT INTO zone 
		(name)
		VALUES (?)
		`, name)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Created zone %s\n", name)
	}
}

// GetZoneID gets the ID of the zone with the given name
func getZoneID(zoneName string) (int8, error) {
	var zoneID int8
	DB.QueryRow(`
		SELECT zoneID 
		FROM zone
		WHERE name=?
		`, zoneName).Scan(&zoneID)
	return zoneID, nil
}

func getAddSpeakerToZonesStmt() (*sql.Stmt, error) {
	stmt, err := DB.Prepare(`INSERT INTO zoneToSpeaker 
		(zoneID, speakerID)
		VALUES (?, ?)
		`)
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

func addSpeakerToDefaultZone(speakerID int8) {
	zoneID, _ := getZoneID("default")
	_, err := addSpeakerToZonesStmt.Exec(zoneID, speakerID)
	if err != nil {
		log.Fatal(err)
	}
}

// SetSpeakerToZoneByName moves the given speaker to the zone with the given name
func SetSpeakerToZoneByName(speaker *ControllerStatus, zoneName string) {
	zoneID, _ := getZoneID(zoneName)
	SetSpeakerToZoneByID(speaker, zoneID)
}

// SetSpeakerToZoneByID moves the given speaker to the zone with the given name
func SetSpeakerToZoneByID(speaker *ControllerStatus, zoneID int8) {
	// update zoneToSpeaker table
	_, err := DB.Exec(`UPDATE zoneToSpeaker 
		SET zoneID = ?
		WHERE speakerID = ?
		`, zoneID, speaker.ID)
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllZones returns all zones in the database
func GetAllZones() []*Zone {
	zoneIDs := getAllZoneIDs()
	zones := []*Zone{}
	for _, zoneID := range zoneIDs {
		zone := GetZone(zoneID)
		zones = append(zones, zone)
	}
	return zones
}

// getAllZoneIDs gets all zone ID of every zone in the database
func getAllZoneIDs() []int8 {
	rows, err := DB.Query(`
		SELECT zoneID
		FROM zone
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	zoneIDs := []int8{}
	for rows.Next() {
		var zoneID int8
		rows.Scan(&zoneID)
		zoneIDs = append(zoneIDs, zoneID)
	}
	return zoneIDs
}

// getZonesSpeakers gets the speakers that belong to the specified zone
func getZonesSpeakers(zoneID int8) []*ControllerStatus {
	rows, err := DB.Query(`
		SELECT s.speakerID, x, y
		FROM speaker s
		INNER JOIN zoneToSpeaker z
		ON s.speakerID = z.speakerId
		WHERE zoneID = ?
	`, zoneID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	speakers := []*ControllerStatus{}
	for rows.Next() {
		var speakerID int8
		var x int
		var y int
		rows.Scan(&speakerID, &x, &y)

		speaker := &ControllerStatus{}
		speaker.X = x
		speaker.Y = y
		speaker.ID = int8(speakerID)
		speakers = append(speakers, speaker)
	}
	return speakers
}

// GetZone gets the Zone with the specified ID from the database
func GetZone(zoneID int8) *Zone {
	// get speakers
	speakers := getZonesSpeakers(zoneID)
	// get name
	var name string
	DB.QueryRow(`
		SELECT name 
		FROM zone
		WHERE zoneID=?
		`, zoneID).Scan(&name)

	zone := &Zone{ID: zoneID, Name: name, Speakers: speakers}
	return zone
}
