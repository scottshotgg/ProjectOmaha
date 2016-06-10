package database

import (
	"database/sql"
	"log"
)

func createZoneTable() {
	_, err := DB.Exec(`
		CREATE TABLE zone (
			zoneID INTEGER PRIMARY KEY AUTOINCREMENT,
			name varchar(50),

			volumeLevel INTEGER DEFAULT 0,
			musicLevel INTEGER DEFAULT 0,
			pagingLevel INTEGER DEFAULT 0,
			soundMaskingLevel INTEGER DEFAULT 0,

			lowerMusicThreshold INTEGER DEFAULT 0,
			lowerPagingThreshold INTEGER DEFAULT 0,
			lowerMaskingThreshold INTEGER DEFAULT 0,

			upperMusicThreshold INTEGER DEFAULT 100,
			upperPagingThreshold INTEGER DEFAULT 100,
			upperMaskingThreshold INTEGER DEFAULT 100,

			fadeTime INTEGER,
			fadeLevel INTEGER,
			effectiveness INTEGER,
			pleasantness INTEGER,
			status INTEGER DEFAULT 0,
			equalizerMode INTEGER DEFAULT 0,
			schedulingMode INTEGER DEFAULT 0,

			eBand0 REAL, 		
			eBand1 REAL, 
			eBand2 REAL, 
			eBand3 REAL, 
			eBand4 REAL, 
			eBand5 REAL, 
			eBand6 REAL, 
			eBand7 REAL, 
			eBand8 REAL, 
			eBand9 REAL, 
			eBand10 REAL, 
			eBand11 REAL, 
			eBand12 REAL, 
			eBand13 REAL, 
			eBand14 REAL, 
			eBand15 REAL, 
			eBand16 REAL, 
			eBand17 REAL, 
			eBand18 REAL, 
			eBand19 REAL, 
			eBand20 REAL,

			mBand0 REAL, 
			mBand1 REAL, 
			mBand2 REAL, 
			mBand3 REAL, 
			mBand4 REAL, 
			mBand5 REAL, 
			mBand6 REAL, 
			mBand7 REAL, 
			mBand8 REAL, 
			mBand9 REAL, 
			mBand10 REAL, 
			mBand11 REAL, 
			mBand12 REAL, 
			mBand13 REAL, 
			mBand14 REAL, 
			mBand15 REAL, 
			mBand16 REAL, 
			mBand17 REAL, 
			mBand18 REAL, 
			mBand19 REAL, 
			mBand20 REAL,

			pBand0 REAL, 
			pBand1 REAL, 
			pBand2 REAL, 
			pBand3 REAL, 
			pBand4 REAL, 
			pBand5 REAL, 
			pBand6 REAL, 
			pBand7 REAL, 
			pBand8 REAL, 
			pBand9 REAL, 
			pBand10 REAL, 
			pBand11 REAL, 
			pBand12 REAL, 
			pBand13 REAL, 
			pBand14 REAL, 
			pBand15 REAL, 
			pBand16 REAL, 
			pBand17 REAL, 
			pBand18 REAL, 
			pBand19 REAL, 
			pBand20 REAL,
			
			tBand0 REAL, 
			tBand1 REAL, 
			tBand2 REAL, 
			tBand3 REAL, 
			tBand4 REAL, 
			tBand5 REAL, 
			tBand6 REAL, 
			tBand7 REAL, 
			tBand8 REAL, 
			tBand9 REAL, 
			tBand10 REAL, 
			tBand11 REAL, 
			tBand12 REAL, 
			tBand13 REAL, 
			tBand14 REAL, 
			tBand15 REAL, 
			tBand16 REAL, 
			tBand17 REAL, 
			tBand18 REAL, 
			tBand19 REAL, 
			tBand20 REAL
		)
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created zone table")
	}
}


func createEqualizerPresetsTableZone() {
	_, err := DB.Exec(`
		CREATE TABLE EqualizerPresetsZone (
			zoneID INTEGER,
			name varchar(50),
			whichPreset INTEGER,
			band0 REAL,
			band1 REAL,
			band2 REAL,
			band3 REAL,
			band4 REAL,
			band5 REAL,
			band6 REAL,
			band7 REAL,
			band8 REAL,
			band9 REAL,
			band10 REAL,
			band11 REAL,
			band12 REAL,
			band13 REAL,
			band14 REAL,
			band15 REAL,
			band16 REAL,
			band17 REAL,
			band18 REAL,
			band19 REAL,
			band20 REAL,
			PRIMARY KEY (zoneID, name)
		)
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created equalizerZone table")
	}
}

