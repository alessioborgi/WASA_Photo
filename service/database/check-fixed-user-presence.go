package database

import (
	"database/sql"
	"log"
)

func (db *appdbimpl) CheckFixedUserPresence(fixedUsername string) (string, error) {

	// Check whether there exists a fixedUsername that is requesting the action is present..
	var exists = 0
	err := db.c.QueryRow(`SELECT COUNT(fixedUsername) FROM Users WHERE fixedUsername = ?`, fixedUsername).Scan(&exists)

	// Check for the error during the Query.
	if err != nil && err != sql.ErrNoRows {
		log.Println("Err: Unexpected Error!")
		return "", err
	} else if exists == 1 {
		// If no strange error during the Query occurs, and exists = 1, we already have the user registered.
		// The User already Exists.

		var username string
		errFixedUsername := db.c.QueryRow(`SELECT username FROM Users WHERE fixedUsername = ?`, fixedUsername).Scan(&username)
		if errFixedUsername != nil {
			log.Println("Err: Unexpected Error!")
			return "", err
		} else {

			// If we arrive here we have that the username has been correclty retrieved. We can therefore return it.
			log.Println("FixedUsername correctly Retrieved from the Database.")
			return username, Ok
		}
	} else {

		// The User does not Exists.
		return "Not Exists", ErrUserDoesNotExist
	}
}
