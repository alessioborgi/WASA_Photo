package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) CheckBanPresence(fixedUsernameBanner string, fixedUsernameBanned string) error {

	// Check whether there exists a Ban between fixedUsernameBanner and fixedUsernameBanned.
	var exists = 0
	err := db.c.QueryRow(`SELECT COUNT(fixedUsernameBanner) FROM Bans WHERE fixedUsernameBanner = ? AND fixedUsernameBanned = ?`, fixedUsernameBanner, fixedUsernameBanned).Scan(&exists)

	// Check for the error during the Query.
	if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Err: Unexpected Error during the Checking of the Ban!")
		return err
	} else if exists == 1 {
		// If no strange error during the Query occurs, and exists = 1, we already have the Ban Exists.
		log.Println("The Ban is present in the Database.")
		return Okay_Error_Inverse
	}

	// If we arrive here it means that the Ban is not present in the DB.
	return ErrBanDoesNotExist
}