func createMusicEqualizerPresetsTableZone() {
	_, err := DB.Exec(`
		CREATE TABLE MusicEqualizerPresetsZone (
			zoneID INTEGER,
			name varchar(50),
			whichPreset INTEGER,
			band0 REAL,
			band1 REAL,
			band2 REAL,
			band3 REAL,
			band4 REAL,
			band5 REAL,
			band6 REAL,
			band7 REAL,
			band8 REAL,
			band9 REAL,
			band10 REAL,
			band11 REAL,
			band12 REAL,
			band13 REAL,
			band14 REAL,
			band15 REAL,
			band16 REAL,
			band17 REAL,
			band18 REAL,
			band19 REAL,
			band20 REAL,
			PRIMARY KEY (zoneID, name)
		)
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created music equalizerZone table")
	}
}

func createPagingEqualizerPresetsTableZone() {
	_, err := DB.Exec(`
		CREATE TABLE PagingEqualizerPresetsZone (
			zoneID INTEGER,
			name varchar(50),
			whichPreset INTEGER,
			band0 REAL,
			band1 REAL,
			band2 REAL,
			band3 REAL,
			band4 REAL,
			band5 REAL,
			band6 REAL,
			band7 REAL,
			band8 REAL,
			band9 REAL,
			band10 REAL,
			band11 REAL,
			band12 REAL,
			band13 REAL,
			band14 REAL,
			band15 REAL,
			band16 REAL,
			band17 REAL,
			band18 REAL,
			band19 REAL,
			band20 REAL,
			PRIMARY KEY (zoneID, name)
		)
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created paging equalizerZone table")
	}
}

