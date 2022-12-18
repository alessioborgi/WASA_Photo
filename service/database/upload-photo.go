package database

import (
	"errors"
	"log"
)

func (db *appdbimpl) UploadPhoto(username string, photo Photo, uuid string) error {

	// Selection of the User profile.
	// Here, you have 4 options, stored in the "authorization" variable:
	// 1) AUTHORIZED: The action requester is the Profile Owner. It can Add the photo.
	// 2) UNAUTHORIZED: The action requester is NOT the Profile Owner. It cannot Add the photo.
	// 3) NOT VALID: The action requester has not inserted a valid Uuid, since it's not present in the DB.
	// 4) "": Returned if we have some errors.

	// Let's now check whether the Username we want to insert is already present in the WASAPhoto Platform.
	_, errUsername := db.CheckUserPresence(username)

	// Check whether the Username I am trying to update with the newUsername, does not exists.
	if errors.Is(errUsername, ErrUserDoesNotExist) {
		log.Println("Err: The Username is trying to update the photo does not exists. Error!")
		return ErrBadRequest
	}

	// Check if strange errors occurs.
	if !errors.Is(errUsername, nil) && !errors.Is(errUsername, Ok) {
		log.Println("Err: Strange error during the Check of User Presence")
		return errUsername
	}

	// If both the Usernames are ok, check the Authorization of the person who is asking the action.
	authorization, errAuth := db.CheckAuthorizationOwnerUsername(username, uuid)

	// Check for the error during the Query.
	if !errors.Is(errAuth, nil) {

		// Check whether we have received some errors during the Authentication.
		return errAuth
	}

	// We can now go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
	if authorization == AUTHORIZED {

		// If the uuid is requesting the action is the actual User Owner.
		log.Println("The User is Authorized to add a Photo")

		// We can execute the Insertion.
		_, errInsertion := db.c.Exec(`INSERT INTO Photos (photoid, fixedUsername, filename, uploadDate, phrase, numberLikes, numberComments) VALUES (?, ?, ?, ?, ?, ?, ?)`,
			photo.Photoid, photo.FixedUsername, photo.Filename, photo.UploadDate, photo.Phrase, photo.NumberLikes, photo.NumberComments)

		// Check if some strage error occurred during the update.
		if !errors.Is(errInsertion, nil) {
			log.Println("Err: Error during Photo Insertion.")
			return errInsertion
		}

		// 4 - return mediaID
		return nil
	}

	// We can now see what to do if the Uuid that is requesting the action is not the User Owner.
	if authorization == NOTAUTHORIZED {

		// If the User was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
		log.Println("Err: The Uuid you are providing is not Authorized to do this action.")
		return ErrUserNotAuthorized
	}

	// Check if we have a NOTVALID auth, i.e., the Uuid is not present in the DB.
	if authorization == NOTVALID {

		log.Println("Err: The Uuid you are providing is not present.")
		return ErrUserNotAuthorized
	}

	// If we arrive here, we encountered other types of problem.
	log.Println("Err: Unexpected Error.")
	return errAuth
}
