package database

import (
	"database/sql"
	"log"
)

const (
	AUTHORIZED    = "Authorized"
	NOTAUTHORIZED = "Not Authorized"
	NOTVALID      = "Not Valid"
	PRESENT       = "Present"
	NOTPRESENT    = "Not Present"
)

func (db *appdbimpl) CheckAuthorizationOwner(fixedUsername string, uuid string) (string, error) {

	// First check whether there exists the Uuid in the DataBase.
	exists, err := db.CheckAuthorizationPresence(uuid)
	log.Println("The Uuid is:", exists, "in the DB.")

	// Check for the error during the Query.
	if err != nil && err != sql.ErrNoRows {

		// Unexpected Error encountered during the query retrieval.
		log.Fatalf("Unexpected Error!")
		return "", err
	} else if exists == PRESENT {

		// Here we have that the Uuid inserted, is one present in the DB.
		// Check whether the uuid that is requesting the action is the owner of the profile.
		var auth bool
		err := db.c.QueryRow(`SELECT uuid == '?' FROM Users	WHERE fixedUsername == '?'`, uuid, fixedUsername).Scan(&auth)

		// Check for the error during the Query.
		if err != nil {
			return "", err
		} else {

			// Returning "Authorized" if it is the owner, "Not Authorized" otherwise.
			if auth == true {

				// Here we have that the person is requesting the action is the account owner, thus it is authorized.
				log.Println("User recognized as the Account Owner.")
				return AUTHORIZED, nil
			} else {

				// Here we have that the person is requesting the action is NOT the account owner, thus it is NOT authorized.
				log.Fatalf("User NOT recognized as the Account Owner.")
				return NOTAUTHORIZED, nil
			}
		}
	} else {

		// Here we have that the Uuid inserted, is NOT one present in the DB.
		log.Fatalf("Inserted Uuid is not one Present in the DB!")
		return NOTVALID, nil
	}
}

func (db *appdbimpl) CheckAuthorizationPresence(uuid string) (string, error) {

	// First check whether there exists the Uuid in the DataBase.
	var exists = 0
	err := db.c.QueryRow(`SELECT COUNT(fixedUsername) FROM Users WHERE uuid == ?`, uuid).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {

		// Unexpected Error encountered during the query retrieval.
		log.Fatalf("Unexpected Error!")
		return "", err
	} else if exists == 1 {
		return PRESENT, nil
	} else {
		return NOTPRESENT, nil
	}
}
