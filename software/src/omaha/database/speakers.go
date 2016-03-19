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
			effectiveness INTEGER,
			pleasantness INTEGER,
			status INTEGER DEFAULT 0,

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
			tBand20
		)
	`) // needs to be an equalizer thing in here
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created speaker table")
	}
}

func createEqualizerPresetsTable() {
	_, err := DB.Exec(`
		CREATE TABLE EqualizerPresets (
			speakerID INTEGER PRIMARY KEY,
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
			band20 REAL
		)
	`) // needs to be an equalizer thing in here
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created speaker table")
	}
}

func createTargetsTable() {
	_, err := DB.Exec(`
		CREATE TABLE Targets (
			speakerID INTEGER PRIMARY KEY,
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
			band20 REAL
		)
	`) // needs to be an equalizer thing in here
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created speaker table")
	}
}

/*func GetAllStatus() {

}*/

// GetAllSpeakers gets all speakers from the database and returns them as a slice of ControllerStatus objects
func GetAllSpeakers() []*ControllerStatus {		
// not sure if we need to get all the bands and stuff, doesnt look like that are used
	rows, err := DB.Query(`
		SELECT speakerID, name, x, y,  
			volumeLevel,
			musicLevel,
			pagingLevel,
			soundMaskingLevel,
			fadeTime,
			fadeLevel,
			effectiveness,
			pleasantness,
			status,

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

		FROM speaker
	`)


	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	speakers := []*ControllerStatus{}
	for rows.Next() {
		var speakerID int
		var name string
		var x int
		var y int
		var volumeLevel int8
		var musicLevel int8
		var pagingLevel int8
		var soundMaskingLevel int8
		var fadeTime int8
		var fadeLevel int8
		var effectiveness int8
		var pleasantness int8
		var status int
		// var currentMode int
		// var whichPreset int
		var pBand0 float64 // not sure if we need to get all the bands and stuff, doesnt look like that are used
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

		// in here get the presets that go along with the speakers and use those variables 

		if err := rows.Scan(
			&speakerID, 
			&name, 
			&x, 
			&y,
			&volumeLevel, 
			&musicLevel,
			&pagingLevel,
			&soundMaskingLevel,
			&fadeTime,
			&fadeLevel,
			&effectiveness,
			&pleasantness,
			&status,		

			// Current preset bands
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
	var effectiveness int8
	var pleasantness int8
	var x int
	var y int
	var name string
	var presetName string		// might need to make a catch for this
	var targetName string		// might need to make a catch for this
	var whichPreset int
	var status int

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

	var pBand0 float64 // not sure if we need to get all the bands and stuff, doesnt look like that are used
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

	err := DB.QueryRow(`
		SELECT name, x, y, volumeLevel,
			musicLevel,
			pagingLevel,
			soundMaskingLevel,
			fadeTime,
			fadeLevel,
			effectiveness,
			pleasantness,
			status,

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
		FROM speaker
		WHERE speakerID=?;
		`, speakerID).Scan(
			&name, 
			&x, 
			&y,
			&volumeLevel,
			&musicLevel,
			&pagingLevel,
			&soundMaskingLevel,
			&fadeTime,
			&fadeLevel,
			&effectiveness,
			&pleasantness,
			&status,

			// Current preset bands
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
	speaker := &ControllerStatus { 
		Name: name, 
		X: x, 
		Y: y, 

		VolumeLevel: [4] int8 {
			volumeLevel, 
			musicLevel, 
			pagingLevel, 
			soundMaskingLevel }, 

		PagingLevel: [2] int8 {
			fadeTime, 
			fadeLevel }, 

		Effectiveness: effectiveness,
		Pleasantness: pleasantness, 
		ID: speakerID, 
		Status: status,

		CurrentPreset: [21] float64 {
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
			speaker.Equalizer = append(speaker.Equalizer, []float64 {band0, band1, band2, band3, band4, band5, band6, band7, band8, band9, band10, band11, band12, band13, band14, band15, band16, band17, band18, band19, band20})
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

	rows, errr := DB.Query(`
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
	FROM Targets
	WHERE speakerID=?;
	`, speakerID)

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&targetName, &whichPreset, 
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
			speaker.TargetNames = append(speaker.TargetNames, targetName)
			speaker.Target = append(speaker.Target, []float64 {band0, band1, band2, band3, band4, band5, band6, band7, band8, band9, band10, band11, band12, band13, band14, band15, band16, band17, band18, band19, band20})
			log.Println(speaker.TargetNames, speaker.Target)		// make sure that the two arrays are the same size

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

func UpdateStatus(ID int8, status int8) {
	log.Println("Saving status")
	_, err := DB.Exec(`
		UPDATE speaker
		SET
			status = ?
		WHERE speakerID = ?
	`, status, ID)	
	if err != nil {
		log.Fatal(err)
	}
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

func UpdateTarget(speakerId int8, name string, constants []string) {

	//var stringOfStatement string = "INSERT into EqualizerPresets VALUES (" + strconv.Itoa(int(speakerId)) + " " + name + "-1" + constants[0] + constants[1] + constants[2] + constants[3] + constants[4] + constants[5] + constants[6] + constants[7] + constants[8] + constants[9] + constants[10] + constants[11] + constants[12] + constants[13] + constants[14] + constants[15] + constants[16] + constants[17] + constants[18] + constants[19] + constants[20] + ")"
	log.Println("SaveTarget: I am firing")
	var stringOfStatement string = "INSERT into Targets VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, errr := DB.Exec(stringOfStatement, speakerId, name, "-1", constants[0], constants[1], constants[2], constants[3], constants[4], constants[5], constants[6], constants[7], constants[8], constants[9], constants[10], constants[11], constants[12], constants[13], constants[14], constants[15], constants[16], constants[17], constants[18], constants[19], constants[20])		// may not even need the speaker object to be passed in, idk why this shit is uneccessarily abstracted
	if errr != nil {				// also remember that the level that we are passing in is not offset, it is the true level, as it should be
		log.Fatal(errr)			
	}
	//log.Println(stringOfStatement)

	//statement, err := DB.Prepare(stringOfStatement)
}


func SaveTarget(speakerId int8, name string, constants []string) {

	//var stringOfStatement string = "INSERT into EqualizerPresets VALUES (" + strconv.Itoa(int(speakerId)) + " " + name + "-1" + constants[0] + constants[1] + constants[2] + constants[3] + constants[4] + constants[5] + constants[6] + constants[7] + constants[8] + constants[9] + constants[10] + constants[11] + constants[12] + constants[13] + constants[14] + constants[15] + constants[16] + constants[17] + constants[18] + constants[19] + constants[20] + ")"
	log.Println("SaveTarget: I am firing")
	var stringOfStatement string = "INSERT into Targets VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
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
			effectiveness = ?,
			pleasantness = ?
		WHERE speakerID = ?
	`, speaker.Effectiveness, speaker.Pleasantness, speaker.ID)
	if err != nil {
		log.Fatal(err)
	}
}

