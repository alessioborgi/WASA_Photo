package api

import (
	"errors"

	"github.com/alessioborgi/WASA_Photo/service/database"
)

type Like struct {
	UsernameLiker string `json:"likeid"`        // This corresponds to the LikeId = UsernameLiker.
	PhotoId       int64  `json:"photoid"`       //
	Username      string `json:"fixedUsername"` //Primary & Foreign Key (PK, FK).
}

// ----- FINAL LIKE FUNCTION -----

// Function Method used to check for the User Validity.
func ValidLike(like Like) bool {
	return regex_username.MatchString(like.UsernameLiker) &&
		like.PhotoId >= 0 &&
		regex_username.MatchString(like.Username)
}

// -----                -----

// Function for handling the population of the struct with data from the DB.
func (l *Like) FromDatabase(like database.Like, db database.AppDatabase) error {
	l.UsernameLiker, err2 = db.GetUsername(like.Likeid)
	if !errors.Is(err2, nil) {
		return err2
	}
	l.PhotoId = like.PhotoId
	l.Username, err1 = db.GetUsername(like.FixedUsername)
	if !errors.Is(err1, nil) {
		return err1
	}
	return nil
}

// ToDatabase returns the User in a Database-Compatible Representation.

func (l *Like) ToDatabase(db database.AppDatabase) database.Like {
	return database.Like{
		// Likeid:  l.UsernameLiker,
		PhotoId: l.PhotoId,
		// FixedUsername: l.Username,
	}
}
