package database

type Account struct {
	UID      int
	UserName string
	Hash     string
}

type ControllerStatus struct {
	LEDOn         	bool 		`json:"ledOn"`
	VolumeLevel[3]  int8 		`json:"volumeLevel"`
	ID            	int8 		`json:"id"`
	X            	int  		`json:"x"`
	Y				int  		`json:"y"`
	AveragingMode	int8 		`json:"averagingMode"`
	Equalizer[21]	int			`json:["equalizer"]`
	PagingLevel		int8		`json:"pagingLevel"`
	/*
	Band0			int			`json:"band0"`	
	Band1			int			`json:"band1"`	
	Band2			int			`json:"band2"`	
	Band3			int			`json:"band3"`	
	Band4			int			`json:"band4"`	
	Band5			int			`json:"band5"`	
	Band6			int			`json:"band6"`	
	Band7			int			`json:"band7"`	
	Band8			int			`json:"band8"`	
	Band9			int			`json:"band9"`	
	Band10			int			`json:"band10"`	
	Band11			int			`json:"band11"`	
	Band12			int			`json:"band12"`	
	Band13			int			`json:"band13"`	
	Band14			int			`json:"band14"`
	Band15			int			`json:"band15"`	
	Band16			int			`json:"band16"`	
	Band17			int			`json:"band17"`	
	Band18			int			`json:"band18"`	
	Band19			int			`json:"band19"`	
	Band20			int			`json:"band20"`	
	*/
}

type Zone struct {
	VolumeLevel int8                `json:"-"`
	ID          int8                `json:"zoneID"`
	Name        string              `json:"name"`
	Speakers    []*ControllerStatus `json:"speakers"`
}
