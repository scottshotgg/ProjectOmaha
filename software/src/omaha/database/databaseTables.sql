CREATE TABLE Speaker (
	SpeakerID int PRIMARY KEY,
	Name varchar(50),
	Volume FLOAT NOT NULL
)


CREATE TABLE  account IF NOT EXISTS (
	uid INTEGER PRIMARY KEY AUTOINCREMENT,
	username VARCHAR(20) NOT NULL UNIQUE,
	name VARCHAR(20),
	hash CHAR(60) NOT NULL
);

CREATE INDEX usernameIndex ON account(username);

CREATE TABLE Zone (
	zoneID INTEGER PRIMARY KEY AUTOINCREMENT,
	name varchar(50)	
)

CREATE TABLE Zones-to-Speakers (
	ZoneID int REFERENCES Zone(ZoneID),
	SpeakerID int REFERENCES Speaker(SpeakerID)
	PRIMARY KEY (ZoneID, SpeakerID)
)

CREATE TABLE Accounts-to-Speakers (
	userID varchar(50) REFERENCES Account(UserID),
	speakerID int REFERENCES Speaker(SpeakerID),
	PRIMARY KEY (userID, speakerID)	
)

CREATE TABLE Accounts-to-Zones (
	userID INTEGER REFERENCES Account(UserID),
	zoneID INTEGER REFERENCES Zone(ZoneID)
	PRIMARY KEY (userID, zoneID)
)