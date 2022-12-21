package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) CheckPhotoPresence(photoid string, fixedUsername string) (string, error) {

	// Check whether the uuid that is requesting the action is the owner of the profile.
	var exists = 0
	err := db.c.QueryRow(`SELECT COUNT(photoid) FROM Photos WHERE photoid = ? AND fixedUsername = ?`, photoid, fixedUsername).Scan(&exists)

	// Check for the error during the Query.
	if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Err: Unexpected Error!")
		return "", err
	} else if exists == 1 {
		// If no strange error during the Query occurs, and exists = 1, we already have the photo saved.
		// The Photo already Exists.
		return photoid, Okay_Error_Inverse
	}

	// The Photo does not Exists.
	return "Not Exists", ErrPhotoDoesNotExist
}
