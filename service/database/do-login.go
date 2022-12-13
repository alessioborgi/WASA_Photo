package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
)

var now = time.Now()

func (db *appdbimpl) DoLogin(username string) (string, error) {

	// First check whether there exists a User with the inserted Username.
	var exists = 0
	err := db.c.QueryRow(`SELECT COUNT(fixedUsername) FROM Users WHERE username == ?`, username).Scan(&exists)

	// Check for the error during the Query.
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("Unexpected Error during the Query of the DB!")
		return "", err
	} else if exists == 1 {
		// USER PROFILE LOGIN:
		// If no strange error occurs during the Query, and exists = 1, we already have the user registered.
		// The User already Exists.
		log.Println("The User already exists!")
		var saved_uuid string

		// Uuid Retrieval.
		err := db.c.QueryRow(`SELECT uuid FROM Users WHERE username == ?`, username).Scan(&saved_uuid)
		if err != nil {
			log.Fatalf("Failed to Retrieve UUID from the DB")
			return "", err
		} else {
			log.Println("Uuid Retrieval Succeeded from the DB!")
			return saved_uuid, Ok
		}
	} else {

		// USER PROFILE CREATION:
		// The User deos not Exists.
		log.Println("The User does not exists!")

		//Uuid Generation.
		var uuid = uuid.Must(uuid.NewV4())
		if err != nil && err != sql.ErrNoRows {
			log.Fatalf("Failed to generate UUID: %v", err)
		} else {
			log.Println("Generated UUID", uuid)
		}

		// Actual User insertion in the DB. Insertion of the actual uuid, username and (after), update the fixedUsername.
		// The rest of the User is completely Standard, in such a way to have that the user is not obliged to add nothing else.
		res, errCretion := db.c.Exec(`INSERT INTO Users (fixedUsername, uuid, username, biography, dateOfCreation, numberOfPhotos, totNumberLikes, totNumberComments, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			uuid.String(), uuid.String(), username, "", now, 0, 0, 0, 0, 0, "", "", "1900-01-01", "surname.matriculation@studenti.uniroma1.it", "", "do not specify")
		if errCretion != nil {
			log.Fatalf("Error During Creation")
			return "Error", errCretion
		} else {

			// User Created Successfully.
			log.Println("User Creation Succeeded!")

			// User's fixedUsername Update.
			lastInsertID, err := res.LastInsertId()
			if err != nil {
				log.Fatalf("User fixedUsername retrieval Error")
				return "", err
			} else {
				log.Println("User fixedUsername retrieval Succedeed.")
				last_id := lastInsertID
				last := "u" + strconv.Itoa(int(last_id))
				fmt.Println("The lastID is:", last)
				_, errUpdate := db.c.Exec(`UPDATE Users SET fixedUsername=? WHERE username = ?`, last, username)
				if errUpdate != nil {
					log.Fatalf("Error During Updatating")
					return "", errUpdate
				} else {
					log.Println("fixedUsername Update Succeeded")
					return uuid.String(), Created
				}
			}
		}
	}
}
