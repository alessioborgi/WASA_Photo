package database

import (
	"errors"
	"log"
)

func (db *appdbimpl) UncommentPhoto(username string, photoid string, commentid string, uuid string) error {

	// Adding a User Follow.
	// Here, you have 4 options, stored in the "authorization" variable:
	// 1) AUTHORIZED: The action requester is the fixedUsernameCommenter. It can proceed to delete the comment, if it is not banned and if it is the Commenter maker.
	// 2) UNAUTHORIZED: The action requester is NOT the fixedUsernameCommenter. It cannot proceed to delete the comment.
	// 3) NOT VALID: The action requester has not inserted a valid Uuid, since it's not present in the DB.
	// 4) "": Returned if we have some errors.

	// 0.1) First of all, I need to check whether the username that has the photo exists.
	fixedUsername, errUsername := db.CheckUserPresence(username)

	// Check whether the Username that owns the photo, does not exists.
	if errors.Is(errUsername, ErrUserDoesNotExist) {
		log.Println("Err: The Username, does not exists.")
		return ErrUserDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errUsername, nil) {
		log.Println("Err: Strange error during the Check of User Presence")
		return errUsername
	}

	// 0.2) Check whether the photo is present in the DB.
	_, errPhoto := db.CheckPhotoPresence(photoid, fixedUsername)

	// Check whether the Username exists.
	if errors.Is(errPhoto, ErrPhotoDoesNotExist) {
		log.Println("Err: The Photo that fixedUsernameCommenter is trying to delete the comment does not exists. Error!")
		return ErrPhotoDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errPhoto, nil) && !errors.Is(errPhoto, Okay_Error_Inverse) {
		log.Println("Err: Strange error during the Check of Photo Presence.")
		return errPhoto
	}

	// 0.3) Getting the FixedUsername of the Commenter.
	fixedUsernameCommenter, errfixedUsernameCommenter := db.GetFixedUsername(uuid)

	// Check whether the Username that commented the photo, does not exists.
	if errors.Is(errUsername, ErrUserDoesNotExist) {
		log.Println("Err: The UsernameCommenter, does not exists.")
		return ErrUserDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errfixedUsernameCommenter, nil) {
		log.Println("Err: Unexpected Error in the Username Commenter Retrieval ")
		return errfixedUsernameCommenter
	}

	// 0.4) Thirdly, we should check whether there exists the comment.
	errCommentRetrieval := db.CheckCommentPresence(commentid, fixedUsername, photoid, fixedUsernameCommenter)
	if errors.Is(errCommentRetrieval, ErrCommentDoesNotExist) {
		log.Println("Err: The Comment didn't exists!")
		return ErrCommentDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errCommentRetrieval, nil) && !errors.Is(errCommentRetrieval, Okay_Error_Inverse) {
		log.Println("Err: Strange error during the Check of Comment Presence")
		return errCommentRetrieval
	}

	// If we arrive here, it means that the Comment is present. Thus we can continue.
	// 0.5) We need now to check whether fixedUsernameLiker is Banned by fixedUsername.
	errBanRetrieval := db.CheckBanPresence(fixedUsername, fixedUsernameCommenter)

	if errors.Is(errBanRetrieval, Okay_Error_Inverse) {
		log.Println("Err: The Ban exists. You cannot delete the comment to the photo!")
		return ErrUserNotAuthorized
	}

	// Check if strange errors occurs.
	if !errors.Is(errBanRetrieval, nil) && !errors.Is(errBanRetrieval, ErrBanDoesNotExist) {
		log.Println("Err: Strange error during the Check of Ban Presence")
		return errBanRetrieval
	}

	// If we arrive here, it means that the Ban is not present. Thus we can continue.

	// Now, we can finally check the Authorization of the person who is asking the action.
	authorization, errAuth := db.CheckAuthorizationOwner(fixedUsernameCommenter, uuid)

	// Check for the error during the Query.
	if !errors.Is(errAuth, nil) {

		// Check whether we have received some errors during the Authentication.
		return errAuth
	}

	// We can now go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
	if authorization == AUTHORIZED {

		// If the uuid is requesting the action(usernameCommenter) is the actual Commenter maker.
		// Actual DB Photo Deletion.
		_, errDeletion := db.c.Exec(`DELETE FROM COMMENTS WHERE commentid = ? AND photoid = ? AND fixedUsername = ? AND commenterFixedUsername = ?`, commentid, photoid, fixedUsername, fixedUsernameCommenter)

		// Check if some strage error occurred during the Deletion.
		if !errors.Is(errDeletion, nil) {
			log.Println("Err: Error during Photo's Comment Deletion.")
			return errDeletion
		}

		// The deletion has been correclty delivered.
		log.Println("Photo's Comment Deletion Successfull.")
		return nil

	}

	if authorization == NOTAUTHORIZED {

		// If NotAuthorized, (thus the usernameLiker is diverse from whoever is asking the action), you cannot proceed to remove the Like.
		log.Println("Err: You should are not able to delete the like of others.")
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
