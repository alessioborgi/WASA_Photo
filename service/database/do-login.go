package database

import (
	"errors"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

var now = time.Now().Format(time.RFC3339)

func (db *appdbimpl) DoLogin(username string) (string, error, string) {

	// First check whether there exists a User with the inserted Username.
	fixedUsername, errUserPresence := db.CheckUserPresence(username)

	// Check whether the fixedUsername  is already in the DB. If the fixedUsername is != ("" AND "Not Exists") means that it is returned a fixedUsername.
	if fixedUsername != "" && fixedUsername != NOTEXISTS {

		// USER PROFILE LOGIN:
		log.Println("The User already exists!")
		var saved_uuid string

		// Uuid Retrieval from the DB.
		err := db.c.QueryRow(`SELECT uuid FROM Users WHERE username == ?`, username).Scan(&saved_uuid)
		if !errors.Is(err, nil) {

			// Uuid failed to be retrieved from the DB.
			log.Println("Err: Failed to Retrieve UUID from the DB")
			return "", err, ""
		} else {

			// Uuid retrieved correctly from the DB.
			log.Println("Uuid Retrieval Succeeded from the DB!")
			return saved_uuid, nil, PRESENT
		}
	}

	// Check then whether it is an error "ErrUserDoesNotExist".
	if errors.Is(errUserPresence, ErrUserDoesNotExist) {

		// USER PROFILE CREATION:
		// The User deos not Exists.
		log.Println("The User does not exists!")

		// Uuid Generation.
		var uuid = uuid.Must(uuid.NewV4())
		log.Println("Generated UUID", uuid)

		// Getting the last fixedUsername + 1
		fixedUsername, errfixedUsername := db.GetLastFixedUsername()
		if !errors.Is(errfixedUsername, nil) {

			// Last fixedUsername failed to be retrieved.
			log.Println("Err: Last fixedUsername failed to be retrieved")
			return "", errfixedUsername, ""
		}

		// Actual User insertion in the DB. Insertion of the actual uuid, username and (after), update the fixedUsername.
		// The rest of the User is completely Standard, in such a way to have that the user is not obliged to add nothing else.
		_, errCreation := db.c.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			fixedUsername, uuid.String(), username, "./service/api/photos/Profile-Already-Created/No-Picture-Available.png", "", now, 0, 0, 0, "", "", "1900-01-01", "surname.matriculation@studenti.uniroma1.it", "", "do not specify")

		// Check whether we have experienced an error from the User Insertion.
		if !errors.Is(errCreation, nil) {
			log.Println("Err: Error During Creation")
			return "Error", errCreation, ""
		}

		// The User has been Created successfully.
		log.Println("User Creation Succeeded!")

		// If we arrive here, we have successfully created the User.
		return uuid.String(), nil, NOTPRESENT
	}

	// Fist, check whether there is an error strange, i.e., that is neither nil nor ErrUserDoesNotExists.
	log.Println("Err: Unexpected Error during the Query of the DB!")
	return "", errUserPresence, ""
}
