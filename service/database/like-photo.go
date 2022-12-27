package database

import (
	"errors"
	"log"
)

func (db *appdbimpl) LikePhoto(username string, photoid string, usernameLiker string, uuid string) error {

	// Adding a User Follow.
	// Here, you have 4 options, stored in the "authorization" variable:
	// 1) AUTHORIZED: The action requester is the usernameLiker. It can proceed to add the like, if it is not banned.
	// 2) UNAUTHORIZED: The action requester is the Profile Owner. It cannot proceed to add the like. (No self-like to the photos)
	// 3) NOT VALID: The action requester has not inserted a valid Uuid, since it's not present in the DB.
	// 4) "": Returned if we have some errors.

	// 0.0) As a premature check, check whether the username that is requesting the action is going to self-follow.
	if username == usernameLiker {
		return ErrBadRequest
	}

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

	// 0.2) Secondly, I need to check only whether the usernameLiker I passed exists in the DB.
	fixedUsernameLiker, errusernameLiker := db.CheckUserPresence(usernameLiker)

	// Check whether the usernameFollowing I am trying to insert, already exists.
	if errors.Is(errusernameLiker, ErrUserDoesNotExist) {
		log.Println("Err: The fixedUsernameLiker that is trying to put the Like, does not exists.")
		return ErrUserDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errusernameLiker, nil) {
		log.Println("Err: Strange error during the Check of usernameLiker Presence")
		return errusernameLiker
	}

	// 0.3) Thirdly, we should check whether there exists the same like already.
	errLikeRetrieval := db.CheckLikePresence(fixedUsername, photoid, fixedUsernameLiker)
	if errors.Is(errLikeRetrieval, Okay_Error_Inverse) {
		log.Println("Err: The Like already exists.")
		return Okay_Error_Inverse
	}

	// Check if strange errors occurs.
	if !errors.Is(errLikeRetrieval, nil) && !errors.Is(errLikeRetrieval, ErrLikeDoesNotExists) {
		log.Println("Err: Strange error during the Check of Follow Presence")
		return errLikeRetrieval
	}

	// If we arrive here, it means that the Like is not present. Thus we can continue.
	// 0.4) We need now to check whether fixedUsernameLiker is Banned by fixedUsername.
	errBanRetrieval := db.CheckBanPresence(fixedUsername, fixedUsernameLiker)

	if errors.Is(errBanRetrieval, Okay_Error_Inverse) {
		log.Println("Err: The Ban exists. You cannot Like the photo!")
		return ErrUserNotAuthorized
	}

	// Check if strange errors occurs.
	if !errors.Is(errBanRetrieval, nil) && !errors.Is(errBanRetrieval, ErrBanDoesNotExist) {
		log.Println("Err: Strange error during the Check of Ban Presence")
		return errBanRetrieval
	}

	// If we arrive here, it means that the Ban is not present. Thus we can continue.

	// Check whether the photo is present in the DB.
	_, errPhoto := db.CheckPhotoPresence(photoid, fixedUsername)

	// Check whether the Username exists.
	if errors.Is(errPhoto, ErrPhotoDoesNotExist) {
		log.Println("Err: The Photo that fixedUsernameLiker is trying to like does not exists. Error!")
		return ErrPhotoDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errPhoto, nil) && !errors.Is(errPhoto, Okay_Error_Inverse) {
		log.Println("Err: Strange error during the Check of Photo Presence")
		return errPhoto
	}

	// Now, we can finally check the Authorization of the person who is asking the action.
	authorization, errAuth := db.CheckAuthorizationOwnerUsername(username, uuid)

	// Check for the error during the Query.
	if !errors.Is(errAuth, nil) {

		// Check whether we have received some errors during the Authentication.
		return errAuth
	}

	// We can now go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
	if authorization == AUTHORIZED {

		// If the uuid is requesting the action(username) is the actual User Owner.
		log.Println("Err: You should not be able to like your photo. This does not make any sense.")
		return ErrUserNotAuthorized
	}

	if authorization == NOTAUTHORIZED {

		// If NotAuthorized, (thus the fixedUsernameLiker is diverse from the Liker), you can proceed to add up the Like since we have already the certainty that no ban is present.
		_, err := db.c.Exec(`INSERT INTO Likes (likeid, photoid, fixedUsername) VALUES (?, ?, ?)`,
			fixedUsernameLiker, photoid, fixedUsername)
		if !errors.Is(err, nil) {
			return err
		}
		log.Println("Photo Like added correclty.")
		return nil

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
