package database

type Account struct {
	UID      int
	UserName string
	Hash     string
}

type ControllerStatus struct {
	Name										string		`json:"name"`
	LEDOn         					bool 			`json:"ledOn"`
	ID            					int8 			`json:"id"`
	X            						int  			`json:"x"`
	Y												int  			`json:"y"`
	Effectiveness						int8 			`json:"effectiveness"`
	Pleasantness						int8 			`json:"Pleasantness"`
	VolumeLevel[4]  				int8 			`json:["volumeLevel"]`
	LowerThreshold[3] 			int8 			`json:["lowerThreshold"]`
	UpperThreshold[3] 			int8 			`json:["upperThreshold"]`
	PagingLevel[2]					int8			`json:["pagingLevel"]`

	Target[][]							float64		`json:["target"]`
	TargetNames[]						string		`json:["targetNames"]`
	CurrentTarget[21] 			float64		`json:"currentTarget"`

	Equalizer[][]						float64		`json:["equalizer"]`
	PresetNames[]						string		`json:["presetNames"]`
	CurrentPreset[21]				float64		`json:"currentPreset"`

	MusicEqualizer[][]			float64		`json:["musicEqualizer"]`
	MusicPresetNames[]			string		`json:["pagingPresetNames"]`
	CurrentMusicPreset[21]	float64		`json:"currentMusicPreset"`

	PagingEqualizer[][]			float64		`json:["pagingEqualizer"]`
	PagingPresetNames[]			string		`json:["pagingPresetNames"]`
	CurrentPagingPreset[21]	float64		`json:"currentPagingPreset"`

	WhichPreset							int				`json:"whichPreset"`
	Status 									int				`json:"status"`
	EqualizerMode						int8			`json:"equalizerMode"`
	SchedulingMode					int8			`json:"schedulingMode"`

	Schedule[][]						int				`json:["schedule"]`
}

type Zone struct {
	Name										string		`json:"name"`
	Speakers    						[]*ControllerStatus 	`json:"speakers"`
	LEDOn         					bool 			`json:"ledOn"`
	ZoneID         					int8 			`json:"zoneID"`
	Effectiveness						int8 			`json:"effectiveness"`
	Pleasantness						int8 			`json:"pleasantness"`
	VolumeLevel[4]  				int8 			`json:["volumeLevel"]`
	LowerThreshold[3] 			int8 			`json:["lowerThreshold"]`
	UpperThreshold[3] 			int8 			`json:["upperThreshold"]`
	PagingLevel[2]					int8			`json:["pagingLevel"]`

	Target[][]							float64		`json:["target"]`
	TargetNames[]						string		`json:["targetNames"]`
	CurrentTarget[21] 			float64		`json:"currentTarget"`

	Equalizer[][]						float64		`json:["equalizer"]`
	PresetNames[]						string		`json:["presetNames"]`
	CurrentPreset[21]				float64		`json:"currentPreset"`

	MusicEqualizer[][]			float64		`json:["musicEqualizer"]`
	MusicPresetNames[]			string		`json:["pagingPresetNames"]`
	CurrentMusicPreset[21]	float64		`json:"currentMusicPreset"`

	PagingEqualizer[][]			float64		`json:["pagingEqualizer"]`
	PagingPresetNames[]			string		`json:["pagingPresetNames"]`
	CurrentPagingPreset[21]	float64		`json:"currentPagingPreset"`

	WhichPreset							int				`json:"whichPreset"`
	Status 									int				`json:"status"`
	EqualizerMode						int8			`json:"equalizerMode"`
	SchedulingMode					int8			`json:"schedulingMode"`

	Schedule[][]						int				`json:["schedule"]`
}
