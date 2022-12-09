package database

import (
	"database/sql"
	"log"

	"github.com/gofrs/uuid"
)

// DOUBT: What does lastInserId does and what from line 13 happens?

func (db *appdbimpl) DoLogin(username string) (string, error) {

	// First check whether there exists a User with the inserted Username.
	var exists = 0
	err := db.c.QueryRow(`SELECT COUNT(fixedUsername) FROM Users WHERE username == ?`, username).Scan(&exists)

	// Check for the error during the Query.
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("Unexpected Error!")
		return "", err
	} else if exists == 1 {
		// If no strange error during the Query occurs, and exists = 1, we already have the user registered.
		// The User already Exists. LOGIN Part.
		log.Println("The User already esists!")
		var saved_uuid string
		err := db.c.QueryRow(`SELECT uuid FROM Users WHERE username == ?`, username).Scan(&saved_uuid)
		if err != nil {
			log.Fatalf("Failed to Retrieve UUID")
			return "", err
		} else {
			log.Println("Uuid Retrieval Succeeded!")
			return saved_uuid, nil
		}

	} else {
		// The User deos not Exists. User CREATION Part.
		log.Println("The User does not exists!")
		var uuid = uuid.Must(uuid.NewV4())
		if err != nil && err != sql.ErrNoRows {
			log.Fatalf("Failed to generate UUID: %v", err)
		} else {
			log.Println("Generated UUID", uuid)
		}

		// Add the snapshot time of when it is added automatically with like "time.now".
		res, errCretion := db.c.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, totNumberLikes, totNumberComments, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			uuid.String(), uuid.String(), username, "0000000000000000000000000000000000000000000000000000000000000000000000", "", "2022-01-01", 0, 0, 0, 0, 0, "", "", "1900-01-01", "surname.matriculation@studenti.uniroma1.it", "", "do not specify")
		if errCretion != nil {
			log.Fatalf("Error During Creation")
			return "Error", errCretion
		} else {
			log.Println("User Creation Succeeded!")
			lastInsertID, err := res.LastInsertId()
			if err != nil {
				log.Fatalf("User fixedUsername retrieval Error")
				return "", err
			} else {
				log.Println("User fixedUsername retrieval Succedeed.")
				var a = lastInsertID
				_, errUpdate := db.c.Exec(`UPDATE Users SET fixedUsername="?" WHERE username = '?'`, string(rune(a)), username)
				if errUpdate != nil {
					log.Fatalf("Error During Updatating")
					return "", errUpdate
				} else {
					log.Println("fixedUsername Update succeeded")
					return uuid.String(), nil
				}
			}
		}
	}
}

// ------------------------------------------------------------------------------------
// import (
// 	"log"

// 	"github.com/gofrs/uuid"
// )

// // Create a Version 4 UUID, panicking on error.
// // Use this form to initialize package-level variables.
// var u1 = uuid.Must(uuid.NewV4())

// func main() {
// 	// Create a Version 4 UUID.
// 	u2, err := uuid.NewV4()
// 	if err != nil {
// 		log.Fatalf("failed to generate UUID: %v", err)
// 	}
// 	log.Printf("generated Version 4 UUID %v", u2)

// 	// Parse a UUID from a string.
// 	s := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
// 	u3, err := uuid.FromString(s)
// 	if err != nil {
// 		log.Fatalf("failed to parse UUID %q: %v", s, err)
// 	}
// 	log.Printf("successfully parsed UUID %v", u3)
// }