// make a command to save the band

func SaveBand(speaker *ControllerStatus, band int, level float64, target bool) { // this should return an error here if something is wrong
	var stringOfStatement string

	if (target == false) {
		stringOfStatement = "UPDATE speaker SET pBand" + strconv.Itoa(band) + " = ? WHERE speakerID = ?"
	} else {
		stringOfStatement = "UPDATE speaker SET tBand" + strconv.Itoa(band) + " = ? WHERE speakerID = ?"
	}
	//stmt, err := DB.Prepare()			// this might need to be prepared afterwards
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

func CreatePagingZone(speakers []int8, name string) error {
	log.Println(name, speakers)

	//stmt, err := DB.Prepare(`INSERT INTO zone (name) VALUES (?)`, name)

	_, ziError := DB.Exec(`INSERT INTO pagingZone (name) VALUES (?)`, name)

	_, zsError := DB.Exec(`SELECT zoneID FROM pagingZone where name=?`, name)
	log.Println(ziError, zsError)

	var zoneID int8

	err := DB.QueryRow(`SELECT zoneID FROM pagingZone where name=?`, name).Scan(&zoneID)
	log.Println(zoneID, err)


	// this may need to be transacted instead of single
	for i := 0; i < len(speakers); i++ {
		_, ztsError := DB.Exec(`INSERT INTO pagingZoneToSpeaker (zoneID, speakerID) VALUES (?, ?)`, zoneID, speakers[i])
		log.Println("I am le running: ", speakers[i], ztsError)
	}
	return nil
}

func CreateZone(speakers []int8, name string) error {
	log.Println(name, speakers)

	//stmt, err := DB.Prepare(`INSERT INTO zone (name) VALUES (?)`, name)

	_, ziError := DB.Exec(`INSERT INTO zone (name) VALUES (?)`, name)

	_, zsError := DB.Exec(`SELECT zoneID FROM zone where name=?`, name)
	log.Println(ziError, zsError)

	var zoneID int8

	err := DB.QueryRow(`SELECT zoneID FROM zone where name=?`, name).Scan(&zoneID)
	log.Println(zoneID, err)


	// this may need to be transacted instead of single
	for i := 0; i < len(speakers); i++ {
		_, ztsError := DB.Exec(`INSERT INTO zoneToSpeaker (zoneID, speakerID) VALUES (?, ?)`, zoneID, speakers[i])
		log.Println("I am le running: ", speakers[i], ztsError)
	}
	return nil
}

func getInsertSpeakerStatement() (*sql.Stmt, error) {
	stmt, err := DB.Prepare(`INSERT INTO speaker 
		(name, x, y, volumeLevel,
			musicLevel,
			pagingLevel,
			soundMaskingLevel,
			fadeTime,
			fadeLevel,
			effectiveness,
			pleasantness,

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
		VALUES ("", ?, ?, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
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
	defer func(speakerStmt, zoneStmt *sql.Stmt, pagingZoneStmt *sql.Stmt) {
		insertSpeakerStmt = speakerStmt
		addSpeakerToZonesStmt = zoneStmt
		addSpeakerToPagingZonesStmt = pagingZoneStmt
	}(insertSpeakerStmt, addSpeakerToZonesStmt, addSpeakerToPagingZonesStmt)
	// change the statements to use this transaction
	insertSpeakerStmt = tx.Stmt(insertSpeakerStmt)
	addSpeakerToZonesStmt = tx.Stmt(addSpeakerToZonesStmt)
	addSpeakerToPagingZonesStmt = tx.Stmt(addSpeakerToPagingZonesStmt)
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
