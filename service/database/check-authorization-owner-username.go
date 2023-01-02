package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) CheckAuthorizationOwnerUsername(username string, uuid string) (string, error) {

	// First check whether there exists the Uuid in the DataBase.
	exists, err := db.CheckAuthorizationPresence(uuid)
	log.Println("The Uuid is:", exists, "in the DB.")

	// Check for the error during the Query.
	if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {

		// Unexpected Error encountered during the query retrieval.
		log.Println("Err: Unexpected Error!")
		return "", err
	} else if exists == PRESENT {

		// Here we have that the Uuid inserted, is one present in the DB.
		// Check whether the uuid that is requesting the action is the owner of the profile.
		var auth bool
		err := db.c.QueryRow(`SELECT uuid == ? FROM Users WHERE username == ?`, uuid, username).Scan(&auth)

		// Check for the error during the Query.
		if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {

			return "", err
		} else if errors.Is(err, sql.ErrNoRows) {

			// It enters here only if the username we are providing does not exists.
			log.Println("The username is not present in the DB.")
			return NOTVALID, ErrUserDoesNotExist
		} else {

			// Returning "Authorized" if it is the owner, "Not Authorized" otherwise.
			if auth {

				// Here we have that the person is requesting the action is the account owner, thus it is authorized.
				log.Println("User recognized as the Account Owner.")
				return AUTHORIZED, nil
			} else {

				// Here we have that the person is requesting the action is NOT the account owner, thus it is NOT authorized.
				log.Println("User NOT recognized as the Account Owner.")
				return NOTAUTHORIZED, nil
			}
		}
	} else {

		// Here we have that the Uuid inserted, is NOT one present in the DB.
		log.Println("Err: Inserted Uuid is not one Present in the DB!")
		return NOTVALID, nil
	}
}
