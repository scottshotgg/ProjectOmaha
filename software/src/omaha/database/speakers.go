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
// tab this shit over if you get bored
// make one of these for the presets also
func createSpeakerTable() {
	_, err := DB.Exec(`
		CREATE TABLE speaker (
			speakerID INTEGER PRIMARY KEY,
			name varchar(50),
			y INTEGER,
			x INTEGER,
			volumeLevel INTEGER,
			musicLevel INTEGER,
			pagingLevel INTEGER,
			soundMaskingLevel INTEGER,
			fadeTime INTEGER,
			fadeLevel INTEGER,
			averagingMode INTEGER,
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
		SELECT speakerID, name, x, y,  
			volumeLevel,
			musicLevel,
			pagingLevel,
			soundMaskingLevel,
			fadeTime,
			fadeLevel,
			averagingMode INTEGER,
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
		var name string
		var volumeLevel int8
		var musicLevel int8
		var pagingLevel int8
		var soundMaskingLevel int8
		var fadeTime int8
		var fadeLevel int8
		var averagingMode int8
		// var currentMode int
		// var whichPreset int
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

		// in here get the presets that go along with the speakers and use those variables 

		if err := rows.Scan(&speakerID, &name, &x, &y,
			&volumeLevel, 
			&musicLevel,
			&pagingLevel,
			&soundMaskingLevel,
			&fadeTime,
			&fadeLevel,
			&averagingMode,
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
		err != nil {
			log.Fatal(err)
		}
		speaker := &ControllerStatus{}
		speaker.X = x
		speaker.Y = y
		speaker.ID = int8(speakerID)
		speaker.Name = name

		speaker.VolumeLevel[0] = int8(volumeLevel)
		speaker.VolumeLevel[1] = int8(musicLevel)
		speaker.VolumeLevel[2] = int8(pagingLevel)
		speaker.VolumeLevel[3] = int8(soundMaskingLevel)

		speaker.PagingLevel[0] = int8(fadeTime)
		speaker.PagingLevel[1] = int8(fadeLevel)
		// might need an averaging setter here later
		speakers = append(speakers, speaker)
	}
	return speakers
}

// GetSpeaker gets a speaker from the database
func GetSpeaker(speakerID int8) *ControllerStatus {
	var volumeLevel int8
	var musicLevel int8
	var pagingLevel int8
	var soundMaskingLevel int8
	var fadeTime int8
	var fadeLevel int8
	var averagingMode int8
	var x int
	var y int
	var name string
	var presetName string		// might need to make a catch for this
	var whichPreset int
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
		SELECT name, x, y, volumeLevel,
			musicLevel,
			pagingLevel,
			soundMaskingLevel,
			fadeTime,
			fadeLevel,
			averagingMode,
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
		`, speakerID).Scan(&name, &x, &y,
			&volumeLevel,
			&musicLevel,
			&pagingLevel,
			&soundMaskingLevel,
			&fadeTime,
			&fadeLevel,
			&averagingMode,
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
	speaker := &ControllerStatus{Name: name, X: x, Y: y, VolumeLevel: [4]int8{volumeLevel, musicLevel, pagingLevel, soundMaskingLevel}, PagingLevel: [2]int8{fadeTime, fadeLevel}, AveragingMode: averagingMode, ID: speakerID, Current: [21]int {band0, band1, band2, band3, band4, band5, band6, band7, band8, band9, band10, band11, band12, band13, band14, band15, band16, band17, band18, band19, band20}}


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
	FROM EqualizerPresets
	WHERE speakerID=?;
	`, speakerID)

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

		/*
		log.Println(name, whichPreset, 
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
		*/
			//constants := []int {band0, band1, band2, band3, band4, band5, band6, band7, band8, band9, band10, band11, band12, band13, band14, band15, band16, band17, band18, band19, band20}
			speaker.PresetNames = append(speaker.PresetNames, presetName)
			speaker.Equalizer = append(speaker.Equalizer, []int {band0, band1, band2, band3, band4, band5, band6, band7, band8, band9, band10, band11, band12, band13, band14, band15, band16, band17, band18, band19, band20})
			log.Println(speaker.PresetNames, speaker.Equalizer)		// make sure that the two arrays are the same size

		//log.Println("current", current)

		// make current the one that is stored in the speaker table
		/*
		if(current == 0) {
			constants := []int {band0, band1, band2, band3, band4, band5, band6, band7, band8, band9, band10, band11, band12, band13, band14, band15, band16, band17, band18, band19, band20}
			speaker.Equalizer  = append(speaker.Equalizer, constants)
		} else {
			speaker.Current = [21]int {band0, band1, band2, band3, band4, band5, band6, band7, band8, band9, band10, band11, band12, band13, band14, band15, band16, band17, band18, band19, band20}
			log.Println("hi from current:", speaker.Current)
		}
		*/
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// it is pulling from here but we need to figure out a way to distinguish which one should be loaded

	//speaker.Equalizer = constant in Current
	// might not need to append if we have the current in 
	//log.Println(speaker.VolumeLevel
	return speaker
}


func SavePreset(speakerId int8, name string, constants []string) {

	//var stringOfStatement string = "INSERT into EqualizerPresets VALUES (" + strconv.Itoa(int(speakerId)) + " " + name + "-1" + constants[0] + constants[1] + constants[2] + constants[3] + constants[4] + constants[5] + constants[6] + constants[7] + constants[8] + constants[9] + constants[10] + constants[11] + constants[12] + constants[13] + constants[14] + constants[15] + constants[16] + constants[17] + constants[18] + constants[19] + constants[20] + ")"
	log.Println("SavePreset: I am firing")
	var stringOfStatement string = "INSERT into EqualizerPresets VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, errr := DB.Exec(stringOfStatement, speakerId, name, "-1", constants[0], constants[1], constants[2], constants[3], constants[4], constants[5], constants[6], constants[7], constants[8], constants[9], constants[10], constants[11], constants[12], constants[13], constants[14], constants[15], constants[16], constants[17], constants[18], constants[19], constants[20])		// may not even need the speaker object to be passed in, idk why this shit is uneccessarily abstracted
	if errr != nil {				// also remember that the level that we are passing in is not offset, it is the true level, as it should be
		log.Fatal(errr)			
	}
	//log.Println(stringOfStatement)

	//statement, err := DB.Prepare(stringOfStatement)
}

// SaveSpeaker saves the speaker provided to the database
// reduce the generality if performance wanes and we need it
func SaveVolume(speaker *ControllerStatus) {
	log.Println("SaveVolume: I am firing")
	_, err := DB.Exec(`
		UPDATE speaker
		SET
			volumeLevel = ?,
			musicLevel = ?,
			pagingLevel = ?,
			soundMaskingLevel = ?
		WHERE speakerID = ?
	`, speaker.VolumeLevel[0], speaker.VolumeLevel[1], speaker.VolumeLevel[2], speaker.VolumeLevel[3], speaker.ID)	// the volumeLevel needs to change
	if err != nil {
		log.Fatal(err)
	}
	log.Println(speaker)
}

func SaveFade(speaker *ControllerStatus) {
	log.Println("SaveFade: I am firing")
	_, err := DB.Exec(`
		UPDATE speaker
		SET
			fadeTime = ?,
			fadeLevel = ?
		WHERE speakerID = ?
	`, speaker.PagingLevel[0], speaker.PagingLevel[1], speaker.ID)	// the volumeLevel needs to change
	if err != nil {
		log.Fatal(err)
	}
}

func SaveAveraging(speaker *ControllerStatus) {
	log.Println("SaveAveraging: I am firing")
	_, err := DB.Exec(`
		UPDATE speaker
		SET
			averagingMode = ?
		WHERE speakerID = ?
	`, speaker.AveragingMode, speaker.ID)
	if err != nil {
		log.Fatal(err)
	}
}

// make a command to save the band

func SaveBand(speaker *ControllerStatus, band int, level int) {

	var stringOfStatement string = "UPDATE speaker SET band" + strconv.Itoa(band) + " = ? WHERE speakerID = ?"
	log.Println("SaveBand: I am firing")

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
			musicLevel,
			pagingLevel,
			soundMaskingLevel,
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
		VALUES (?, ?, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
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
	log.Println("Getting speakers....")
	log.Println(speakerLocs)
	return speakerLocs.SpeakerLocations
}
