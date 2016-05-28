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
	`) // needs to be an equalizer thing in here
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created speaker table")
	}
}

/*
paBand0 REAL, 
			paBand1 REAL, 
			paBand2 REAL, 
			paBand3 REAL, 
			paBand4 REAL, 
			paBand5 REAL, 
			paBand6 REAL, 
			paBand7 REAL, 
			paBand8 REAL, 
			paBand9 REAL, 
			paBand10 REAL, 
			paBand11 REAL, 
			paBand12 REAL, 
			paBand13 REAL, 
			paBand14 REAL, 
			paBand15 REAL, 
			paBand16 REAL, 
			paBand17 REAL, 
			paBand18 REAL, 
			paBand19 REAL, 
			paBand20 REAL,
*/
func createEqualizerPresetsTable() {
	_, err := DB.Exec(`
		CREATE TABLE EqualizerPresets (
			speakerID INTEGER,
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
			PRIMARY KEY (speakerID, name)
		)
	`) // needs to be an equalizer thing in here
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created equalizer table")
	}
}

func createMusicEqualizerPresetsTable() {
	_, err := DB.Exec(`
		CREATE TABLE MusicEqualizerPresets (
			speakerID INTEGER,
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
			PRIMARY KEY (speakerID, name)
		)
	`) // needs to be an equalizer thing in here
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created music equalizer table")
	}
}

func createPagingEqualizerPresetsTable() {
	_, err := DB.Exec(`
		CREATE TABLE PagingEqualizerPresets (
			speakerID INTEGER,
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
			PRIMARY KEY (speakerID, name)
		)
	`) // needs to be an equalizer thing in here
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created paging equalizer table")
	}
}

func createTargetsTable() {
	_, err := DB.Exec(`
		CREATE TABLE Targets (
			speakerID INTEGER,
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
			PRIMARY KEY (speakerID, name)
		)
	`) // needs to be an equalizer thing in here
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created targets table")
	}
}

func createSchedulingTable() {
	_, err := DB.Exec(`
		CREATE TABLE Scheduling (
			speakerID INTEGER,
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
	`) // needs to be an equalizer thing in here
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created scheduling table")
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
		var status int
		var equalizerMode int8
		var schedulingMode int8
		// var currentMode int
		// var whichPreset int
		var eBand0 float64			// why do we need to pull all of this stuff, this can be optimized!!
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

		var pBand0 float64			// why do we need to pull all of this stuff, this can be optimized!!
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

			// we may not even need to scan all this.....

			// Current preset bands
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
		speaker.Status = status
		speaker.EqualizerMode = equalizerMode
		speaker.SchedulingMode = schedulingMode

		speaker.VolumeLevel[0] = int8(volumeLevel)
		speaker.VolumeLevel[1] = int8(musicLevel)
		speaker.VolumeLevel[2] = int8(pagingLevel)
		speaker.VolumeLevel[3] = int8(soundMaskingLevel)

		speaker.LowerThreshold[0] = int8(lowerMusicThreshold)
		speaker.LowerThreshold[1] = int8(lowerPagingThreshold)
		speaker.LowerThreshold[2] = int8(lowerMusicThreshold)

		speaker.UpperThreshold[0] = int8(upperMusicThreshold)
		speaker.UpperThreshold[1] = int8(upperPagingThreshold)
		speaker.UpperThreshold[2] = int8(upperMusicThreshold)

		speaker.PagingLevel[0] = int8(fadeTime)
		speaker.PagingLevel[1] = int8(fadeLevel)
		// might need an averaging setter here later
		speakers = append(speakers, speaker)
	}
	return speakers
}

// GetSpeaker gets a speaker from the database
func GetSpeaker(speakerID int8) *ControllerStatus {
	var speakerid int8

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
	var x int
	var y int
	var name string
	var presetName string		// might need to make a catch for this
	var targetName string		// might need to make a catch for this
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


//	var interval int
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
		SELECT name, x, y, volumeLevel,
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

			// Current preset bands
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
		ID: speakerID, 
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
			
			speaker.PresetNames = append(speaker.PresetNames, presetName)
			speaker.Equalizer = append(speaker.Equalizer, []float64 {band0, band1, band2, band3, band4, band5, band6, band7, band8, band9, band10, band11, band12, band13, band14, band15, band16, band17, band18, band19, band20})
			log.Println(speaker.PresetNames, speaker.Equalizer)		// make sure that the two arrays are the same size
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
	FROM MusicEqualizerPresets
	WHERE speakerID=?;
	`, speakerID)

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
			
			speaker.MusicPresetNames = append(speaker.MusicPresetNames, presetName)
			speaker.MusicEqualizer = append(speaker.MusicEqualizer, []float64 {band0, band1, band2, band3, band4, band5, band6, band7, band8, band9, band10, band11, band12, band13, band14, band15, band16, band17, band18, band19, band20})
			log.Println(speaker.MusicPresetNames, speaker.MusicEqualizer)		// make sure that the two arrays are the same size
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
	FROM PagingEqualizerPresets
	WHERE speakerID=?;
	`, speakerID)

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
			
			speaker.PagingPresetNames = append(speaker.PagingPresetNames, presetName)
			speaker.PagingEqualizer = append(speaker.PagingEqualizer, []float64 {band0, band1, band2, band3, band4, band5, band6, band7, band8, band9, band10, band11, band12, band13, band14, band15, band16, band17, band18, band19, band20})
			log.Println(speaker.PagingPresetNames, speaker.PagingEqualizer)		// make sure that the two arrays are the same size
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
	FROM Targets
	WHERE speakerID=?;
	`, speakerID)

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


	rows, errr = DB.Query(`
		SELECT * 
		FROM Scheduling
		WHERE speakerID = ?
		`, speakerID)

	for rows.Next() {
		err = rows.Scan(
			//&interval,
			&speakerid,
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

		//log.Println(day)
		// interval needs to go in here as well
		speaker.Schedule = append(speaker.Schedule, []int {day, interval, start, end, hour0, hour1, hour2, hour3, hour4, hour5, hour6, hour7, hour8, hour9, hour10, hour11, hour12, hour13, hour14, hour15, hour16, hour17, hour18, hour19, hour20, hour21, hour22, hour23})

	}
	//log.Println(speaker.Music)

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

	log.Println(ID, status)
}