func createTargetsTableZone() {
	_, err := DB.Exec(`
		CREATE TABLE TargetsZone (
			zoneID INTEGER,
			name varchar(50),
			whichPreset INTEGER,
			band0 REAL,
			band1 REAL,
			band2 REAL,
			band3 REAL,
			band4 REAL,
			band5 REAL,
			band6 REAL,
			band7 REAL,
			band8 REAL,
			band9 REAL,
			band10 REAL,
			band11 REAL,
			band12 REAL,
			band13 REAL,
			band14 REAL,
			band15 REAL,
			band16 REAL,
			band17 REAL,
			band18 REAL,
			band19 REAL,
			band20 REAL,
			PRIMARY KEY (zoneID, name)
		)
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created targetsZone table")
	}
}

func createSchedulingTableZone() {
	_, err := DB.Exec(`
		CREATE TABLE SchedulingZone (
			zoneID INTEGER,
			day INT,
			interval INT,
			start INT,
			end INT,
			hour0 INT,
			hour1 INT,
			hour2 INT,
			hour3 INT,
			hour4 INT,
			hour5 INT,
			hour6 INT,
			hour7 INT,
			hour8 INT,
			hour9 INT,
			hour10 INT,
			hour11 INT,
			hour12 INT,
			hour13 INT,
			hour14 INT,
			hour15 INT,
			hour16 INT,
			hour17 INT,
			hour18 INT,
			hour19 INT,
			hour20 INT,
			hour21 INT,
			hour22 INT,
			hour23 INT
		)
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created schedulingZone table")
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

// createPagingZoneTable creates the zone table in the database
func createPagingZoneTable() {
	_, err := DB.Exec(`
		CREATE TABLE pagingZone (
			zoneID INTEGER PRIMARY KEY AUTOINCREMENT,
			name varchar(50)	
		)
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created pagingZone table")
		AddPagingZone("all_paging")
	}
}

// createPagingZoneToSpeakerTable creates the zoneToSpeaker table in the database
func createPagingZoneToSpeakerTable() {
	_, err := DB.Exec(`
		CREATE TABLE pagingZoneToSpeaker (
			zoneID INTEGER REFERENCES pagingZone(ZoneID),
			speakerID INTEGER REFERENCES Speaker(SpeakerID),
			PRIMARY KEY (ZoneID, SpeakerID),
			FOREIGN KEY(zoneID) REFERENCES pagingZone,
			FOREIGN KEY(speakerID) REFERENCES speaker
		)
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created pagingZoneToSpeaker table")
	}
}

// AddZone adds a zone to the database with the specified name
func AddZone(name string) {
	result, err := DB.Exec(`INSERT INTO zone 
			(name,
			volumeLevel,
			musicLevel,
			pagingLevel,
			soundMaskingLevel,

			lowerMusicThreshold,
			lowerPagingThreshold,
			lowerMaskingThreshold,

			upperMusicThreshold,
			upperPagingThreshold,
			upperMaskingThreshold,

			fadeTime,
			fadeLevel,
			effectiveness,
			pleasantness,
			status,
			equalizerMode,

			eBand0, 		
			eBand1, 
			eBand2, 
			eBand3, 
			eBand4, 
			eBand5, 
			eBand6, 
			eBand7, 
			eBand8, 
			eBand9, 
			eBand10, 
			eBand11, 
			eBand12, 
			eBand13, 
			eBand14, 
			eBand15, 
			eBand16, 
			eBand17, 
			eBand18, 
			eBand19, 
			eBand20,

			mBand0, 
			mBand1, 
			mBand2, 
			mBand3, 
			mBand4, 
			mBand5, 
			mBand6, 
			mBand7, 
			mBand8, 
			mBand9, 
			mBand10, 
			mBand11, 
			mBand12, 
			mBand13, 
			mBand14, 
			mBand15, 
			mBand16, 
			mBand17, 
			mBand18, 
			mBand19, 
			mBand20,

			pBand0, 
			pBand1, 
			pBand2, 
			pBand3, 
			pBand4, 
			pBand5, 
			pBand6, 
			pBand7, 
			pBand8, 
			pBand9, 
			pBand10, 
			pBand11, 
			pBand12, 
			pBand13, 
			pBand14, 
			pBand15, 
			pBand16, 
			pBand17, 
			pBand18, 
			pBand19, 
			pBand20,
			
			tBand0, 
			tBand1, 
			tBand2, 
			tBand3, 
			tBand4, 
			tBand5, 
			tBand6, 
			tBand7, 
			tBand8, 
			tBand9, 
			tBand10, 
			tBand11, 
			tBand12, 
			tBand13, 
			tBand14, 
			tBand15, 
			tBand16, 
			tBand17, 
			tBand18, 
			tBand19, 
			tBand20)
		VALUES (?, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		`, name)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Created zone %s\n", name)
	}
	id, err := result.LastInsertId()
	for day := 0; day < 7; day++ {
		_, err = DB.Exec(`INSERT into SchedulingZone 
			VALUES (?, ?, -1, -1, -1, 100, 100, 100, 100, 100, 
						100, 100, 100, 100, 100, 100, 100, 100, 100, 
						100, 100, 100, 100, 100, 100, 100, 100, 100, 100)
					`, id, day)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Added day %s\n", day, "to \"all_masking\"")
		}
	}

}



