package database

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
)

var now = time.Now()

func (db *appdbimpl) DoLogin(username string) (string, error) {

	// First check whether there exists a User with the inserted Username.
	_, errUserPresence := db.CheckUserPresence(username)
	// var exists = 0
	// err := db.c.QueryRow(`SELECT COUNT(fixedUsername) FROM Users WHERE username == ?`, username).Scan(&exists)

	// Check whether the error is "Ok". If we have it to be Ok, it means it is already present in the DB.
	if errors.Is(errUserPresence, Ok) {

		// USER PROFILE LOGIN:
		log.Println("The User already exists!")
		var saved_uuid string

		// Uuid Retrieval from the DB.
		err := db.c.QueryRow(`SELECT uuid FROM Users WHERE username == ?`, username).Scan(&saved_uuid)
		if err != nil {

			// Uuid failed to be retrieved from the DB.
			log.Fatalf("Failed to Retrieve UUID from the DB")
			return "", err
		} else {

			// Uuid retrieved correctly from the DB.
			log.Println("Uuid Retrieval Succeeded from the DB!")
			return saved_uuid, Ok
		}
	}

	//Check then whether it is an error "ErrUserDoesNotExist".
	if errors.Is(errUserPresence, ErrUserDoesNotExist) {

		// USER PROFILE CREATION:
		// The User deos not Exists.
		log.Println("The User does not exists!")

		//Uuid Generation.
		var uuid = uuid.Must(uuid.NewV4())
		log.Println("Generated UUID", uuid)

		// Actual User insertion in the DB. Insertion of the actual uuid, username and (after), update the fixedUsername.
		// The rest of the User is completely Standard, in such a way to have that the user is not obliged to add nothing else.
		res, errCretion := db.c.Exec(`INSERT INTO Users (fixedUsername, uuid, username, biography, dateOfCreation, numberOfPhotos, totNumberLikes, totNumberComments, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			"", uuid.String(), username, "", now, 0, 0, 0, 0, 0, "", "", "1900-01-01", "surname.matriculation@studenti.uniroma1.it", "", "do not specify")

		// Check whether we have experienced an error from the User Insertion.
		if errCretion != nil {
			log.Fatalf("Error During Creation")
			return "Error", errCretion
		}

		// The User has been Created successfully.
		log.Println("User Creation Succeeded!")

		// User's fixedUsername Update with the last id fixedUsername.
		lastInsertID, err := res.LastInsertId()

		//Check whether we have some error in the LastId retrieval.
		if err != nil {
			log.Fatalf("User fixedUsername retrieval Error")
			return "", err
		}

		// The lastId was retrieved successfully.
		log.Println("User fixedUsername retrieval Succedeed.")

		// Build up the lastInsertId
		last_id := lastInsertID
		last := "u" + strconv.Itoa(int(last_id))

		// Perform an Update of the fixedUsername on the User just created.
		_, errUpdate := db.c.Exec(`UPDATE Users SET fixedUsername=? WHERE username = ?`, last, username)

		// Check whether we have some errors during the update in the DB.
		if errUpdate != nil {
			log.Fatalf("Error During Updatating")
			return "", errUpdate
		}

		// If we arrive here, we have successfully created the User.
		log.Println("fixedUsername Update Succeeded")
		return uuid.String(), Created
	}

	// Fist, check whether there is an error strange, i.e., that is neither nil nor ErrUserDoesNotExists.
	fmt.Println(errUserPresence)
	log.Fatalf("Unexpected Error during the Query of the DB!")
	return "", errUserPresence
}