func UpdateSchedule(speakerid int8, day int, interval int, start int, end int, times [24]int) {
	log.Println("UpdateSchedule")
	log.Println(speakerid, start, end, times, day)

	for x := 0; x < len(times); x++ {

		var stringOfStatement string = "UPDATE Scheduling SET hour" + strconv.Itoa(x) + "=? WHERE speakerID=? AND day=?"
		_, err := DB.Exec(stringOfStatement, times[x], speakerid, day)	
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("updated hour", x, speakerid, times[x], day)
		}

		stringOfStatement = "UPDATE Scheduling SET interval=?, start=?, end=? WHERE speakerID=? AND day=?"
		_, err = DB.Exec(stringOfStatement, interval, start, end, speakerid, day)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("updated hour", x, speakerid, times[x], day)
		}
	}
}


func UpdateScheduleZone(zoneid int8, day int, interval int, start int, end int, times [24]int) {
	log.Println("UpdateScheduleZone")
	log.Println(zoneid, start, end, times, day)

	for x := 0; x < len(times); x++ {

		var stringOfStatement string = "UPDATE SchedulingZone SET hour" + strconv.Itoa(x) + "=? WHERE zoneID=? AND day=?"
		_, err := DB.Exec(stringOfStatement, times[x], zoneid, day)	
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("updated hour", x, zoneid, times[x], day)
		}

		stringOfStatement = "UPDATE SchedulingZone SET interval=?, start=?, end=? WHERE zoneID=? AND day=?"
		_, err = DB.Exec(stringOfStatement, interval, start, end, zoneid, day)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("updated hour", x, zoneid, times[x], day)
		}
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


