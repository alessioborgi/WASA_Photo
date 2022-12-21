package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) CheckLikePresence(fixedUsername string, photoid string, likeid string) error {

	// Check whether there exists a Like between likeid(fixedUsernameLiker) and fixedUsername, in photid.
	var exists = 0
	err := db.c.QueryRow(`SELECT COUNT(likeid) FROM Likes WHERE fixedUsername = ? AND photoid = ? AND likeid = ?`, fixedUsername, photoid, likeid).Scan(&exists)

	// Check for the error during the Query.
	if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Err: Unexpected Error!")
		return err
	} else if exists == 1 {

		// If no strange error during the Query occurs, and exists = 1, we already have the Follow Exists.
		log.Println("The Like is present in the Database.")
		return Okay_Error_Inverse
	}

	// If we arrive here it means that the Ban is not present in the DB.
	return ErrLikeDoesNotExists
}
