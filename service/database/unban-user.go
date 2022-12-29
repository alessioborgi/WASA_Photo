package database

import (
	"errors"
	"log"
)

func (db *appdbimpl) UnbanUser(username string, usernameBanned string, uuid string) error {
	// Deletion of a User Ban.
	// Here, you have 4 options, stored in the "authorization" variable:
	// 1) AUTHORIZED: The action requester is the Profile Owner. It can proceed to delete the Ban only whether username and usernameBanned is not the same (we cannot eliminate a self ban).
	// 2) UNAUTHORIZED: The action requester is NOT the Profile Owner. It cannot remove the Ban.
	// 3) NOT VALID: The action requester has not inserted a valid Uuid, since it's not present in the DB.
	// 4) "": Returned if we have some errors.

	// 0.0) As a premature check, check whether the username that is requesting the action is going to sel-ban.
	if username == usernameBanned {
		return ErrBadRequest
	}

	// 0.1) First of all, I need to check whether the username that has the Ban exists (that must be also the uuid itself, check later).
	fixedUsernameBanner, errUsername := db.CheckUserPresence(username)

	// Check whether theUsername I am trying to update, does not exists.
	if errors.Is(errUsername, ErrUserDoesNotExist) {
		log.Println("Err: The fixedUsernameBanner, does not exists.")
		return ErrUserDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errUsername, nil) {
		log.Println("Err: Strange error during the Check of User Presence")
		return errUsername
	}

	// 0.2) Secondly, I need to check only whether the usernameBanned I passed exists in the DB.
	fixedUsernameBanned, errusernameBanned := db.CheckUserPresence(usernameBanned)

	// Check whether the usernameBanned I am trying to insert, already exists.
	if errors.Is(errusernameBanned, ErrUserDoesNotExist) {
		log.Println("Err: The fixedUsernameBanned, does not exists.")
		return ErrUserDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errusernameBanned, nil) {
		log.Println("Err: Strange error during the Check of User Presence")
		return errusernameBanned
	}

	// 0.3) Thirdly, we should check whether there exists an actual Ban to delete between fixedusernameBanner and fixedusernameBanned.
	_, errBanRetrieval := db.CheckBanPresence(fixedUsernameBanner, fixedUsernameBanned)
	if errors.Is(errBanRetrieval, ErrBanDoesNotExist) {
		log.Println("Err: The Ban, does not exists.")
		return ErrBanDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errBanRetrieval, nil) {
		log.Println("Err: Strange error during the Check of Ban Presence")
		return errBanRetrieval
	}

	// If we arrive here, it means that the Ban is present. Thus we can continue.
	// First of all, check the Authorization of the person who is asking the action.
	authorization, errAuth := db.CheckAuthorizationOwnerUsername(username, uuid)

	// Check for the error during the Query.
	if !errors.Is(errAuth, nil) {

		// Check whether we have received some errors during the Authentication.
		return errAuth
	}

	// We can now go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
	if authorization == AUTHORIZED {

		// If the uuid is requesting the action is the actual User Owner.
		// If Authorized, you can proceed to remove the Ban without any problem.
		_, err := db.c.Exec(`DELETE FROM Bans WHERE fixedUsernameBanner=? AND fixedUsernameBanned = ?`, fixedUsernameBanner, fixedUsernameBanned)
		if !errors.Is(err, nil) {
			return err
		}

		// The Insertion went well.
		return nil
	}

	// We can now see what to do if the Uuid that is requesting the action is not the User Owner.
	if authorization == NOTAUTHORIZED {

		// If the Use was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
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