func SaveMusicPreset(speakerId int8, name string, constants []string) {

	//var stringOfStatement string = "INSERT into EqualizerPresets VALUES (" + strconv.Itoa(int(speakerId)) + " " + name + "-1" + constants[0] + constants[1] + constants[2] + constants[3] + constants[4] + constants[5] + constants[6] + constants[7] + constants[8] + constants[9] + constants[10] + constants[11] + constants[12] + constants[13] + constants[14] + constants[15] + constants[16] + constants[17] + constants[18] + constants[19] + constants[20] + ")"
	log.Println("SaveMusicPreset: I am firing")
	var stringOfStatement string = "INSERT into MusicEqualizerPresets VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, errr := DB.Exec(stringOfStatement, speakerId, name, "-1", constants[0], constants[1], constants[2], constants[3], constants[4], constants[5], constants[6], constants[7], constants[8], constants[9], constants[10], constants[11], constants[12], constants[13], constants[14], constants[15], constants[16], constants[17], constants[18], constants[19], constants[20])		// may not even need the speaker object to be passed in, idk why this shit is uneccessarily abstracted
	if errr != nil {				// also remember that the level that we are passing in is not offset, it is the true level, as it should be
		log.Fatal(errr)			
	}
	//log.Println(stringOfStatement)

	//statement, err := DB.Prepare(stringOfStatement)
}


func SavePagingPreset(speakerId int8, name string, constants []string) {

	//var stringOfStatement string = "INSERT into EqualizerPresets VALUES (" + strconv.Itoa(int(speakerId)) + " " + name + "-1" + constants[0] + constants[1] + constants[2] + constants[3] + constants[4] + constants[5] + constants[6] + constants[7] + constants[8] + constants[9] + constants[10] + constants[11] + constants[12] + constants[13] + constants[14] + constants[15] + constants[16] + constants[17] + constants[18] + constants[19] + constants[20] + ")"
	log.Println("SavePagingPreset: I am firing")
	var stringOfStatement string = "INSERT into PagingEqualizerPresets VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, errr := DB.Exec(stringOfStatement, speakerId, name, "-1", constants[0], constants[1], constants[2], constants[3], constants[4], constants[5], constants[6], constants[7], constants[8], constants[9], constants[10], constants[11], constants[12], constants[13], constants[14], constants[15], constants[16], constants[17], constants[18], constants[19], constants[20])		// may not even need the speaker object to be passed in, idk why this shit is uneccessarily abstracted
	if errr != nil {				// also remember that the level that we are passing in is not offset, it is the true level, as it should be
		log.Fatal(errr)			
	}
	//log.Println(stringOfStatement)

	//statement, err := DB.Prepare(stringOfStatement)
}


// wtf is this for, cant remember why I created this
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


func SavePresetZone(zoneid int8, name string, constants []string) {

	//var stringOfStatement string = "INSERT into EqualizerPresets VALUES (" + strconv.Itoa(int(speakerId)) + " " + name + "-1" + constants[0] + constants[1] + constants[2] + constants[3] + constants[4] + constants[5] + constants[6] + constants[7] + constants[8] + constants[9] + constants[10] + constants[11] + constants[12] + constants[13] + constants[14] + constants[15] + constants[16] + constants[17] + constants[18] + constants[19] + constants[20] + ")"
	log.Println("SavePresetZone: I am firing")
	var stringOfStatement string = "INSERT into EqualizerPresetsZone VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, errr := DB.Exec(stringOfStatement, zoneid, name, "-1", constants[0], constants[1], constants[2], constants[3], constants[4], constants[5], constants[6], constants[7], constants[8], constants[9], constants[10], constants[11], constants[12], constants[13], constants[14], constants[15], constants[16], constants[17], constants[18], constants[19], constants[20])		// may not even need the speaker object to be passed in, idk why this shit is uneccessarily abstracted
	if errr != nil {				// also remember that the level that we are passing in is not offset, it is the true level, as it should be
		log.Fatal(errr)			
	}
	//log.Println(stringOfStatement)

	//statement, err := DB.Prepare(stringOfStatement)
}

