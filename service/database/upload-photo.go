package database

import "log"

func (db *appdbimpl) UploadPhoto(username string, photo Photo) (Photo, error) {

	// Adding a User Photo. Here we can distinguish two cases:
	//1) We have that the User can add the photo since the photo insertion action has been requested by the profile owner's coinciding with the action requester.
	//2) We have that the User cannot add the photo since the photo insertion action has NOT been requested by the profile owner's coinciding with the action requester.

	// authorization, errAuth := db.CheckAuthorizationOwner(photo.FixedUsername, uuid)

	//Check for the error during the Query.
	// if errAuth != nil {
	// 	return photo, errAuth
	// } else {
	// Go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).

	// if authorization == "Authorized" {

	//If Authorized, you can proceed to add up the Photo Object without any problem.

	// As first thing I need to retrieve the fixedUsername of the Username.
	var fixedUsername string
	err := db.c.QueryRow(`SELECT fixedUsername FROM Users WHERE username == ?`, username).Scan(&fixedUsername)
	if err != nil {
		log.Fatalf("Failed to Retrieve fixedUsername from the DB")
		return Photo{}, err
	} else {
		// If we arrive here, we correctly retrieved the fixedUsername.
		log.Println("fixedUsername Retrieval Succeeded from the DB!")

		// We can therefore proceed to Insert the photo.
		res, err := db.c.Exec(`INSERT INTO Photos (photoid, fixedUsername, filename, uploadDate, phrase, numberLikes, numberComments, latitude, longitude) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			photo.Photoid, fixedUsername, photo.Filename, photo.UploadDate, photo.Phrase, photo.NumberLikes, photo.NumberComments, photo.Latitude, photo.Longitude)
		if err != nil {
			return Photo{}, err
		} else {
			// User Created Successfully.
			log.Println("User's Photo Creation Succeeded!")

			// User's fixedUsername Update.
			lastInsertID, err := res.LastInsertId()
			if err != nil {
				log.Fatalf("Photo's photoid retrieval Error")
				return Photo{}, err
			} else {
				log.Println("Photo's photoid retrieval Succedeed.")
				var old_photoid = photo.Photoid
				var a = lastInsertID
				_, errUpdate := db.c.Exec(`UPDATE Photos SET photoid=? WHERE fixedUsername = ? AND photoid = ?`, string(rune(a)), username, old_photoid)
				if errUpdate != nil {
					log.Fatalf("Error During Updatating")
					return Photo{}, errUpdate
				} else {
					log.Println("fixedUsername Update Succeeded")
					return photo, nil
				}
			}
		}
	}

	// 	} else {
	// 		//If the User was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
	// 		return photo, ErrUserNotAuthorized
	// 	}
	// }
}
