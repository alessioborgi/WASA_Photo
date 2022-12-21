package database

import (
	"errors"
	"log"
)

func (db *appdbimpl) DeleteUser(username string, uuid string) error {
	// Deletion of the User profile. Here we can distinguish two cases:
	// Here, you have 4 options, stored in the "authorization" variable:
	// 1) AUTHORIZED: The action requester is the Profile Owner.
	// 2) UNAUTHORIZED: The action requester is NOT the Profile Owner.
	// 3) NOT VALID: The action requester has not inserted a valid Uuid, since it's not present in the DB.
	// 4) "": Returned if we have some errors.

	// First of all, check the Authorization of the person who is asking the action.
	authorization, errAuth := db.CheckAuthorizationOwnerUsername(username, uuid)

	// Check for the error during the Query.
	if !errors.Is(errAuth, nil) {
		return errAuth
	}

	// We can now go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
	if authorization == AUTHORIZED {

		// If the uuid is requesting the action is the actual User Owner.

		// First of all, I need to check whether the fixedUsername on which uuid wants to do the action exists.
		_, errUsername := db.CheckUserPresence(username)

		// Check whether the fixedUsername I am trying to delete, does not exists.
		if errors.Is(errUsername, ErrUserDoesNotExist) {
			log.Println("Err: The fixedUsername I am trying to delete, does not exists.")
			return ErrUserDoesNotExist
		}

		// Check if strange errors occurs.
		if !errors.Is(errUsername, nil) && !errors.Is(errUsername, Ok) {
			log.Println("Err: Strange error during the Check of User Presence")
			return errUsername
		}

		// Here I arrive if the fixedUsername I am trying to delete exists(Ok). I have the fixedUsername passed in input.

		// Perform the actual Deletion of the User profile from the DB.
		_, errDeletion := db.c.Exec(`DELETE FROM Users WHERE username=?`, username)
		if !errors.Is(errDeletion, nil) {
			log.Println("Error encountered during the User Deletion in the DB.")
			return errDeletion
		}

		return Ok
	}

	// We can now see what to do if the Uuid that is requesting the action is not the User Owner.
	if authorization == NOTAUTHORIZED {

		//If the User was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
		log.Println("Err: The Uuid you are providing is not Authorized to do this action.")
		return ErrUserNotAuthorized
	}

	// Check if we have a NOTVALID auth, i.e., the Uuuid is not present in the DB.
	if authorization == NOTVALID {

		log.Println("Err: The Uuid you are providing is not present.")
		return ErrUserNotAuthorized
	}

	// If we arrive here, we encountered other types of problem.
	log.Println("Err: Unexpected Error.")
	return errAuth
}
