package database

import (
	"errors"
	"log"
	"time"
)

func (db *appdbimpl) CommentPhoto(username string, photoid string, comment Comment, uuid string) error {

	// Adding a User Follow.
	// Here, you have 4 options, stored in the "authorization" variable:
	// 1) AUTHORIZED: The action requester is the usernameCommenter. It can proceed to add the comment, if it is not banned.
	// 2) UNAUTHORIZED: The action requester is the Profile Owner. It can proceed to add the comments.
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
	if !errors.Is(errUsername, nil) && !errors.Is(errUsername, Ok) {
		log.Println("Err: Strange error during the Check of User Presence")
		return errUsername
	}

	// Check whether the photo is present in the DB.
	_, errPhoto := db.CheckPhotoPresence(photoid, fixedUsername)

	// Check whether the Username exists.
	if errors.Is(errPhoto, ErrPhotoDoesNotExist) {
		log.Println("Err: The Photo that fixedUsernameCommenter is trying to comment does not exists. Error!")
		return ErrPhotoDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errPhoto, nil) && !errors.Is(errPhoto, Ok) {
		log.Println("Err: Strange error during the Check of Photo Presence")
		return errPhoto
	}

	// Getting the FixedUsername of the Commenter
	fixedUsernameCommenter, errfixedUsernameCommenter := db.GetFixedUsername(uuid)

	// Check if strange errors occurs.
	if !errors.Is(errfixedUsernameCommenter, nil) {
		log.Println("Err: Unexpected Error in the Username Commenter Retrieval ")
		return errfixedUsernameCommenter
	}

	// 0.3) We need now to check whether fixedUsernameCommenter is Banned by fixedUsername.
	errBanRetrieval := db.CheckBanPresence(fixedUsername, fixedUsernameCommenter)

	if errors.Is(errBanRetrieval, Ok) {
		log.Println("Err: The Ban exists. You cannot Comment the photo!")
		return ErrUserNotAuthorized
	}

	// Check if strange errors occurs.
	if !errors.Is(errBanRetrieval, nil) && !errors.Is(errBanRetrieval, ErrBanDoesNotExist) {
		log.Println("Err: Strange error during the Check of Ban Presence")
		return errBanRetrieval
	}

	// If we arrive here, it means that the Ban is not present. Thus we can continue.

	// Now, we can finally check the Authorization of the person who is asking the action.
	authorization, errAuth := db.CheckAuthorizationOwnerUsername(username, uuid)

	// Check for the error during the Query.
	if !errors.Is(errAuth, nil) {

		// Check whether we have received some errors during the Authentication.
		return errAuth
	}

	// We can now go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
	if authorization == AUTHORIZED || authorization == NOTAUTHORIZED {

		// If the uuid is requesting the action(username) is the actual User Owner, or
		// If NotAuthorized, (thus the fixedUsernameCommenter is diverse from the Liker), you can proceed to add up the Comment
		// since we have already the certainty that no ban is present.

		// We first need to get the commentId, that corresponds to the last commentId+1.
		commentid, errComment := db.GetLastCommentId()
		if !errors.Is(errComment, nil) {
			log.Println("Err: Error encoutered in the GetLastCommentId")
			return errComment
		}

		// We can now finally insert the Comment in the DB.
		_, err := db.c.Exec(`INSERT INTO Comments (commentid, photoid, fixedUsername, commenterFixedUsername, phrase, uploadDate) VALUES (?, ?, ?, ?, ?, ?)`,
			commentid, photoid, fixedUsername, fixedUsernameCommenter, comment.Phrase, time.Now().Format(time.RFC3339))

		if !errors.Is(err, nil) {
			return err
		}

		log.Println("Photo Comment added correclty.")
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

func GetLastCommentId() {
	panic("unimplemented")
}
