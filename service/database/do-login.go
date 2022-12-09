package database

import (
	"log"

	"github.com/gofrs/uuid"
)

//DOUBT: What does lastInserId does and what from line 13 happens?

func (db *appdbimpl) DoLogin(u User) (string, error) {

	//First check whether there exists a User with the inserted Username.
	var exists bool
	err := db.c.QueryRow(`SELECT true
	FROM Users
	WHERE username == '?'`, u.Username).Scan(&exists)

	//Check for the error during the Query.
	if err != nil {
		return "", err
	} else {
		//If no error during the Query occurs, checking whether we already have the user registered, thus going for the Login, or whether we have to do the User Registration.
		if exists == true {
			// The User already Exists. LOGIN Part.
			var uuid string
			errLogin := db.c.QueryRow(`SELECT uuid
			FROM Users
			WHERE fixedUsername == '?'`, u.FixedUsername).Scan(&uuid)
			if errLogin != nil {
				return "", errLogin
			} else {
				return uuid, nil
			}
		} else {
			// The User deos not Exists. User CREATION Part.
			var uuid = uuid.Must(uuid.NewV4())
			if err != nil {
				log.Fatalf("failed to generate UUID: %v", err)
			} else {
				log.Printf("generated Version 4 UUID %v", uuid)
			}
			// res, err := db.c.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, totNumberLikes, totNumberComments, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			_, errCretion := db.c.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, totNumberLikes, totNumberComments, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				u.FixedUsername, uuid.String(), u.Username, u.PhotoProfile, u.Biography, u.DateOfCreation, u.NumberOfPhotos, u.TotNumberLikes, u.TotNumberComments, u.NumberFollowers, u.NumberFollowing, u.Name, u.Surname, u.DateOfBirth, u.Email, u.Nationality, u.Gender)
			if errCretion != nil {
				return "Error", errCretion
			} else {
				return uuid.String(), nil
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
