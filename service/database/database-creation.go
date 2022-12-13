package database

// ----- USER-RELATED: -----

// YYYY-MM-DD HH:MM:SS.SSS is the format to use then for date and date-time, respecting the the ISO8601 string format.
// IF NOT EXISTS used for cheecking whether the table we are creting is not present in the DB.
// DEFAULT is used for checking the fixing the default values.
// UNIQUE is used for saying that the value must be unique among all rows.
// Note that photoProfile and Biography coul also be NULL.
const user_table = `CREATE TABLE IF NOT EXISTS Users (
	fixedUsername TEXT NOT NULL PRIMARY KEY, 
	uuid TEXT NOT NULL UNIQUE,
	username TEXT NOT NULL UNIQUE,
	biography TEXT,
	dateOfCreation TEXT NOT NULL DEFAULT "0000-01-01T00:00:00Z",									
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
const ban_table = `CREATE TABLE IF NOT EXISTS Bans (
	fixedUsernameBanner TEXT NOT NULL, 
	fixedUsernameBanned TEXT NOT NULL,
	uploadDate TEXT NOT NULL DEFAULT "0000-01-01T00:00:00Z",
	Motivation TEXT NOT NULL DEFAULT "Spam",
	PRIMARY KEY (fixedUsernameBanner, fixedUsernameBanned),
	FOREIGN KEY (fixedUsernameBanner) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE,
	FOREIGN KEY (fixedUsernameBanned) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE
	);`

const follow_table = `CREATE TABLE IF NOT EXISTS Follows (
	fixedUsername TEXT NOT NULL, 
	fixedUsernameFollowing TEXT NOT NULL,
	uploadDate TEXT NOT NULL DEFAULT "0000-01-01T00:00:00Z",
	PRIMARY KEY (fixedUsername, fixedUsernameFollowing),
	FOREIGN KEY (fixedUsername) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE,
	FOREIGN KEY (fixedUsernameFollowing) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE
	);`

// ----- PHOTO-RELATED: -----
// Note here that phrase, latitude, longitude could also be NULL.
const photo_table = `CREATE TABLE IF NOT EXISTS Photos (
	photoid INTEGER NOT NULL, 
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

// Notice that here I declared (commentid,photoid, fixedUsername) as PK. Commenter is not needed since it can make more than one comment.
const comment_table = `CREATE TABLE IF NOT EXISTS Comments (
	commentid INTEGER NOT NULL, 
	photoid INTEGER NOT NULL,
	fixedUsername TEXT NOT NULL, 
	phrase TEXT NOT NULL,
	commenterFixedUsername TEXT NOT NULL,
	uploadDate TEXT NOT NULL DEFAULT "0000-01-01T00:00:00Z",
	PRIMARY KEY (commentid, photoid, fixedUsername),
	FOREIGN KEY (photoid) REFERENCES Photo (photoid) ON DELETE CASCADE,
	FOREIGN KEY (fixedUsername) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE,
	FOREIGN KEY (commenterFixedUsername) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE
	);`

// Here the key (likeid, photoid, fixedUsername), means that likeid(fixedUsername) has put a like on the photoid photo of the user(fixedUsername).
// LikeId = fixedUsername of the Liker
// fixedUsername = fixedusername of the person being liked.
const (
	like_table = `CREATE TABLE IF NOT EXISTS Likes (
	likeid TEXT NOT NULL, 
	photoid INTEGER NOT NULL,
	fixedUsername TEXT NOT NULL,
	uploadDate TEXT NOT NULL DEFAULT "0000-01-01T00:00:00Z",
	PRIMARY KEY (likeid, photoid, fixedUsername),
	FOREIGN KEY (likeid) REFERENCES UserProfile (fixedUsername) ON DELETE CASCADE,
	FOREIGN KEY (photoid) REFERENCES Photo (photoid) ON DELETE CASCADE,
	FOREIGN KEY (fixedUsername) REFERENCES UserProfle (fixedUsername) ON DELETE CASCADE
	);`
	query_presence_user    = `SELECT name FROM sqlite_master WHERE type='table' AND name='Users';`
	query_presence_ban     = `SELECT name FROM sqlite_master WHERE type='table' AND name='Bans';`
	query_presence_follow  = `SELECT name FROM sqlite_master WHERE type='table' AND name='Follows';`
	query_presence_photo   = `SELECT name FROM sqlite_master WHERE type='table' AND name='Photos';`
	query_presence_comment = `SELECT name FROM sqlite_master WHERE type='table' AND name='Comments';`
	query_presence_like    = `SELECT name FROM sqlite_master WHERE type='table' AND name='Likes';`

	delete_users    = `DROP TABLE Users;`
	delete_bans     = `DROP TABLE Bans;`
	delete_follows  = `DROP TABLE Follows;`
	delete_photos   = `DROP TABLE Photos;`
	delete_comments = `DROP TABLE Comments;`
	delete_likes    = `DROP TABLE Likes;`
	delete_alessio  = `DELETE FROM Users WHERE fixedUsername=alessio01`
)

var (
	database             = []string{user_table, ban_table, follow_table, photo_table, comment_table, like_table}
	query_table_presence = []string{query_presence_user, query_presence_ban, query_presence_follow, query_presence_photo, query_presence_comment, query_presence_like}
	delete_tables        = []string{delete_users, delete_bans, delete_follows, delete_photos, delete_comments, delete_likes}
)
