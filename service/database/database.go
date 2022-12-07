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
)

// Error Function.
var ErrUserDoesNotExist = errors.New("User does not Exists!")
var ErrPhotoDoesNotExist = errors.New("Photo does not Exists!")
var ErrCommentDoesNotExist = errors.New("Comment does not Exists!")
var ErrUserNotAuthorized = errors.New("User not Authorized!")

// User Struct has been declared in the "db-struct-user.go" file.

// -----
// IDEAS: I think i should add the uuid of the User which is doing all the things in all the functions except for the GetUsers().
// DOUBTS: Doubts on Session: Shouuld it return uuid or fixedUsername?. I think uuid!
// TO DO: ADD Uuid at every input Function except DoLogin.
// -----

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	// -----
	// MANDATORY
	// -----

	// SESSION:
	//DoLogin() creates a new Username given in input a Username. If does not already exists and returns a uuid, or,  if it already exists, simply returns a uuid.
	//DoLogin(username string) (string, error)
	// DoLogin(user User) (string, error)	//Maybe in this way?

	// PARTICULAR USER:
	//(Security Required: Needs Uuid of the action requester).
	// SetMyUsername(), given the fixedUsername in input together with a newUsername, updates the User's Username.
	SetMyUsername(fixedUsername string, newUsername string, uuid string) error

	// USER's PHOTO COLLECTION:
	//(Security Required: Needs Uuid of the action requester).
	// UploadPhoto() creates a new User's Photo(Post) in the database, given in input the Photo Object. It returns an Photo Object.
	//UploadPhoto(photo Photo) (Photo, error)

	//(Security Required: Needs Uuid of the action requester).
	// DeletePhoto() removes a User's Photo given the fixedUsername and the photoId in input.
	DeletePhoto(fixedUsername string, photoId int, uuid string) error

	// PARTICULAR FOLLOW:
	//(Security Required: Needs Uuid of the action requester).
	// FollowUser() creates a new User's Follow in the database, given in input the Follow Object. It returns a Follow Object.
	FollowUser(follow Follow, uuid string) (Follow, error)

	//(Security Required: Needs Uuid of the action requester).
	// UnfollowUser() removes a User's Follow given the fixedUsername, and the FollowindId(i.e., the fixedUsername of the Person that the fixedUsername wants to delete from the following list).
	UnfollowUser(fixedUsername string, followingId string, uuid string) error

	// PARTICULAR BAN:
	//(Security Required: Needs Uuid of the action requester).
	// BanUser() creates a new User's Ban in the database, given in input the Ban Object. It returns a Ban Object.
	BanUser(ban Ban, uuid string) (Ban, error)

	//(Security Required: Needs Uuid of the action requester).
	// UnbanUser() removes a User's Ban given the fixedUsername, and the BanId(i.e., the fixedUsername of the Banned Person).
	UnbanUser(fixedUsername string, banId string, uuid string) error

	// PARTICULA USER:
	//(Security Required: Needs Uuid of the action requester).
	// GetUserProfile() returns the User Profile requested given in input the fixedUsername.
	GetUserProfile(fixedUsername string, uuid string) (User, error)

	// USER STREAM:
	//(Security Required: Needs Uuid of the action requester).
	// GetMyStream() returns a list of Photos pertaining to the User's following list. We provide in input a fixedUsername.
	GetMyStream(fixedUsername string, uuid string) ([]Photo, error)

	// PARTICULAR LIKE:
	//(Security Required: Needs Uuid of the action requester).
	// LikePhoto() creates a new User's Photo Like in the database, given in input the Like Object. It returns a Like Object.
	//LikePhoto(like Like, uuid string) (Like, error)

	//(Security Required: Needs Uuid of the action requester).
	// UnlikePhoto() removes a User's Photo Like given the fixedUsername, the photoId and the fixedUsername of the Liker in input.
	UnlikePhoto(fixedUsername string, photoId int, fixedUsernameLiker string, uuid string) error

	// USER's PHOTO COMMENTS COLLECTION:
	//(Security Required: Needs Uuid of the action requester).
	// CommentPhoto() creates a new User's Photo Comment in the database, given in input the Comment Object. It returns a Comment Object.
	//CommentPhoto(comment Comment, uuid string) (Comment, error)

	//(Security Required: Needs Uuid of the action requester).
	// UncommentPhoto() removes a User's Photo Comment given the fixedUsername, the photoId and the commentId in input.
	UncommentPhoto(fixedUsername string, photoId int, commentId int, uuid string) error

	// -----
	// OPTIONAL
	// -----

	// USERS COLLECTION:
	//(Security Required: Needs Uuid of the action requester).
	// GetUsers() returns the list of fixedUsername.
	GetUsers(uuid string) ([]string, error)

	// PARTICULAR USER:
	//(Security Required: Needs Uuid of the action requester).
	// DeleteUsername() removes the User given the fixedUsername in input.
	DeleteUsername(fixedUsername string, uuid string) error

	//(Security Required: Needs Uuid of the action requester).
	// GetPhotos() returns the list of Photos of a given user, given in input a fixedUsername.
	GetPhotos(fixedUsername string, uuid string) ([]Photo, error)

	// PARTICULAR PHOTO:
	//(Security Required: Needs Uuid of the action requester).
	// SetPhoto() updates a User's Photo, replacing it with the new value of the Phrase in the argument, in addition to a fixedUsername of the User and the PhotoId.
	SetPhoto(fixedUsername string, photoId int, newPhrase string, uuid string) error

	// USER's PHOTO COMMENTS COLLECTION:
	//(Security Required: Needs Uuid of the action requester).
	// GetPhotoComments() returns the list of Photos's Comments of a given User Photo, given in input a fixedUsername, and the photoId.
	GetPhotoComments(fixedUsername string, photoId int, uuid string) ([]Comment, error)

	// PARTICULAR COMMENT:
	//(Security Required: Needs Uuid of the action requester).
	// SetComment(), given the fixedUsername in input together with a photoId, a commentId and a newComment(Phrase), updates the User's Username.
	SetComment(fixedUsername string, photoId int, commentId int, newComment string, uuid string) error

	// USER's PHOTO LIKES:
	//(Security Required: Needs Uuid of the action requester).
	// GetPhotoLikes() returns the list of Photos's Likes of a given User Photo, given in input a fixedUsername, and the photoId.
	GetPhotoLikes(fixedUsername string, photoId int, uuid string) ([]Like, error)

	// USER's BANS COLLECTION:
	//(Security Required: Needs Uuid of the action requester).
	// GetBannedUsers() returns the list of User's Bans, given in input a fixedUsername.
	GetBannedUsers(fixedUsername string, uuid string) ([]Ban, error)

	// USER's FOLLOWERS COLLECTION:
	//(Security Required: Needs Uuid of the action requester).
	// GetFollowers() returns the list of User's Followers(Follow Objects), given in input a fixedUsername.
	GetFollowers(fixedUsername string, uuid string) ([]Follow, error)

	// USER's FOLLOWINGS COLLECTION:
	//(Security Required: Needs Uuid of the action requester).
	// GetFollowing() returns the list of User's Followings(Follow Objects), given in input a fixedUsername.
	GetFollowing(fixedUsername string, uuid string) ([]Follow, error)

	// -----
	// SPECIAL
	// -----

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

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	for i := 0; i < len(database); i++ {

		//Check whether, for every table, we have the Table.
		err := db.QueryRow(query_table_presence[i]).Scan(&tableName)
		if errors.Is(err, sql.ErrNoRows) {

			// If the table is not present, create it.
			table_creation := database[i]
			_, err = db.Exec(table_creation)
			if err != nil {
				fmt.Println("Error in Creating the Database Structure of the table: ", i)
				return nil, fmt.Errorf("Error in Creating the Database Structure: %w", err)
			} else {
				fmt.Println("Creation of the table number: ", i, "succeeded!")
			}
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
