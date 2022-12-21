package database

import (
	"errors"
	"log"
)

func (db *appdbimpl) GetMyStream(username string, uuid string) ([]Photo, error) {

	// Selection of the User profile.
	// Here, you have 4 options, stored in the "authorization" variable:
	// 1) AUTHORIZED: The action requester is the Profile Owner. It can get the photos.
	// 2) UNAUTHORIZED: The action requester is NOT the Profile Owner. It cannot get the photos.
	// 3) NOT VALID: The action requester has not inserted a valid Uuid, since it's not present in the DB.
	// 4) "": Returned if we have some errors.

	// Let's now check whether the Username we want to insert is already present in the WASAPhoto Platform.
	fixedUsername, errUsername := db.CheckUserPresence(username)

	// Check whether the Username exists.
	if errors.Is(errUsername, ErrUserDoesNotExist) {
		log.Println("Err: The Username that is trying to get its StreamPhotos does not exists. Error!")
		return nil, ErrUserDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errUsername, nil) && !errors.Is(errUsername, Ok) {
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
	if authorization == AUTHORIZED {

		// If the User is the profile owner, it can get the list of Photos of its Stream.
		photos, err := db.c.Query(`SELECT p.photoid, p.fixedUsername, p.filename, p.uploadDate, p.phrase, p.numberLikes, p.numberComments
			FROM Photos as p JOIN (SELECT fixedUsernameFollowing FROM Follows WHERE fixedUsername == ?) as f ON p.fixedUsername = f.fixedUsernameFollowing
			ORDER BY p.uploadDate DESC`, fixedUsername)

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

	// Check if we have a NOTAUTHORIZED auth, i.e., the Uuid is not present in the DB.
	if authorization == NOTAUTHORIZED {

		// We simply cannot view the Stream of others.
		log.Println("Err: You are not authorized to view Stream of others.")
		return nil, ErrUserNotAuthorized
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
