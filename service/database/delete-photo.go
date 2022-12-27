package database

import (
	"errors"
	"log"
)

func (db *appdbimpl) DeletePhoto(username string, photoid string, uuid string) (string, error) {

	// Selection of the User profile.
	// Here, you have 4 options, stored in the "authorization" variable:
	// 1) AUTHORIZED: The action requester is the Profile Owner. It can Delete the photo.
	// 2) UNAUTHORIZED: The action requester is NOT the Profile Owner. It cannot Delete the photo.
	// 3) NOT VALID: The action requester has not inserted a valid Uuid, since it's not present in the DB.
	// 4) "": Returned if we have some errors.

	// Let's now check whether the Username we want to insert is already present in the WASAPhoto Platform.
	fixedUsername, errUsername := db.CheckUserPresence(username)

	// Check whether the Username I am trying to update with the newUsername, does not exists.
	if errors.Is(errUsername, ErrUserDoesNotExist) {
		log.Println("Err: The Username is trying to delete the photo does not exists. Error!")
		return "", ErrUserDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errUsername, nil) {
		log.Println("Err: Strange error during the Check of User Presence")
		return "", errUsername
	}

	// If we arrive here, it means that the Username exists. Check whether the photo is present in the DB.
	_, errPhoto := db.CheckPhotoPresence(photoid, fixedUsername)

	// Check whether the Username exists.
	if errors.Is(errPhoto, ErrPhotoDoesNotExist) {
		log.Println("Err: The Photo I am trying to get does not exists. Error!")
		return "", ErrPhotoDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errPhoto, nil) && !errors.Is(errPhoto, Okay_Error_Inverse) {
		log.Println("Err: Strange error during the Check of Photo Presence")
		return "", errPhoto
	}

	// If both the Usernames are ok, check the Authorization of the person who is asking the action.
	authorization, errAuth := db.CheckAuthorizationOwnerUsername(username, uuid)

	// Check for the error during the Query.
	if !errors.Is(errAuth, nil) {

		// Check whether we have received some errors during the Authentication.
		return "", errAuth
	}

	// We can now go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
	if authorization == AUTHORIZED {

		// If the uuid is requesting the action is the actual User Owner.
		log.Println("The User is Authorized to Delete a Photo")

		// Retrieving the filename path.
		var filename string
		errFilename := db.c.QueryRow(`SELECT filename FROM Photos WHERE fixedUsername = ? AND photoid = ?`, fixedUsername, photoid).Scan(&filename)

		// Check if some strage error occurred during the Selection.
		if !errors.Is(errFilename, nil) {
			log.Println("Err: Error during Photo Deletion.")
			return "", errFilename
		}

		// Actual DB Photo Deletion.
		_, errDeletion := db.c.Exec(`DELETE FROM Photos WHERE fixedUsername = ? AND photoid = ?`, fixedUsername, photoid)

		// Check if some strage error occurred during the Deletion.
		if !errors.Is(errDeletion, nil) {
			log.Println("Err: Error during Photo Deletion.")
			return "", errDeletion
		}

		// The deletion has been correclty delivered.
		log.Println("Photo Deletion Successfull.")
		return filename, nil

	}

	// We can now see what to do if the Uuid that is requesting the action is not the User Owner.
	if authorization == NOTAUTHORIZED {

		// If the User was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
		log.Println("Err: The Uuid you are providing is not Authorized to do this action.")
		return "", ErrUserNotAuthorized
	}

	// Check if we have a NOTVALID auth, i.e., the Uuid is not present in the DB.
	if authorization == NOTVALID {

		log.Println("Err: The Uuid you are providing is not present.")
		return "", ErrUserNotAuthorized
	}

	// If we arrive here, we encountered other types of problem.
	log.Println("Err: Unexpected Error.")
	return "", errAuth
}
