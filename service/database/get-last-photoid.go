package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) GetLastPhotoId(username string) (int64, error) {

	// Check whether there exists first the fixedUsername that is requesting the action.
	fixedUsername, errUsername := db.CheckUserPresence(username)

	// Check whether theUsername I am trying to update, does not exists.
	if errors.Is(errUsername, ErrUserDoesNotExist) {
		log.Println("Err: The fixedUsername, does not exists.")
		return 0, ErrUserDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errUsername, nil) && !errors.Is(errUsername, Ok) {
		log.Println("Err: Strange error during the Check of User Presence")
		return 0, errUsername
	}

	// If no strange error during the Query occurs, and the User exists, we already have the user registered.
	// We can proceed to get the last photoId.
	var photoid_existence = 0

	errPhotoId := db.c.QueryRow(`SELECT COUNT(photoid) FROM Photos WHERE fixedUsername = ?`, fixedUsername).Scan(&photoid_existence)
	if errors.Is(errPhotoId, sql.ErrNoRows) {

		// If we have no rows, we return that the photoid must be 1
		log.Println("No photos yet. Inserting First Photo!")
		return 1, nil
	} else if !errors.Is(errPhotoId, nil) {

		// If we encounter any other type of error, return error.
		log.Println("Err: Unexpected Error!")
		return 0, errPhotoId
	} else {

		// If we arrive here we have that the photoid has been correclty retrieved. We can therefore return the photoid+1.
		log.Println("photoid correctly retrieved from the Database.")
		return int64(photoid_existence) + 1, nil
	}

}
