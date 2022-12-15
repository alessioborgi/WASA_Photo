package database

import (
	"errors"
	"fmt"
	"log"
)

// DOUBT: Do I have to pass it the entire User Object? Do I have to pass to it the new Username?

func (db *appdbimpl) SetMyUsername(fixedUsername string, newUsername string, uuid string) error {

	// Selection of the User profile.
	// Here, you have 4 options, stored in the "authorization" variable:
	// 1) AUTHORIZED: The action requester is the Profile Owner.
	// 2) UNAUTHORIZED: The action requester is NOT the Profile Owner.
	// 3) NOT VALID: The action requester has not inserted a valid Uuid, since it's not present in the DB.
	// 4) "": Returned if we have some errors.

	// First of all, check the Authorization of the person who is asking the action.
	authorization, errAuth := db.CheckAuthorizationOwner(fixedUsername, uuid)

	fmt.Println(authorization, errAuth)
	// Check for the error during the Query.
	if errAuth != nil {

		// Check whether we have received some errors during the Authentication.
		return errAuth
	}

	// We can now go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
	if authorization == AUTHORIZED {

		// If the uuid is requesting the action is the actual User Owner.
		// First of all, I need to check whether the fixedUsername on which uuid wants to do the action exists.
		_, errUsername := db.CheckFixedUserPresence(fixedUsername)
		fmt.Println("errUsername is: ", errUsername)

		// Check if strange errors occurs.
		if !errors.Is(errUsername, nil) && !errors.Is(errUsername, ErrUserDoesNotExist) && !errors.Is(errUsername, Ok) {
			log.Println("Err: Strange error during the Check of User Presence")
			return errUsername
		}

		// Check whether theUsername I am trying to update, does not exists.
		if errors.Is(errUsername, ErrUserDoesNotExist) {
			log.Println("Err: The Username I am trying to update, does not exists.")
			return ErrUserDoesNotExist
		}

		// Here I arrive if the Username I am trying to update exists. I have the fixedUsername passed in input.

		// Perform the Update of the Username.
		res, errUpdate := db.c.Exec(`UPDATE Users SET username=? WHERE fixedUsername=?`, newUsername, fixedUsername)

		// Check if some strage error occurred during the update.
		if !errors.Is(errUpdate, nil) {
			log.Println("Err: Error during Update.")
			return errUpdate
		}

		// Here arrives if no strange errors occurred.
		affected, err := res.RowsAffected()
		if !errors.Is(err, nil) {
			log.Println("Err: Error during the rowsAffected retrieval.")
			return err
		} else if affected == 0 {
			// If we didn't modified any row, then the User didn't exist.
			log.Println("Err: You didn't modified any row, meaning that the User doesn't exists.")
			return ErrUserDoesNotExist
		}
		return nil
	}

	// We can now see what to do if the Uuid that is requesting the action is not the User Owner.
	if authorization == NOTAUTHORIZED {

		//If the Use was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
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
