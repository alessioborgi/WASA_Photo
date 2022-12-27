package database

import (
	"errors"
	"log"
)

func (db *appdbimpl) GetPhotoLikes(username string, photoid string, uuid string) ([]string, error) {

	// Retrieving the Users Ban List.
	// Here, you have 4 options, stored in the "authorization" variable:
	// 1) AUTHORIZED: The action requester is the Profile Owner. It can proceed to retrieve its list of Photoid's Likes usernames.
	// 2) UNAUTHORIZED: The action requester is NOT the Profile Owner. It can take the list, provided that it has not been banned by username.
	// 3) NOT VALID: The action requester has not inserted a valid Uuid, since it's not present in the DB.
	// 4) "": Returned if we have some errors.

	// Variable for returning the slice of Likes.
	var LikesList []string

	// 0.1) First of all, I need to check whether the username that owns the photo exists.
	fixedUsername, errfixedUsername := db.CheckUserPresence(username)

	// Check whether the Username I am trying to update, does not exists.
	if errors.Is(errfixedUsername, ErrUserDoesNotExist) {
		log.Println("Err: The Username I am trying to get its photoid's likes List, does not exists.")
		return nil, ErrUserDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errfixedUsername, nil) {
		log.Println("Err: Strange error during the Check of User Presence")
		return nil, errfixedUsername
	}

	// If we arrive here, it means that the Username exists. Check whether the photo is present in the DB.
	_, errPhoto := db.CheckPhotoPresence(photoid, fixedUsername)

	// Check whether the Username exists.
	if errors.Is(errPhoto, ErrPhotoDoesNotExist) {
		log.Println("Err: The Photo I am trying to get its Likes does not exists. Error!")
		return nil, ErrPhotoDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errPhoto, nil) {
		log.Println("Err: Strange error during the Check of Photo Presence")
		return nil, errPhoto
	}

	// If we arrive here, we have that, errPhoto= nil, and therefore it all ok.

	// If both the Usernames are ok, check the Authorization of the person who is asking the action.
	authorization, errAuth := db.CheckAuthorizationOwner(fixedUsername, uuid)

	// Check for the error during the Query.
	if !errors.Is(errAuth, nil) {

		// Check whether we have received some errors during the Authentication.
		return nil, errAuth
	}

	// We can now go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
	if authorization == AUTHORIZED || authorization == NOTAUTHORIZED {

		// If the uuid is requesting the action is the actual User Owner.

		// If the User was not "Authorized", i.e. it is not the Profile Owner, it must be checked whether you are banned or not.
		// 0.1) First of all, I need to get back the fixedUsername of the requester.
		fixedUsernameRequester, errfixedUsernameRequester := db.GetFixedUsername(uuid)

		// Check if strange errors occurs.
		if !errors.Is(errfixedUsernameRequester, nil) {
			log.Println("Err: Unexpected Error in the Username Requester Retrieval ")
			return nil, errfixedUsernameRequester
		}

		// If we arrive here, we have correclty retrieved the Requester Username.
		// Proceed to check whether it is Banned or not.
		ban_presence, errBanRetrieval := db.CheckBanPresence(fixedUsername, fixedUsernameRequester)

		if ban_presence == PRESENT {
			log.Println("Err: The Ban exists. You cannot get the Photo's Likes List.")
			return nil, ErrUserNotAuthorized
		}

		// Check if strange errors occurs.
		if !errors.Is(errBanRetrieval, nil) && !errors.Is(errBanRetrieval, ErrBanDoesNotExist) {
			log.Println("Err: Strange error during the Check of Ban Presence")
			return nil, errBanRetrieval
		}

		// If we arrive here, the user is Authorized.
		// We now need to retrieve the Usernames from the fixedUsername Banned List.
		likes, err := db.c.Query(`SELECT likeid
 				FROM Likes
 				WHERE fixedUsername = ? AND photoid = ?`, fixedUsername, photoid)

		// Here we read the resultset and we build the list to be returned.
		// Check if we have encountered some error in the retrieval of the query.
		if !errors.Is(err, nil) {

			// We check first whether the users retrieval caused an error.
			log.Println("Err: Error encountered during the Query in the DB.")
			return nil, err
		}

		// If no error occur, we can proceed on elaborating the DB Response.
		log.Println("No error encountered during the Query in the DB.")
		defer func() { _ = likes.Close() }()

		// Here we read the resultset and we build the list of Usernames to be returned.
		var u string
		for likes.Next() {
			err = likes.Scan(&u)
			if !errors.Is(err, nil) {
				log.Println("Err: Error encountered during the scan.")
				return nil, err
			}

			// Add up to the UsernameList the fixedUsername.
			LikesList = append(LikesList, u)
		}
		// If we have encountered some error in the likes variable.
		if likes.Err() != nil {
			log.Println("Err: Error encountered on likes")
			return nil, err
		}

		// Then, check whether the returned list has length equal to zero. Return No Content if so.
		if len(LikesList) == 0 {
			return nil, ErrNoContent
		}

		// We now need to retrieve the Usernames from the fixedUsername followers List.
		var like_list []string
		var usr string
		for i := 0; i < len(LikesList); i++ {
			err := db.c.QueryRow(`SELECT username FROM Users WHERE fixedUsername == ?`, LikesList[i]).Scan(&usr)
			if !errors.Is(err, nil) {

				// If we encounter some error during the Username Retrieval.
				log.Println("Err: Error encountered during the Username retrieval in the DB.")
				return nil, err
			}

			// Add up to the UsernameList the Username.
			like_list = append(like_list, usr)

		}
		// If we arrive here, we have a result with at least one Username, return it.
		return like_list, nil

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
