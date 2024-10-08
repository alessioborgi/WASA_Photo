/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

// Errors Declaration
var (
	ErrUserDoesNotExist     = errors.New("The User does not Exists!")
	ErrPhotoDoesNotExist    = errors.New("The Photo does not Exists!")
	ErrCommentDoesNotExist  = errors.New("The Comment does not Exists!")
	ErrUserNotAuthorized    = errors.New("The User not Authorized!")
	ErrUserProfileOwner     = errors.New("The User is the Profile Owner!")
	ErrInternalServerError  = errors.New("Internal Server Error!")
	ErrNoContent            = errors.New("There isn't any object you are searching for in the WASAPhoto Platform!")
	ErrBadRequest           = errors.New("The action you requested cannot be parsed due to a Bad Request!")
	ErrBanDoesNotExist      = errors.New("The Ban does not Exists!")
	ErrFollowDoesNotExist   = errors.New("The Follow does not Exists!")
	ErrLikeDoesNotExists    = errors.New("The Like does not Exists!")
	ErrCommentDoesNotExists = errors.New("The Comment does not Exists!")
)

// User Struct has been declared in the "db-struct-user.go" file.

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	// -----
	// MANDATORY
	// -----

	// SESSION:
	// DoLogin() creates a new Username given in input a Username. If does not already exists and returns a uuid, or,  if it already exists, simply returns a uuid.
	DoLogin(username string) (string, error, string)

	// PARTICULAR USER:
	// (Security Required: Needs Uuid of the action requester).
	// SetUser(), given the fixedUsername in input together with a newUsername, updates the User's Username.
	SetUser(username string, user User, uuid string) (string, error)

	// USER's PHOTO COLLECTION:
	// (Security Required: Needs Uuid of the action requester).
	// UploadPhoto() creates a new User's Photo(Post) in the database, given in input the Photo Object. It returns an Photo Object.
	UploadPhoto(username string, photo Photo, uuid string) error

	// (Security Required: Needs Uuid of the action requester).
	// DeletePhoto() removes a User's Photo given the username and the photoId in input. Returns the filename path and an error.
	DeletePhoto(username string, photoid string, uuid string) (string, error)

	// PARTICULAR FOLLOW:
	// (Security Required: Needs Uuid of the action requester).
	// FollowUser() creates a new User's Follow in the database, given in input the Follow Object. It returns a Follow Object.
	FollowUser(username string, usernameFollowing, uuid string) (string, error)

	// (Security Required: Needs Uuid of the action requester).
	// UnfollowUser() removes a User's Follow given the fixedUsername, and the FollowindId(i.e., the fixedUsername of the Person that the fixedUsername wants to delete from the following list).
	UnfollowUser(username string, usernameFollowing string, uuid string) error

	// PARTICULAR BAN:
	// (Security Required: Needs Uuid of the action requester).
	// BanUser() creates a new User's Ban in the database, given in input the username of the profile owner and the username of the person I want to ban. It returns nothing.
	BanUser(username string, usernameBanned string, uuid string) (string, error)

	// (Security Required: Needs Uuid of the action requester).
	// UnbanUser() removes a User's Ban given the fixedUsername, and the BanId(i.e., the fixedUsername of the Banned Person).
	UnbanUser(username string, usernameBanned string, uuid string) error

	// PARTICULA USER:
	// (Security Required: Needs Uuid of the action requester).
	// GetUserProfile() returns the User Profile requested given in input the fixedUsername.
	GetUserProfile(username string, uuid string) (User, error)

	// PARTICULAR LIKE:
	// (Security Required: Needs Uuid of the action requester).
	// LikePhoto() creates a new User's Photo Like in the database, given in input the Like Object. It returns a Like Object.
	LikePhoto(username string, photoid string, usernameLiker string, uuid string) (string, error)

	// (Security Required: Needs Uuid of the action requester).
	// UnlikePhoto() removes a User's Photo Like given the fixedUsername, the photoId and the fixedUsername of the Liker in input.
	UnlikePhoto(username string, photoid string, usernameLiker string, uuid string) error

	// USER's PHOTO COMMENTS COLLECTION:
	// (Security Required: Needs Uuid of the action requester).
	// CommentPhoto() creates a new User's Photo Comment in the database, given in input the Comment Object. It returns a Commentid.
	CommentPhoto(username string, photoid string, comment Comment, uuid string) (int, error)

	// (Security Required: Needs Uuid of the action requester).
	// UncommentPhoto() removes a User's Photo Comment given the fixedUsername, the photoId and the commentId in input.
	UncommentPhoto(username string, photoid string, commentid string, uuid string) error

	// USER STREAM:
	// (Security Required: Needs Uuid of the action requester).
	// GetMyStream() returns a list of Photos pertaining to the User's following list. We provide in input a fixedUsername.
	GetMyStream(username string, uuid string) ([]Photo, error)

	// -----
	// OPTIONAL
	// -----

	// USERS COLLECTION:
	// (Security Required: Needs Uuid of the action requester).
	// GetUsers() returns the list of fixedUsername.
	GetUsers(uuid string) ([]string, error)

	// PARTICULAR USER:
	// (Security Required: Needs Uuid of the action requester).
	// DeleteUsername() removes the User given the fixedUsername in input.
	DeleteUser(username string, uuid string) error

	// (Security Required: Needs Uuid of the action requester).
	// GetPhotos() returns the list of Photos of a given user, given in input a fixedUsername.
	GetPhotos(username string, uuid string) ([]Photo, error)

	// PARTICULAR PHOTO:
	// (Security Required: Needs Uuid of the action requester).
	// GetPhoto() return a User's Photo, given the fixedUsername and the photoid in input.
	GetPhoto(username string, photoid string, uuid string) (Photo, error)

	// USER's PHOTO COMMENTS COLLECTION:
	// (Security Required: Needs Uuid of the action requester).
	// GetPhotoComments() returns the list of Photos's Comments of a given User Photo, given in input a fixedUsername, and the photoId.
	GetPhotoComments(username string, photoid string, uuid string) ([]Comment, error)

	// USER's PHOTO LIKES:
	// (Security Required: Needs Uuid of the action requester).
	// GetPhotoLikes() returns the list of Photos's Likes of a given User Photo, given in input a fixedUsername, and the photoId.
	GetPhotoLikes(username string, photoid string, uuid string) ([]string, error)

	// USER's BANS COLLECTION:
	// (Security Required: Needs Uuid of the action requester).
	// GetBannedUsers() returns the list of User's Bans, given in input a username.
	GetBannedUsers(username string, uuid string) ([]string, error)

	// USER's FOLLOWERS COLLECTION:
	// (Security Required: Needs Uuid of the action requester).
	// GetFollowers() returns the list of User's Followers(Follow Objects), given in input a fixedUsername.
	GetFollowers(username string, uuid string) ([]string, error)

	// USER's FOLLOWINGS COLLECTION:
	// (Security Required: Needs Uuid of the action requester).
	// GetFollowing() returns the list of User's Followings(Follow Objects), given in input a fixedUsername.
	GetFollowings(username string, uuid string) ([]string, error)

	// -----
	// SPECIAL
	// -----

	// GetLastFixedUsername is used for getting the last fixedUsername+1, to be used in the DB.
	GetLastFixedUsername() (string, error)

	// GetUsername is useful for getting the Username given the fixedUsername in input.
	GetUsername(fixedUsername string) (string, error)

	// GetLastPhotoId is useful for getting back the lastPhotoId.
	GetLastPhotoId(username string) (int64, error)

	// CheckUserPresence given in input the username, returns the fixedUsername.
	CheckUserPresence(username string) (string, error)

	// Ping checks whether the database is available or not (in that case, an error will be returned).
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("Database is required when building a AppDatabase")
	}

	// Used to Turn On the foreign Key Mechanism.
	_, err1 := db.Exec(turn_on_fk)
	if !errors.Is(err1, nil) {
		log.Println("Error Encountered during the FK Turning on")
	} else {
		log.Println("FK correctly Turned On.")
	}

	// This code is only used during development if we do some change on the database schema.
	// for i := 0; i < len(delete_tables); i++ {
	// 	_, err := db.Exec(delete_tables[i])
	// 	if !errors.Is(err, nil) {
	// 		log.Println("Error Encountered during the Table Deletion", i)
	// 	} else {
	// 		log.Println("Table", i, "deleted correctly.")
	// 	}
	// }

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	for i := 0; i < len(database); i++ {

		// Check whether, for every table, we have the Table.
		err := db.QueryRow(query_table_presence[i]).Scan(&tableName)
		if errors.Is(err, sql.ErrNoRows) {

			// If the table is not present, create it.
			table_creation := database[i]
			_, err = db.Exec(table_creation)
			if !errors.Is(err, nil) {
				log.Println("Error in Creating the Database Structure of the table: ", i)
				return nil, fmt.Errorf("Error in Creating the Database Structure: %w", err)
			} else {
				log.Println("Creation of the table number: ", i, "succeeded!")
			}
		} else {
			log.Println("The Table", i, "is already present!")
		}
	}

	// Triggers insertion.
	var trigger string
	for i := 0; i < len(triggers); i++ {

		// Check whether, for every Trigger, we have it.
		err := db.QueryRow(query_trigger_presence[i]).Scan(&trigger)

		if errors.Is(err, sql.ErrNoRows) {

			// If the table is not present, create it.
			trigger_creation := triggers[i]
			_, err = db.Exec(trigger_creation)
			if !errors.Is(err, nil) {
				log.Println("Error in Creating the Trigger: ", i)
				return nil, fmt.Errorf("Error in Creating the Trigger: %w", err)
			} else {
				log.Println("Creation of the Trigger number: ", i, "succeeded!")
			}
		} else {
			log.Println("The Trigger", i, "is already present!")
		}

	}

	// ADMIN USER PROFILE CREATION (myself): alessioborgi01
	// This is done for Development Purpose, in such a way to leave always a sort of backdoor.

	// First check whether there are any other users in the table.
	var exists = 0
	err := db.QueryRow(`SELECT COUNT(fixedUsername) FROM Users`).Scan(&exists)

	// Check for the error during the Query.
	if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		log.Fatalf("Unexpected Error during the Query of the DB!")
		return &appdbimpl{
			c: db,
		}, err
	} else if exists == 0 {

		// If no user is in the Users Table, go beyond and add the Admin User Profile(alessioborgi01) and others.
		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u18", "18181818-1818-1818-1818-181818181818", "Rosalinda", "", "I am Rosalinda Bros", now, 0, 0, 0, "Rosalinda", "Bros", "2000-01-01", "rosalinda.1952442@studenti.uniroma1.it", "Italian", "female")
		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u17", "17171717-1717-1717-1717-171717171717", "Peach", "", "I am Peach Bros", now, 0, 0, 0, "Peach", "Bros", "2000-01-01", "peach.1952442@studenti.uniroma1.it", "Italian", "female")
		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u16", "16161616-1616-1616-1616-161616161616", "Yoshi", "", "I am Yoshi Bros", now, 0, 0, 0, "Yoshi", "Bros", "2000-01-01", "yoshi.1952442@studenti.uniroma1.it", "Italian", "male")
		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u15", "15151515-1515-1515-1515-151515151515", "Mario", "", "I am Mario Bros", now, 0, 0, 0, "Mario", "Bros", "2000-01-01", "mario.1952442@studenti.uniroma1.it", "Italian", "male")
		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u14", "14141414-1414-1414-1414-141414141414", "TipoTimido", "", "I am Tipo Timido", now, 0, 0, 0, "Tipo", "Timido", "2000-01-01", "tipo.1952442@studenti.uniroma1.it", "Italian", "male")
		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u13", "13131313-1313-1313-1313-131313131313", "Lakitu", "", "I am Lakitu Bros", now, 0, 0, 0, "Lakitu", "Bros", "2000-01-01", "lakitu.1952442@studenti.uniroma1.it", "Italian", "male")
		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u12", "12121212-1212-1212-1212-121212121212", "Koopa", "", "I am Koopa Bros", now, 0, 0, 0, "Koopa", "Bros", "2000-01-01", "koopa.1952442@studenti.uniroma1.it", "Italian", "male")
		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u11", "10101010-1010-1010-1010-101010101010", "DonkeyKong", "", "I am Donkey Kong", now, 0, 0, 0, "Donkey", "Kong", "2000-01-01", "donkey.1952442@studenti.uniroma1.it", "Italian", "male")
		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u10", "99999999-9999-9999-9999-9999999999999999", "Bowser", "", "I am Bowser Bros", now, 0, 0, 0, "Bowser", "Bros", "2000-01-01", "bowser.1952442@studenti.uniroma1.it", "Italian", "male")
		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u9", "88888888-8888-8888-8888-888888888888", "Wario", "", "I am Wario Bros", now, 0, 0, 0, "Wario", "Bros", "2000-01-01", "wario.1952442@studenti.uniroma1.it", "Italian", "male")
		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u8", "77777777-7777-7777-7777-777777777777", "Luigi", "", "I am Luigi Bros", now, 0, 0, 0, "Luigi", "Bros", "2000-01-01", "luigi.1952442@studenti.uniroma1.it", "Italian", "male")

		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u7", "66666666-6666-6666-6666-666666666666", "chicca78", "", "I am the WASAPhoto Owner's Aunt", now, 0, 0, 0, "Francesca", "Borgi", "1978-10-07", "mario.1952442@studenti.uniroma1.it", "Italian", "male")
		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u6", "55555555-5555-5555-5555-555555555555", "samantha01", "", "I am the WASAPhoto Owner's Girlfriend", now, 0, 0, 0, "Samantha", "Sorrentino", "2001-10-26", "samantha.1952442@studenti.uniroma1.it", "Italian", "female")
		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u5", "44444444-4444-4444-4444-444444444444", "aurora46", "", "I am the WASAPhoto Owner's Grandma", now, 0, 0, 0, "Aurora", "Melis", "1946-08-27", "aurora.1952442@studenti.uniroma1.it", "Italian", "female")
		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u4", "33333333-3333-3333-3333-333333333333", "alice05", "", "I am the WASAPhoto Owner's Sister", now, 0, 0, 0, "Alice", "Borgi", "2005-12-16", "alice.1952442@studenti.uniroma1.it", "Italian", "female")
		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u3", "22222222-2222-2222-2222-222222222222", "andrea71", "", "I am the WASAPhoto Owner's Dad", now, 0, 0, 0, "Andrea", "Borgi", "1971-09-06", "andrea.1952442@studenti.uniroma1.it", "Italian", "male")
		_, _ = db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u2", "11111111-1111-1111-1111-111111111111", "anna69", "", "I am the WASAPhoto Owner's Mum", now, 0, 0, 0, "Anna", "Mauti", "1969-07-29", "anna.1952442@studenti.uniroma1.it", "Italian", "female")
		_, errCretion := db.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"u1", "00000000-0000-0000-0000-000000000000", "alessioborgi01", "", "I am the WASAPhoto Owner", now, 0, 0, 0, "Alessio", "Borgi", "2001-04-17", "borgi.1952442@studenti.uniroma1.it", "Italian", "male")

		if !errors.Is(errCretion, nil) {
			log.Fatalf("Error During Alessio's Account Creation")
			return &appdbimpl{
				c: db,
			}, errCretion
		} else {

			// If no error occurs, User Profile Creation.
			log.Println("WASAPhoto's Owner Account Created: alessioborgi01")
			return &appdbimpl{
				c: db,
			}, nil
		}
	} else {

		// We arrive here if we have already the Admin User to be Present in the DB.
		log.Println("User Admin already Created!")
		return &appdbimpl{
			c: db,
		}, nil
	}
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
