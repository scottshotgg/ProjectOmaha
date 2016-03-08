package database

type Account struct {
	UID      int
	UserName string
	Hash     string
}

type ControllerStatus struct {
	Name				string		`json:"name"`
	LEDOn         		bool 		`json:"ledOn"`
	ID            		int8 		`json:"id"`
	X            		int  		`json:"x"`
	Y					int  		`json:"y"`
	Effectiveness		int8 		`json:"effectiveness"`
	Pleasantness		int8 		`json:"Pleasantness"`
	VolumeLevel[4]  	int8 		`json:["volumeLevel"]`
	PagingLevel[2]		int8		`json:["pagingLevel"]`
	Target[][]			int			`json:["target"]`
	CurrentTarget[21] 	int			`json:"currentTarget"`
	TargetNames[]		string		`json:["targetNames"]`
	Equalizer[][]		int			`json:["equalizer"]`
	CurrentPreset[21]	int			`json:"currentPreset"`
	PresetNames[]		string		`json:["presetNames"]`
	WhichPreset			int			`json:"WhichPreset"`
}

type Zone struct {
	VolumeLevel int8                `json:"-"`		// what is this ???
	ID          int8                `json:"zoneID"`
	Name        string              `json:"name"`
	Speakers    []*ControllerStatus `json:"speakers"`
}
