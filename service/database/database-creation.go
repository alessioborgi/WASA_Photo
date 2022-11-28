package database

// YYYY-MM-DD HH:MM:SS.SSS is the format to use then for date and date-time, respecting the the ISO8601 string format.
// IF NOT EXISTS used for cheecking whether the table we are creting is not present in the DB.
// DEFAULT is used for checking the fixing the default values.
// UNIQUE is used for saying that the value must be unique among all rows.
var user_table = `CREATE TABLE IF NOT EXIST UserProfile (
	fixedUsername TEXT NOT NULL PRIMARY KEY, 
	username TEXT NOT NULL UNIQUE,
	uuid TEXT NOT NULL UNIQUE,
	photoProfile BLOB NOT NULL,
	biography TEXT NOT NULL,
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

// FOREIGN KEY (attribute_table) REFERENCES external_table (external_attribute) ON DELETE CASCADE
// ON DELETE CASCADE used for removing the ban if the users are removed.
var ban_table = `CREATE TABLE IF NOT EXIST Ban (
	banner TEXT NOT NULL UNIQUE, 
	banned TEXT NOT NULL UNIQUE,
	PRIMARY KEY (banner, banned),
	uploadDate TEXT NOT NULL DEFAULT "0000-01-01",
	Motivation BLOB NOT NULL DEFAULT "Spam",
	FOREIGN KEY (banner) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE,
	FOREIGN KEY (banned) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE,
	);`

var follow_table = `CREATE TABLE IF NOT EXIST Follow (
	follower TEXT NOT NULL UNIQUE, 
	followed TEXT NOT NULL UNIQUE,
	uploadDate TEXT NOT NULL DEFAULT "0000-01-01",
	PRIMARY KEY (follower, followed),
	FOREIGN KEY (follower) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE,
	FOREIGN KEY (followed) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE,
	);`

// Notice that here I declared only (commentid,photoid) as PK because adding also Commenter fixedUsername has not sense.
var comment_table = `CREATE TABLE IF NOT EXIST Comment (
	commentid INTEGER AUTOINCREMENT, 
	photoid INTEGER NOT NULL,
	phrase TEXT NOT NULL,
	commenter TEXT NOT NULL,
	uploadDate TEXT NOT NULL DEFAULT "0000-01-01",
	PRIMARY KEY (commentid, photoid),
	FOREIGN KEY (photoid) REFERENCES Photo (photoid) ON DELETE CASCADE,
	FOREIGN KEY (commenter) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE,
	);`