func SaveMusicPresetZone(zoneid int8, name string, constants []string) {

	//var stringOfStatement string = "INSERT into EqualizerPresets VALUES (" + strconv.Itoa(int(speakerId)) + " " + name + "-1" + constants[0] + constants[1] + constants[2] + constants[3] + constants[4] + constants[5] + constants[6] + constants[7] + constants[8] + constants[9] + constants[10] + constants[11] + constants[12] + constants[13] + constants[14] + constants[15] + constants[16] + constants[17] + constants[18] + constants[19] + constants[20] + ")"
	log.Println("SaveMusicPresetZone: I am firing")
	var stringOfStatement string = "INSERT into MusicEqualizerPresetsZone VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, errr := DB.Exec(stringOfStatement, zoneid, name, "-1", constants[0], constants[1], constants[2], constants[3], constants[4], constants[5], constants[6], constants[7], constants[8], constants[9], constants[10], constants[11], constants[12], constants[13], constants[14], constants[15], constants[16], constants[17], constants[18], constants[19], constants[20])		// may not even need the speaker object to be passed in, idk why this shit is uneccessarily abstracted
	if errr != nil {				// also remember that the level that we are passing in is not offset, it is the true level, as it should be
		log.Fatal(errr)			
	}
	//log.Println(stringOfStatement)

	//statement, err := DB.Prepare(stringOfStatement)
}

func SavePagingPresetZone(zoneid int8, name string, constants []string) {

	//var stringOfStatement string = "INSERT into EqualizerPresets VALUES (" + strconv.Itoa(int(speakerId)) + " " + name + "-1" + constants[0] + constants[1] + constants[2] + constants[3] + constants[4] + constants[5] + constants[6] + constants[7] + constants[8] + constants[9] + constants[10] + constants[11] + constants[12] + constants[13] + constants[14] + constants[15] + constants[16] + constants[17] + constants[18] + constants[19] + constants[20] + ")"
	log.Println("SavePagingPresetZone: I am firing")
	var stringOfStatement string = "INSERT into PagingEqualizerPresetsZone VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, errr := DB.Exec(stringOfStatement, zoneid, name, "-1", constants[0], constants[1], constants[2], constants[3], constants[4], constants[5], constants[6], constants[7], constants[8], constants[9], constants[10], constants[11], constants[12], constants[13], constants[14], constants[15], constants[16], constants[17], constants[18], constants[19], constants[20])		// may not even need the speaker object to be passed in, idk why this shit is uneccessarily abstracted
	if errr != nil {				// also remember that the level that we are passing in is not offset, it is the true level, as it should be
		log.Fatal(errr)			
	}
	//log.Println(stringOfStatement)

	//statement, err := DB.Prepare(stringOfStatement)
}


// wtf is this for, cant remember why I created this
func UpdateTargetZone(zoneid int8, name string, constants []string) {

	//var stringOfStatement string = "INSERT into EqualizerPresets VALUES (" + strconv.Itoa(int(speakerId)) + " " + name + "-1" + constants[0] + constants[1] + constants[2] + constants[3] + constants[4] + constants[5] + constants[6] + constants[7] + constants[8] + constants[9] + constants[10] + constants[11] + constants[12] + constants[13] + constants[14] + constants[15] + constants[16] + constants[17] + constants[18] + constants[19] + constants[20] + ")"
	log.Println("SaveTargetZone: I am firing")
	var stringOfStatement string = "INSERT into TargetsZone VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, errr := DB.Exec(stringOfStatement, zoneid, name, "-1", constants[0], constants[1], constants[2], constants[3], constants[4], constants[5], constants[6], constants[7], constants[8], constants[9], constants[10], constants[11], constants[12], constants[13], constants[14], constants[15], constants[16], constants[17], constants[18], constants[19], constants[20])		// may not even need the speaker object to be passed in, idk why this shit is uneccessarily abstracted
	if errr != nil {				// also remember that the level that we are passing in is not offset, it is the true level, as it should be
		log.Fatal(errr)			
	}
	//log.Println(stringOfStatement)

	//statement, err := DB.Prepare(stringOfStatement)
}


