package database

import (
	"errors"
	"log"
)

func (db *appdbimpl) GetPhotos(username string, uuid string) ([]Photo, error) {

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
		return nil, ErrUserDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errUsername, nil) {
		log.Println("Err: Strange error during the Check of User Presence")
		return nil, errUsername
	}

	// If both the Usernames are ok, check the Authorization of the person who is asking the action.
	authorization, errAuth := db.CheckAuthorizationOwnerUsername(username, uuid)

	// Check for the error during the Query.
	if !errors.Is(errAuth, nil) {

		// Check whether we have received some errors during the Authentication.
		return nil, errAuth
	}

	// We can now go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
	if authorization == AUTHORIZED || authorization == NOTAUTHORIZED {

		// If the User was not "Authorized", i.e. it is not the Profile Owner, it must be checked whether you are banned or not.
		// 0.1) First of all, I need to check whether the username that wants to know the Followings of exists.
		fixedUsernameRequester, errfixedUsernameRequester := db.GetFixedUsername(uuid)

		// Check if strange errors occurs.
		if !errors.Is(errfixedUsernameRequester, nil) {
			log.Println("Err: Unexpected Error in the Username Requester Retrieval ")
			return nil, errfixedUsernameRequester
		}

		// If we arrive here, we have correclty retrieved the Requester Username.
		// Proceed to check whether it is Banned or not.
		errBanRetrieval := db.CheckBanPresence(fixedUsername, fixedUsernameRequester)

		if errors.Is(errBanRetrieval, Okay_Error_Inverse) {
			log.Println("Err: The Ban exists. You cannot get Photos it.")
			return nil, ErrUserNotAuthorized
		}

		// Check if strange errors occurs.
		if !errors.Is(errBanRetrieval, nil) && !errors.Is(errBanRetrieval, ErrBanDoesNotExist) {
			log.Println("Err: Strange error during the Check of Ban Presence")
			return nil, errBanRetrieval
		}

		// If we arrive here, the User is either the profile owner or it is not Banned, thus can get the list of photos.
		// If the uuid is requesting the action is the actual User Owner, get the list of Photos.
		photos, err := db.c.Query(`SELECT *
			FROM Photos
			WHERE fixedUsername = ?
			ORDER BY uploadDate DESC`, fixedUsername)

		// Check for the error during the Query.
		if !errors.Is(err, nil) {
			return nil, err
		}

		// If I arrive here, I got some Photos.
		// Defer the Photos closure. This is a Best-Practice.
		defer func() { _ = photos.Close() }()

		// Here we read the resultset and we build the list to be returned.

		// Variable Declaration.
		var photoList []Photo

		for photos.Next() {
			var p Photo
			err = photos.Scan(&p.Photoid, &p.FixedUsername, &p.Filename, &p.UploadDate, &p.Phrase, &p.NumberLikes, &p.NumberComments)
			if !errors.Is(err, nil) {
				return nil, err
			}
			// Append to the photoList if no error occurs.
			photoList = append(photoList, p)
		}

		// If we have encountered some error in the photos variable.
		if photos.Err() != nil {
			log.Println("Err: Error encountered on followings")
			return nil, err
		}

		// Check whether the list of Photos is empty.
		if len(photoList) == 0 {
			return nil, ErrNoContent
		}

		// If we arrive here, the photoList is not empty, thus, we can return it.
		return photoList, nil
	}

	// Check if we have a NOTVALID auth, i.e., the Uuid is not present in the DB.
	if authorization == NOTVALID {

		log.Println("Err: The Uuid you are providing is not present.")
		return nil, ErrUserNotAuthorized
	}

	// If we arrive here, we encountered other types of problem.
	log.Println("Err: Unexpected Error.")
	return nil, errAuth

}
