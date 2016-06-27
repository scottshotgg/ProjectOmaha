CREATE TABLE Speaker (
	speakerID INTEGER PRIMARY KEY,
	name varchar(50),
	y INTEGER,
	x INTEGER
)

CREATE TABLE  account IF NOT EXISTS (
	uid INTEGER PRIMARY KEY AUTOINCREMENT,
	username VARCHAR(20) NOT NULL UNIQUE,
	name VARCHAR(20),
	hash CHAR(60) NOT NULL
);

CREATE INDEX usernameIndex ON account(username);

create table accountSession (
	sessionKey text primary key,
	userID int not null
)

CREATE TABLE zone (
	zoneID INTEGER PRIMARY KEY AUTOINCREMENT,
	name varchar(50)	
)

CREATE TABLE zoneToSpeaker (
	zoneID int REFERENCES Zone(ZoneID),
	speakerID int REFERENCES Speaker(SpeakerID)
	PRIMARY KEY (ZoneID, SpeakerID)
)

CREATE TABLE pagingZone (
	zoneID INTEGER PRIMARY KEY AUTOINCREMENT,
	name varchar(50)	
)

CREATE TABLE pagingZoneToSpeaker (
	zoneID int REFERENCES PagingZone(ZoneID),
	speakerID int REFERENCES Speaker(SpeakerID)
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