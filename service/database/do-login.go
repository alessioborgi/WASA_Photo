package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

//DOUBT: What does lastInserId does and what from line 13 happens?

func (db *appdbimpl) DoLogin(username string) (string, error) {

	//First check whether there exists a User with the inserted Username.
	var exists = 0
	// fmt.Println("SELECT 1 FROM Users WHERE username == '?'", username)
	// fmt.Println(exists)
	err := db.c.QueryRow(`SELECT 1 FROM Users WHERE username == '?'`, username).Scan(&exists)
	// fmt.Println(exists)
	// fmt.Println(err)
	//Check for the error during the Query.
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Unexpected Error!")
		return "", err
	} else if exists == 1 {
		//If no strange error during the Query occurs, and exists = 1, we already have the user registered.
		// The User already Exists. LOGIN Part.
		fmt.Println("The User already esists!")
		var uuid string
		errLogin := db.c.QueryRow(`SELECT uuid
		FROM Users
		WHERE username == '?'`, username).Scan(&uuid)
		fmt.Println(exists)
		if errLogin != nil {
			return "", errLogin
		} else {
			return uuid, nil
		}
	} else {
		// The User deos not Exists. User CREATION Part.
		fmt.Println("The User does not esists!")
		var uuid = uuid.Must(uuid.NewV4())
		if err != nil && err != sql.ErrNoRows {
			log.Fatalf("failed to generate UUID: %v", err)
		} else {
			log.Printf("generated Version 4 UUID %v", uuid)
		}
		res, errCretion := db.c.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, totNumberLikes, totNumberComments, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			uuid.String(), uuid.String(), username, "0000000000000000000000000000000000000000000000000000000000000000000000", "", "2022-01-01", 0, 0, 0, 0, 0, "", "", "1900-01-01", "surname.matriculation@studenti.uniroma1.it", "", "do not specify")
		if errCretion != nil {
			fmt.Println("Error Insert")
			return "Error", errCretion
		} else {
			fmt.Println("Good Insert")
			lastInsertID, err := res.LastInsertId()
			if err != nil {
				return "", err
			} else {
				var a = lastInsertID

				return string(rune(a)), nil
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
