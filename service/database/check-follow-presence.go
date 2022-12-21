package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) CheckFollowPresence(fixedUsername string, fixedUsernameFollowing string) error {

	// Check whether there exists a Follow between fixedUsername and fixedUsernameFollowing.
	var exists = 0
	err := db.c.QueryRow(`SELECT COUNT(fixedUsernameFollowing) FROM Follows WHERE fixedUsername = ? AND fixedUsernameFollowing = ?`, fixedUsername, fixedUsernameFollowing).Scan(&exists)

	// Check for the error during the Query.
	if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Err: Unexpected Error!")
		return err
	} else if exists == 1 {

		// If no strange error during the Query occurs, and exists = 1, we already have the Follow Exists.
		log.Println("The Follow is present in the Database.")
		return Okay_Error_Inverse
	}

	// If we arrive here it means that the Ban is not present in the DB.
	return ErrFollowDoesNotExist
}
