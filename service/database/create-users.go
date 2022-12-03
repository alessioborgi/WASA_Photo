package database

//DOUBT: What does lastInserId does and what from line 13 happens?

func (db *appdbimpl) DoLogin(u User) (User, error) {
	// res, err := db.c.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, totNumberLikes, totNumberComments, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
	_, err := db.c.Exec(`INSERT INTO Users (fixedUsername, uuid, username, photoProfile, biography, dateOfCreation, numberOfPhotos, totNumberLikes, totNumberComments, numberFollowers, numberFollowing, name, surname, dateOfBirth, email, nationality, gender) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		u.FixedUsername, u.Uuid, u.Username, u.PhotoProfile, u.Biography, u.DateOfCreation, u.NumberOfPhotos, u.TotNumberLikes, u.TotNumberComments, u.NumberFollowers, u.NumberFollowing, u.Name, u.Surname, u.DateOfBirth, u.Email, u.Nationality, u.Gender)
	if err != nil {
		return u, err
	}
	//I added this line. Remove if below is necessary.
	return u, err

	// lastInsertID, err := res.LastInsertId()
	// if err != nil {
	// 	return u, err
	// }

	// f.ID = uint64(lastInsertID)
	// return f, nil
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