// AddPagingZone adds a paging zone to the database with the specified name
func AddPagingZone(name string) {
	_, err := DB.Exec(`INSERT INTO pagingZone 
		(name)
		VALUES (?)
		`, name)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Created paging zone %s\n", name)
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

// getPagingZoneID gets the ID of the zone with the given name
func getPagingZoneID(zoneName string) (int8, error) {
	var zoneID int8
	DB.QueryRow(`
		SELECT ZoneID 
		FROM pagingZone
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

func getAddSpeakerToPagingZonesStmt() (*sql.Stmt, error) {
	stmt, err := DB.Prepare(`INSERT INTO pagingZoneToSpeaker 
		(zoneID, speakerID)
		VALUES (?, ?)
		`)
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

func getAddSpeakerToSchedulingStmt() (*sql.Stmt, error) {
	stmt, err := DB.Prepare(`INSERT INTO Scheduling
		VALUES (?, ?, -1, -1, -1, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100)
		`)
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

func addSpeakerToAllZones(speakerID int8) {
	zoneID, _ := getZoneID("all_masking")
	_, err := addSpeakerToZonesStmt.Exec(zoneID, speakerID)
	if err != nil {
		log.Fatal(err)
	}

	pagingZoneID, _ := getPagingZoneID("all_paging")
	_, errr := addSpeakerToPagingZonesStmt.Exec(pagingZoneID, speakerID)
	if errr != nil {
		log.Fatal(errr)
	}
}

func addSpeakerSchedule(speakerID int8) {
	for day := 0; day < 7; day++ {
		_, err := addSpeakerToSchedulingStmt.Exec(speakerID, day)
		if err != nil {
			log.Fatal(err)
		}
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


func GetAllPagingZones() []*Zone {
	zoneIDs := getAllPagingZoneIDs()
	zones := []*Zone{}
	for _, zoneID := range zoneIDs {
		zone := GetPagingZone(zoneID)
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

func getAllPagingZoneIDs() []int8 {
	rows, err := DB.Query(`
		SELECT zoneID
		FROM pagingZone
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
	log.Println(zoneIDs)
	return zoneIDs
}

// getZonesSpeakers gets the speakers that belong to the specified zone
func getZonesSpeakers(zoneID int8) []*ControllerStatus {
	rows, err := DB.Query(`
		SELECT s.speakerID, x, y, status
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
		var status int
		rows.Scan(&speakerID, &x, &y, &status)

		speaker := &ControllerStatus{}
		speaker.X = x
		speaker.Y = y
		speaker.ID = int8(speakerID)
		speaker.Status = status
		speakers = append(speakers, speaker)
	}
	return speakers
}

// getZonesSpeakers gets the speakers that belong to the specified zone
func getPagingZonesSpeakers(zoneID int8) []*ControllerStatus {
	rows, err := DB.Query(`
		SELECT s.speakerID, x, y, status
		FROM speaker s
		INNER JOIN pagingZoneToSpeaker z
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
		var status int
		rows.Scan(&speakerID, &x, &y, &status)

		speaker := &ControllerStatus{}
		speaker.X = x
		speaker.Y = y
		speaker.ID = int8(speakerID)
		speaker.Status = status
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

	var volumeLevel int8
	var musicLevel int8
	var pagingLevel int8
	var soundMaskingLevel int8
	var lowerMusicThreshold int8
	var lowerPagingThreshold int8
	var lowerMaskingThreshold int8
	var upperMusicThreshold int8
	var upperPagingThreshold int8
	var upperMaskingThreshold int8
	var fadeTime int8
	var fadeLevel int8
	var effectiveness int8
	var pleasantness int8
	var presetName string
	var targetName string
	var whichPreset int
	var status int
	var equalizerMode int8
	var schedulingMode int8

	var band0 float64
	var band1 float64
	var band2 float64
	var band3 float64
	var band4 float64
	var band5 float64
	var band6 float64
	var band7 float64
	var band8 float64
	var band9 float64
	var band10 float64
	var band11 float64
	var band12 float64
	var band13 float64
	var band14 float64
	var band15 float64
	var band16 float64
	var band17 float64
	var band18 float64
	var band19 float64
	var band20 float64

	var eBand0 float64
	var eBand1 float64
	var eBand2 float64
	var eBand3 float64
	var eBand4 float64
	var eBand5 float64
	var eBand6 float64
	var eBand7 float64
	var eBand8 float64
	var eBand9 float64
	var eBand10 float64
	var eBand11 float64
	var eBand12 float64
	var eBand13 float64
	var eBand14 float64
	var eBand15 float64
	var eBand16 float64
	var eBand17 float64
	var eBand18 float64
	var eBand19 float64
	var eBand20 float64

	var mBand0 float64
	var mBand1 float64
	var mBand2 float64
	var mBand3 float64
	var mBand4 float64
	var mBand5 float64
	var mBand6 float64
	var mBand7 float64
	var mBand8 float64
	var mBand9 float64
	var mBand10 float64
	var mBand11 float64
	var mBand12 float64
	var mBand13 float64
	var mBand14 float64
	var mBand15 float64
	var mBand16 float64
	var mBand17 float64
	var mBand18 float64
	var mBand19 float64
	var mBand20 float64

	var pBand0 float64
	var pBand1 float64
	var pBand2 float64
	var pBand3 float64
	var pBand4 float64
	var pBand5 float64
	var pBand6 float64
	var pBand7 float64
	var pBand8 float64
	var pBand9 float64
	var pBand10 float64
	var pBand11 float64
	var pBand12 float64
	var pBand13 float64
	var pBand14 float64
	var pBand15 float64
	var pBand16 float64
	var pBand17 float64
	var pBand18 float64
	var pBand19 float64
	var pBand20 float64

	var tBand0 float64
	var tBand1 float64
	var tBand2 float64
	var tBand3 float64
	var tBand4 float64
	var tBand5 float64
	var tBand6 float64
	var tBand7 float64
	var tBand8 float64
	var tBand9 float64
	var tBand10 float64
	var tBand11 float64
	var tBand12 float64
	var tBand13 float64
	var tBand14 float64
	var tBand15 float64
	var tBand16 float64
	var tBand17 float64
	var tBand18 float64
	var tBand19 float64
	var tBand20 float64


	var day int
	var interval int
	var start int
	var end int
	var hour0 int
	var hour1 int
	var hour2 int
	var hour3 int
	var hour4 int
	var hour5 int
	var hour6 int
	var hour7 int
	var hour8 int
	var hour9 int
	var hour10 int
	var hour11 int
	var hour12 int
	var hour13 int
	var hour14 int
	var hour15 int
	var hour16 int
	var hour17 int
	var hour18 int
	var hour19 int
	var hour20 int
	var hour21 int
	var hour22 int
	var hour23 int

	err := DB.QueryRow(`
		SELECT name, volumeLevel,
			musicLevel,
			pagingLevel,
			soundMaskingLevel,

			lowerMusicThreshold,
			lowerPagingThreshold,
			lowerMaskingThreshold,

			upperMusicThreshold,
			upperPagingThreshold,
			upperMaskingThreshold,

			fadeTime,
			fadeLevel,
			effectiveness,
			pleasantness,
			status,
			equalizerMode,
			schedulingMode,

			eBand0, 		
			eBand1, 
			eBand2, 
			eBand3, 
			eBand4, 
			eBand5, 
			eBand6, 
			eBand7, 
			eBand8, 
			eBand9, 
			eBand10, 
			eBand11, 
			eBand12, 
			eBand13, 
			eBand14, 
			eBand15, 
			eBand16, 
			eBand17, 
			eBand18, 
			eBand19, 
			eBand20,

			mBand0, 
			mBand1, 
			mBand2, 
			mBand3, 
			mBand4, 
			mBand5, 
			mBand6, 
			mBand7, 
			mBand8, 
			mBand9, 
			mBand10, 
			mBand11, 
			mBand12, 
			mBand13, 
			mBand14, 
			mBand15, 
			mBand16, 
			mBand17, 
			mBand18, 
			mBand19, 
			mBand20,

			pBand0, 
			pBand1, 
			pBand2, 
			pBand3, 
			pBand4, 
			pBand5, 
			pBand6, 
			pBand7, 
			pBand8, 
			pBand9, 
			pBand10, 
			pBand11, 
			pBand12, 
			pBand13, 
			pBand14, 
			pBand15, 
			pBand16, 
			pBand17, 
			pBand18, 
			pBand19, 
			pBand20,
			
			tBand0, 
			tBand1, 
			tBand2, 
			tBand3, 
			tBand4, 
			tBand5, 
			tBand6, 
			tBand7, 
			tBand8, 
			tBand9, 
			tBand10, 
			tBand11, 
			tBand12, 
			tBand13, 
			tBand14, 
			tBand15, 
			tBand16, 
			tBand17, 
			tBand18, 
			tBand19, 
			tBand20

		FROM zone
		WHERE zoneID=?;
		`, zoneID).Scan(
			&name,

			&volumeLevel,
			&musicLevel,
			&pagingLevel,
			&soundMaskingLevel,
			&lowerMusicThreshold,
			&lowerPagingThreshold,
			&lowerMaskingThreshold,
			&upperMusicThreshold,
			&upperPagingThreshold,
			&upperMaskingThreshold,
			&fadeTime,
			&fadeLevel,
			&effectiveness,
			&pleasantness,
			&status,
			&equalizerMode,
			&schedulingMode,

			// Current equalizer bands
			&eBand0,
			&eBand1,
			&eBand2,
			&eBand3,
			&eBand4,
			&eBand5,
			&eBand6,
			&eBand7,
			&eBand8,
			&eBand9,
			&eBand10,
			&eBand11,
			&eBand12,
			&eBand13,
			&eBand14,
			&eBand15,
			&eBand16,
			&eBand17,
			&eBand18,
			&eBand19,
			&eBand20,

			// Current music preset bands
			&mBand0,
			&mBand1,
			&mBand2,
			&mBand3,
			&mBand4,
			&mBand5,
			&mBand6,
			&mBand7,
			&mBand8,
			&mBand9,
			&mBand10,
			&mBand11,
			&mBand12,
			&mBand13,
			&mBand14,
			&mBand15,
			&mBand16,
			&mBand17,
			&mBand18,
			&mBand19,
			&mBand20,

			// Current paging bands
			&pBand0,
			&pBand1,
			&pBand2,
			&pBand3,
			&pBand4,
			&pBand5,
			&pBand6,
			&pBand7,
			&pBand8,
			&pBand9,
			&pBand10,
			&pBand11,
			&pBand12,
			&pBand13,
			&pBand14,
			&pBand15,
			&pBand16,
			&pBand17,
			&pBand18,
			&pBand19,
			&pBand20,

			// Current target bands
			&tBand0,
			&tBand1,
			&tBand2,
			&tBand3,
			&tBand4,
			&tBand5,
			&tBand6,
			&tBand7,
			&tBand8,
			&tBand9,
			&tBand10,
			&tBand11,
			&tBand12,
			&tBand13,
			&tBand14,
			&tBand15,
			&tBand16,
			&tBand17,
			&tBand18,
			&tBand19,
			&tBand20)

	if err != nil {
		log.Fatal(err)
	}
	zone := &Zone { 
		Name: name,
		Speakers: speakers,

		VolumeLevel: [4] int8 {
			volumeLevel, 
			musicLevel, 
			pagingLevel, 
			soundMaskingLevel }, 

		LowerThreshold: [3] int8 {
			lowerMusicThreshold,
			lowerPagingThreshold,
			lowerMaskingThreshold },		

		UpperThreshold: [3] int8 {
			upperMusicThreshold,
			upperPagingThreshold,
			upperMaskingThreshold	},
 
		PagingLevel: [2] int8 {
			fadeTime, 
			fadeLevel }, 

		Effectiveness: effectiveness,
		Pleasantness: pleasantness, 
		ZoneID: zoneID, 
		Status: status,
		EqualizerMode: equalizerMode,
		SchedulingMode: schedulingMode,

		CurrentPreset: [21] float64 {
			eBand0, eBand1, eBand2, eBand3, eBand4, eBand5, eBand6, eBand7, eBand8, eBand9, eBand10, eBand11, 
			eBand12, eBand13, eBand14, eBand15, eBand16, eBand17, eBand18, eBand19, eBand20 },

		CurrentMusicPreset: [21] float64 {
			mBand0, mBand1, mBand2, mBand3, mBand4, mBand5, mBand6, mBand7, mBand8, mBand9, mBand10, mBand11, 
			mBand12, mBand13, mBand14, mBand15, mBand16, mBand17, mBand18, mBand19, mBand20 },

		CurrentPagingPreset: [21] float64 {
			pBand0, pBand1, pBand2, pBand3, pBand4, pBand5, pBand6, pBand7, pBand8, pBand9, pBand10, pBand11, 
			pBand12, pBand13, pBand14, pBand15, pBand16, pBand17, pBand18, pBand19, pBand20 },

		CurrentTarget: [21] float64 {
			tBand0, tBand1, tBand2, tBand3, tBand4, tBand5, tBand6, tBand7, tBand8, tBand9, tBand10, tBand11, 
			tBand12, tBand13, tBand14, tBand15, tBand16, tBand17, tBand18, tBand19, tBand20 }}


	rows, err := DB.Query(`
	SELECT 	name, whichPreset, 
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
	FROM EqualizerPresetsZone
	WHERE zoneID=?;
	`, zoneID)

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&presetName, &whichPreset, 
				&band0,
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
			
			zone.PresetNames = append(zone.PresetNames, presetName)
			zone.Equalizer = append(zone.Equalizer, []float64 {band0, band1, band2, band3, band4, band5, band6, band7, band8, band9, band10, band11, band12, band13, band14, band15, band16, band17, band18, band19, band20})
			log.Println(zone.PresetNames, zone.Equalizer)		// make sure that the two arrays are the same size
	}

	rows, err = DB.Query(`
	SELECT 	name, whichPreset, 
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
	FROM MusicEqualizerPresetsZone
	WHERE zoneID=?;
	`, zoneID)

	for rows.Next() {

		err = rows.Scan(&presetName, &whichPreset, 
				&band0,
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
			
			zone.MusicPresetNames = append(zone.MusicPresetNames, presetName)
			zone.MusicEqualizer = append(zone.MusicEqualizer, []float64 {band0, band1, band2, band3, band4, band5, band6, band7, band8, band9, band10, band11, band12, band13, band14, band15, band16, band17, band18, band19, band20})
			log.Println(zone.MusicPresetNames, zone.MusicEqualizer)		// make sure that the two arrays are the same size
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	rows, err = DB.Query(`
	SELECT 	name, whichPreset, 
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
	FROM PagingEqualizerPresetsZone
	WHERE zoneID=?;
	`, zoneID)

	for rows.Next() {

		err = rows.Scan(&presetName, &whichPreset, 
				&band0,
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
			
			zone.PagingPresetNames = append(zone.PagingPresetNames, presetName)
			zone.PagingEqualizer = append(zone.PagingEqualizer, []float64 {band0, band1, band2, band3, band4, band5, band6, band7, band8, band9, band10, band11, band12, band13, band14, band15, band16, band17, band18, band19, band20})
			log.Println(zone.PagingPresetNames, zone.PagingEqualizer)		// make sure that the two arrays are the same size
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	rows, errr := DB.Query(`
	SELECT name, whichPreset, 
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
	FROM TargetsZone
	WHERE zoneID=?;
	`, zoneID)

	for rows.Next() {
		err = rows.Scan(
			&targetName, 
			&whichPreset, 
			&band0,
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

		if errr != nil {
			log.Fatal(err)
		}

		zone.TargetNames = append(zone.TargetNames, targetName)
		zone.Target = append(zone.Target, []float64 {band0, band1, band2, band3, band4, band5, band6, band7, band8, band9, band10, band11, band12, band13, band14, band15, band16, band17, band18, band19, band20})
		log.Println(zone.TargetNames, zone.Target)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}


	rows, errr = DB.Query(`
		SELECT * 
		FROM SchedulingZone
		WHERE zoneID = ?
		`, zoneID)

	for rows.Next() {
		err = rows.Scan(
			&zoneID,
			&day,
			&interval,
			&start,
			&end,
			&hour0,
			&hour1,
			&hour2,
			&hour3,
			&hour4,
			&hour5,
			&hour6,
			&hour7,
			&hour8,
			&hour9,
			&hour10,
			&hour11,
			&hour12,
			&hour13,
			&hour14,
			&hour15,
			&hour16,
			&hour17,
			&hour18,
			&hour19,
			&hour20,
			&hour21,
			&hour22,
			&hour23)

		if err != nil {
			log.Fatal(err)
		}

		zone.Schedule = append(zone.Schedule, []int {day, interval, start, end, hour0, hour1, hour2, hour3, hour4, hour5, hour6, hour7, hour8, hour9, hour10, hour11, hour12, hour13, hour14, hour15, hour16, hour17, hour18, hour19, hour20, hour21, hour22, hour23})
	}

	return zone
}

// GetZone gets the Zone with the specified ID from the database
func GetPagingZone(zoneID int8) *Zone {
	speakers := getPagingZonesSpeakers(zoneID)
	var name string
	DB.QueryRow(`
		SELECT name 
		FROM pagingZone
		WHERE zoneID=?
		`, zoneID).Scan(&name)

	zone := &Zone{ZoneID: zoneID, Name: name, Speakers: speakers}
	
	return zone
}
