package database

type Account struct {
	UID      int
	UserName string
	Hash     string
}

type ControllerStatus struct {
	Name					string		`json:"name"`
	LEDOn         			bool 		`json:"ledOn"`
	ID            			int8 		`json:"id"`
	X            			int  		`json:"x"`
	Y						int  		`json:"y"`
	Effectiveness			int8 		`json:"effectiveness"`
	Pleasantness			int8 		`json:"Pleasantness"`
	VolumeLevel[4]  		int8 		`json:["volumeLevel"]`
	PagingLevel[2]			int8		`json:["pagingLevel"]`

	Target[][]				float64		`json:["target"]`
	TargetNames[]			string		`json:["targetNames"]`
	CurrentTarget[21] 		float64		`json:"currentTarget"`

	Equalizer[][]			float64		`json:["equalizer"]`
	PresetNames[]			string		`json:["presetNames"]`
	CurrentPreset[21]		float64		`json:"currentPreset"`

	MusicEqualizer[][]		float64		`json:["musicEqualizer"]`
	MusicPresetNames[]		string		`json:["pagingPresetNames"]`
	CurrentMusicPreset[21]	float64		`json:"currentMusicPreset"`

	PagingEqualizer[][]		float64		`json:["pagingEqualizer"]`
	PagingPresetNames[]		string		`json:["pagingPresetNames"]`
	CurrentPagingPreset[21]	float64		`json:"currentPagingPreset"`

	WhichPreset				int			`json:"whichPreset"`			// really need to delete this or w/e
	Status 					int			`json:"status"`
}

type Zone struct {
	VolumeLevel int8                `json:"-"`		// what is this ???
	ID          int8                `json:"zoneID"`
	Name        string              `json:"name"`
	Speakers    []*ControllerStatus `json:"speakers"`
}
