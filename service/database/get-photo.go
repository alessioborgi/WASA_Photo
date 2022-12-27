package database

import (
	"errors"
	"log"
)

func (db *appdbimpl) GetPhoto(username string, photoid string, uuid string) (Photo, error) {

	// Selection of the User profile.
	// Here, you have 4 options, stored in the "authorization" variable:
	// 1) AUTHORIZED: The action requester is the Profile Owner. It can get the photos.
	// 2) UNAUTHORIZED: The action requester is NOT the Profile Owner. It can get the photos only if not banned.
	// 3) NOT VALID: The action requester has not inserted a valid Uuid, since it's not present in the DB.
	// 4) "": Returned if we have some errors.

	// Let's now check whether the Username we want to insert is already present in the WASAPhoto Platform.
	fixedUsername, errUsername := db.CheckUserPresence(username)

	// Check whether the Username exists.
	if errors.Is(errUsername, ErrUserDoesNotExist) {
		log.Println("Err: The Username I am trying to get the Photos does not exists. Error!")
		return Photo{}, ErrUserDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errUsername, nil) {
		log.Println("Err: Strange error during the Check of User Presence")
		return Photo{}, errUsername
	}

	// If we arrive here, it means that the Username exists. Check whether the photo is present in the DB.
	_, errPhoto := db.CheckPhotoPresence(photoid, fixedUsername)

	// Check whether the Username exists.
	if errors.Is(errPhoto, ErrPhotoDoesNotExist) {
		log.Println("Err: The Photo I am trying to get does not exists. Error!")
		return Photo{}, ErrPhotoDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errPhoto, nil) {
		log.Println("Err: Strange error during the Check of Photo Presence")
		return Photo{}, errPhoto
	}

	// If we arrive here, we have that, errPhoto= nil, and therefore it all ok.

	// If I arrive here, both the Username and the Photo exists.
	// Check the Authorization of the person who is asking the action.
	authorization, errAuth := db.CheckAuthorizationOwnerUsername(username, uuid)

	// Check for the error during the Query.
	if !errors.Is(errAuth, nil) {

		// Check whether we have received some errors during the Authentication.
		return Photo{}, errAuth
	}

	// We can now go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
	if authorization == AUTHORIZED || authorization == NOTAUTHORIZED {

		// If the User was not "Authorized", i.e. it is not the Profile Owner, it must be checked whether you are banned or not.
		// 0.1) First of all, I need to check whether the username that wants to know the Followings of exists.
		fixedUsernameRequester, errfixedUsernameRequester := db.GetFixedUsername(uuid)

		// Check if strange errors occurs.
		if !errors.Is(errfixedUsernameRequester, nil) {
			log.Println("Err: Unexpected Error in the Username Requester Retrieval ")
			return Photo{}, errfixedUsernameRequester
		}

		// If we arrive here, we have correclty retrieved the Requester Username.
		// Proceed to check whether it is Banned or not.
		ban_presence, errBanRetrieval := db.CheckBanPresence(fixedUsername, fixedUsernameRequester)

		if ban_presence == PRESENT {
			log.Println("Err: The Ban exists. You cannot get the Photo.")
			return Photo{}, ErrUserNotAuthorized
		}

		// Check if strange errors occurs.
		if !errors.Is(errBanRetrieval, nil) && !errors.Is(errBanRetrieval, ErrBanDoesNotExist) {
			log.Println("Err: Strange error during the Check of Ban Presence")
			return Photo{}, errBanRetrieval
		}

		// If we arrive here, the User is either the profile owner or it is not Banned, thus can get the photo.
		var photo Photo
		err := db.c.QueryRow(`SELECT *
			FROM Photos
			WHERE fixedUsername = ? AND photoid = ?
			ORDER BY uploadDate DESC`, fixedUsername, photoid).Scan(&photo.Photoid, &photo.FixedUsername, &photo.Filename, &photo.UploadDate, &photo.Phrase, &photo.NumberLikes, &photo.NumberComments)

		// Check for the error during the Query.
		if !errors.Is(err, nil) {

			// If we have encountered some errors in the Query retrieval.
			log.Println("Err: Unexpected Error! During the Query Retrieval!")
			return Photo{}, err
		}

		// Otherwise we have retrieved the Photo Correctly.
		log.Println("Photo retrieved correctly!")
		return photo, nil
	}

	// Check if we have a NOTVALID auth, i.e., the Uuid is not present in the DB.
	if authorization == NOTVALID {

		log.Println("Err: The Uuid you are providing is not present.")
		return Photo{}, ErrUserNotAuthorized
	}

	// If we arrive here, we encountered other types of problem.
	log.Println("Err: Unexpected Error.")
	return Photo{}, errAuth

}
