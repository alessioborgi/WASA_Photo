package database

// ----- USER-RELATED: -----

// YYYY-MM-DD HH:MM:SS.SSS is the format to use then for date and date-time, respecting the the ISO8601 string format.
// IF NOT EXISTS used for cheecking whether the table we are creting is not present in the DB.
// DEFAULT is used for checking the fixing the default values.
// UNIQUE is used for saying that the value must be unique among all rows.
// Note that photoProfile and Biography coul also be NULL.
var user_table = `CREATE TABLE IF NOT EXIST UserProfile (
	fixedUsername TEXT NOT NULL PRIMARY KEY, 
	username TEXT NOT NULL UNIQUE,
	uuid TEXT NOT NULL UNIQUE,
	photoProfile BLOB,
	biography TEXT,
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
	banner TEXT NOT NULL, 
	banned TEXT NOT NULL,
	PRIMARY KEY (banner, banned),
	uploadDate TEXT NOT NULL DEFAULT "0000-01-01T00:00:00Z",
	Motivation TEXT NOT NULL DEFAULT "Spam",
	FOREIGN KEY (banner) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE,
	FOREIGN KEY (banned) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE
	);`

var follow_table = `CREATE TABLE IF NOT EXIST Follow (
	follower TEXT NOT NULL, 
	followed TEXT NOT NULL,
	uploadDate TEXT NOT NULL DEFAULT "0000-01-01T00:00:00Z",
	PRIMARY KEY (follower, followed),
	FOREIGN KEY (follower) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE,
	FOREIGN KEY (followed) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE
	);`

// ----- PHOTO-RELATED: -----
// Note here that phrase, latitude, longitude could also be NULL.
var photo_table = `CREATE TABLE IF NOT EXIST Photo (
	photoid TEXT NOT NULL, 
	fixedUsername TEXT NOT NULL,
	filename BLOB NOT NULL,
	uploadDate TEXT NOT NULL DEFAULT "0000-01-01T00:00:00Z",
	phrase TEXT,
	numberLikes INTEGER NOT NULL DEFAULT 0,
	numberComments INTEGER NOT NULL DEFAULT 0,
	latitude FLOAT,
	longitude FLOAT,
	PRIMARY KEY (photoid, fixedUsername),
	FOREIGN KEY (fixedUsername) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE 
	);`

// Notice that here I declared (commentid,photoid, user) as PK. Commenter is not needed since it can make more than one comment.
var comment_table = `CREATE TABLE IF NOT EXIST Comment (
	commentid INTEGER AUTOINCREMENT, 
	photoid INTEGER NOT NULL,
	user TEXT NOT NULL, 
	phrase TEXT NOT NULL,
	commenter TEXT NOT NULL,
	uploadDate TEXT NOT NULL DEFAULT "0000-01-01T00:00:00Z",
	PRIMARY KEY (commentid, photoid, user),
	FOREIGN KEY (photoid) REFERENCES Photo (photoid) ON DELETE CASCADE,
	FOREIGN KEY (user) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE,
	FOREIGN KEY (commenter) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE
	);`

// Here the key (likeid, photoid, user), menas that likeid(fixedUsername) has put alike on the photoid photo of the user(fixedUsername).
var like_table = `CREATE TABLE IF NOT EXIST Like (
	likeid TEXT NOT NULL, 
	photoid INTEGER NOT NULL,
	user TEXT NOT NULL,
	uploadDate TEXT NOT NULL DEFAULT "0000-01-01T00:00:00Z",
	PRIMARY KEY (likeid, photoid, user),
	FOREIGN KEY (likeid) REFERENCES UserProfile (fixedUsername) ON DELETE CASCADE,
	FOREIGN KEY (photoid) REFERENCES Photo (photoid) ON DELETE CASCADE,
	FOREIGN KEY (user) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE,
	);`