func SaveTargetZone(zoneid int8, name string, constants []string) {

	//var stringOfStatement string = "INSERT into EqualizerPresets VALUES (" + strconv.Itoa(int(speakerId)) + " " + name + "-1" + constants[0] + constants[1] + constants[2] + constants[3] + constants[4] + constants[5] + constants[6] + constants[7] + constants[8] + constants[9] + constants[10] + constants[11] + constants[12] + constants[13] + constants[14] + constants[15] + constants[16] + constants[17] + constants[18] + constants[19] + constants[20] + ")"
	log.Println("SaveTargetZone: I am firing")
	var stringOfStatement string = "INSERT into TargetsZone VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, errr := DB.Exec(stringOfStatement, zoneid, name, "-1", constants[0], constants[1], constants[2], constants[3], constants[4], constants[5], constants[6], constants[7], constants[8], constants[9], constants[10], constants[11], constants[12], constants[13], constants[14], constants[15], constants[16], constants[17], constants[18], constants[19], constants[20])		// may not even need the speaker object to be passed in, idk why this shit is uneccessarily abstracted
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

func SaveVolumeZone(zone *Zone) {
	log.Println("SaveVolumeZone: I am firing")
	_, err := DB.Exec(`
		UPDATE zone
		SET
			volumeLevel = ?,
			musicLevel = ?,
			pagingLevel = ?,
			soundMaskingLevel = ?
		WHERE zoneID = ?
	`, zone.VolumeLevel[0], zone.VolumeLevel[1], zone.VolumeLevel[2], zone.VolumeLevel[3], zone.ZoneID)	// the volumeLevel needs to change
	if err != nil {
		log.Fatal(err)
	}
	log.Println(zone)
}


func UpdateThreshold(speaker int8, musicMin int8, musicMax int8, pagingMin int8, pagingMax int8, maskingMin int8, maskingMax int8) {
	log.Println("UpdateThreshold: I am firing", speaker)
	_, err := DB.Exec(`
		UPDATE speaker
		SET 
			lowerMusicThreshold = ?,
			lowerPagingThreshold = ?,
			lowerMaskingThreshold = ?,
			upperMusicThreshold = ?,
			upperPagingThreshold = ?,
			upperMaskingThreshold = ?
		WHERE speakerID = ?
			`, musicMin, musicMax, pagingMin, pagingMax, maskingMin, maskingMax, speaker)
	if err != nil {
		log.Fatal(err)
	}
}


func UpdateThresholdZone(zone int8, musicMin int8, musicMax int8, pagingMin int8, pagingMax int8, maskingMin int8, maskingMax int8) {
	log.Println("UpdateThreshold: I am firing", zone)
	_, err := DB.Exec(`
		UPDATE zone
		SET 
			lowerMusicThreshold = ?,
			lowerPagingThreshold = ?,
			lowerMaskingThreshold = ?,
			upperMusicThreshold = ?,
			upperPagingThreshold = ?,
			upperMaskingThreshold = ?
		WHERE zoneID = ?
			`, musicMin, musicMax, pagingMin, pagingMax, maskingMin, maskingMax, zone)
	if err != nil {
		log.Fatal(err)
	}
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


func SaveFadeZone(zone *Zone) {
	log.Println("SaveFade: I am firing")
	_, err := DB.Exec(`
		UPDATE zone
		SET
			fadeTime = ?,
			fadeLevel = ?
		WHERE zoneID = ?
	`, zone.PagingLevel[0], zone.PagingLevel[1], zone.ZoneID)	// the volumeLevel needs to change
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


func SaveAveragingZone(zone *Zone) {
	log.Println("SaveAveragingZone: I am firing")
	_, err := DB.Exec(`
		UPDATE zone
		SET
			effectiveness = ?,
			pleasantness = ?
		WHERE zoneID = ?
	`, zone.Effectiveness, zone.Pleasantness, zone.ZoneID)
	if err != nil {
		log.Fatal(err)
	}
}


func SaveBand(speaker *ControllerStatus, band int, level float64, typeOfPreset int) { // this should return an error here if something is wrong
	var stringOfStatement string

	switch typeOfPreset {
		case 0:
			stringOfStatement = "UPDATE speaker SET eBand" + strconv.Itoa(band) + " = ? WHERE speakerID = ?"
		case 1:
			stringOfStatement = "UPDATE speaker SET mBand" + strconv.Itoa(band) + " = ? WHERE speakerID = ?"
		case 2:
			stringOfStatement = "UPDATE speaker SET pBand" + strconv.Itoa(band) + " = ? WHERE speakerID = ?"
		case 3:
			stringOfStatement = "UPDATE speaker SET tBand" + strconv.Itoa(band) + " = ? WHERE speakerID = ?"
		default:
			log.Println("SAVEBAND WAS PASSED AN INVALID PRESET OR TARGET VALUE")
	}
	//stmt, err := DB.Prepare()			// this might need to be prepared afterwards
	log.Println("SaveBand: I am firing")

	//log.Println(stringOfStatement)

	//statement, err := DB.Prepare(stringOfStatement)
	// need to check whether or not this is null or w/e before sending it off
	_, errr := DB.Exec(stringOfStatement, level, speaker.ID)		// may not even need the speaker object to be passed in, idk why this shit is uneccessarily abstracted
	if errr != nil {				// also remember that the level that we are passing in is not offset, it is the true level, as it should be
		log.Fatal(errr)			
	}
}

func SaveBandZone(zone *Zone, band int, level float64, typeOfPreset int) { // this should return an error here if something is wrong
	var stringOfStatement string

	switch typeOfPreset {
		case 0:
			stringOfStatement = "UPDATE zone SET eBand" + strconv.Itoa(band) + " = ? WHERE zoneID = ?"
		case 1:
			stringOfStatement = "UPDATE zone SET mBand" + strconv.Itoa(band) + " = ? WHERE zoneID = ?"
		case 2:
			stringOfStatement = "UPDATE zone SET pBand" + strconv.Itoa(band) + " = ? WHERE zoneID = ?"
		case 3:
			stringOfStatement = "UPDATE zone SET tBand" + strconv.Itoa(band) + " = ? WHERE zoneID = ?"
		default:
			log.Println("SAVEBANDZONE WAS PASSED AN INVALID PRESET OR TARGET VALUE")
	}

	log.Println("SaveBandZone: I am firing")

	//log.Println(stringOfStatement)

	//statement, err := DB.Prepare(stringOfStatement)
	// need to check whether or not this is null or w/e before sending it off
	_, errr := DB.Exec(stringOfStatement, level, zone.ZoneID)		// may not even need the speaker object to be passed in, idk why this shit is uneccessarily abstracted
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
		//addSpeakerToAllZones(int8(id), zoneStmt2)
		log.Printf("Created speaker %d at (%d, %d)\n", id, loc.X, loc.Y)
		return int8(id)
	}
	return 0
}

func ChangeEQMode(speaker int8, mode int8) error {
	var statement = "UPDATE speaker SET equalizerMode = ? WHERE speakerID = ?"
	_, err := DB.Exec(statement, mode, speaker)

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func ChangeEQModeZone(zone int8, mode int8) error {
	var statement = "UPDATE zone SET equalizerMode = ? WHERE zoneID = ?"
	_, err := DB.Exec(statement, mode, zone)

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func ChangeSchedulingMode(speaker int8, mode int) error {
	var statement = "UPDATE speaker SET schedulingMode = ? WHERE speakerID = ?"
	_, err := DB.Exec(statement, mode, speaker)

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func ChangeSchedulingModeZone(zone int8, mode int) error {
	var statement = "UPDATE zone SET schedulingMode = ? WHERE zoneID = ?"
	_, err := DB.Exec(statement, mode, zone)

	if err != nil {
		log.Fatal(err)
	}
	return nil
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
		(name, x, y, 
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
		VALUES ("", ?, ?, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
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
	addSpeakerToSchedulingStmt = tx.Stmt(addSpeakerToSchedulingStmt)

	speakerLocations := getSpeakerLocations()
	for _, speakerLoc := range speakerLocations {
		id := addSpeaker(speakerLoc)
		addSpeakerToAllZones(id)
		addSpeakerSchedule(id)
	}
	tx.Commit()
}

func populateZoneTable() {
	AddZone("all_masking")
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
