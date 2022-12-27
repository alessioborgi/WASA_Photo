package database

import (
	"database/sql"
	"errors"
	"log"
)

// If the return value is nil, we have that the Comment is present in the DB. Otherwise either it does not exists or it is an error.
func (db *appdbimpl) CheckCommentPresence(commentid string, fixedUsername string, photoid string, fixedUsernameCommenter string) error {

	// Check whether there exists a Comment.
	var exists = 0
	err := db.c.QueryRow(`SELECT COUNT(commentid) FROM Comments WHERE commentid = ? AND fixedUsername = ? AND photoid = ? AND commenterFixedUsername = ?`, commentid, fixedUsername, photoid, fixedUsernameCommenter).Scan(&exists)

	// Check for the error during the Query.
	if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Err: Unexpected Error!")
		return err
	} else if exists == 1 {

		// If no strange error during the Query occurs, and exists = 1, we already have the Comment Exists.
		log.Println("The Comment is present in the Database.")
		return nil
	}

	// If we arrive here it means that the Comment is not present in the DB.
	return ErrCommentDoesNotExist
}
