package database

// YYYY-MM-DD HH:MM:SS.SSS is the format to use then for date and date-time, respecting the the ISO8601 string format.
// IF NOT EXISTS used for cheecking whether the table we are creting is not present in the DB.
// DEFAULT is used for checking the fixing the default values.
// UNIQUE is used for saying that the value must be unique among all rows.
var user_table = `CREATE TABLE IF NOT EXIST UserProfile (
	fixedUsername TEXT NOT NULL PRIMARY KEY, 
	username TEXT NOT NULL,
	uuid TEXT NOT NULL,
	photoProfile BLOB NOT NULL,
	biography TEXT NOT NULL DEFAULT " ",
	dateOfCreation TEXT NOT NULL DEFAULT "0000-01-01",									
	numberOfPhotos INTEGER NOT NULL DEFAULT 0,
	totNumberLikes INTEGER NOT NULL DEFAULT 0,
	totNumberComments INTEGER NOT NULL DEFAULT 0,
	numberFollowers INTEGER NOT NULL DEFAULT 0,
	numberFollowing INTEGER NOT NULL DEFAULT 0,
	name TEXT NOT NULL,
	surname TEXT NOT NULL,
	dateOfBirth TEXT NOT NULL DEFAULT "0000-01-01",
	email TEXT NOT NULL DEFAULT "NOT INSERTED",
	nationality TEXT NOT NULL DEFAULT "NOT INSERTED",
	gender TEXT NOT NULL DEFAULT "DO NOT SPECIFY"
	);`
