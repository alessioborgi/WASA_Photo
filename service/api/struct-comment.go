package api

import (
	"errors"

	"github.com/alessioborgi/WASA_Photo/service/database"
)

type DateTime string

type Comment struct {
	Commentid         int64    `json:"commentid"`
	PhotoId           int64    `json:"photoid"`
	Username          string   `json:"username"`          // Note that here there is no FixedUsername as in the DB Struct, but the Username.
	CommenterUsername string   `json:"commenterUsername"` // Note that here there is no CommenterFixedUsername as in the DB Struct, but the CommenterUsername.
	Phrase            string   `json:"phrase"`
	UploadDate        DateTime `json:"uploadDate"`
}

// ----- FINAL COMMENT FUNCTION -----

// Function Method used to check for the User Validity.
func ValidComment(comment Comment) bool {
	return comment.Commentid >= 0 &&
		comment.PhotoId >= 0 &&
		regex_username.MatchString(comment.Username) &&
		regex_username.MatchString(comment.CommenterUsername) &&
		len(comment.Phrase) >= 5 && len(comment.Phrase) <= 1000
}

// -----                -----

var err1 error
var err2 error

// Function for handling the population of the struct with data from the DB.
func (c *Comment) FromDatabase(comment database.Comment, db database.AppDatabase) error {
	c.Commentid = comment.Commentid
	c.PhotoId = comment.PhotoId
	c.Username, err1 = db.GetUsername(comment.FixedUsername)
	if !errors.Is(err1, nil) {
		return err1
	}
	c.CommenterUsername, err2 = db.GetUsername(comment.CommenterFixedUsername)
	if !errors.Is(err2, nil) {
		return err2
	}
	c.Phrase = comment.Phrase
	c.UploadDate = DateTime(comment.UploadDate)
	return nil
}

// ToDatabase returns the User in a Database-Compatible Representation.

func (c *Comment) ToDatabase(db database.AppDatabase) database.Comment {
	return database.Comment{
		// Commentid:              c.Commentid,
		// PhotoId:                c.PhotoId,
		// FixedUsername:          c.FixedUsername,
		// CommenterFixedUsername: c.CommenterFixedUsername,
		Phrase: c.Phrase,
		// UploadDate:             string(c.UploadDate),
	}
}
