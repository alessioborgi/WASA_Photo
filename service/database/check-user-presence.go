package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) CheckUserPresence(username string) (string, error) {

	// Check whether the uuid that is requesting the action is the owner of the profile.
	var exists = 0
	err := db.c.QueryRow(`SELECT COUNT(fixedUsername) FROM Users WHERE username = ?`, username).Scan(&exists)

	// Check for the error during the Query.
	if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Err: Unexpected Error!")
		return "", err
	} else if exists == 1 {
		// If no strange error during the Query occurs, and exists = 1, we already have the user registered.
		// The User already Exists.

		var fixedUsername string
		errFixedUsername := db.c.QueryRow(`SELECT fixedUsername FROM Users WHERE username = ?`, username).Scan(&fixedUsername)
		if !errors.Is(errFixedUsername, nil) {
			log.Println("Err: Unexpected Error!")
			return "", err
		} else {

			// If we arrive here we have that the username has been correclty retrieved. We can therefore return it.
			log.Println("fixedUsername correctly retrieved from the Database.")
			return fixedUsername, nil
		}
	} else {

		// The User does not Exists.
		return "Not Exists", ErrUserDoesNotExist
	}
}
