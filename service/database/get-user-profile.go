package database

import (
	"errors"
	"log"
)

func (db *appdbimpl) GetUserProfile(username string, uuid string) (User, error) {

	// Selection of the User profile.
	// Here, you have 4 options, stored in the "authorization" variable:
	// 1) AUTHORIZED: The action requester is the Profile Owner.
	// 2) UNAUTHORIZED: The action requester is NOT the Profile Owner.
	// 3) NOT VALID: The action requester has not inserted a valid Uuid, since it's not present in the DB.
	// 4) "": Returned if we have some errors.
	// First of all, check the Authorization of the person who is asking the action .
	authorization, errAuth := db.CheckAuthorizationOwnerUsername(username, uuid)

	// Check for the error during the Query.
	if !errors.Is(errAuth, nil) {

		// Returning empty user and error if some errors occurs.
		return User{}, errAuth
	}

	// If arrive here, you have encountered No Errors.

	// Variable for returning the UserProfile.
	var user User

	// 1) First, check the authorization. In the you are AUTHORIZED and NOTAUTHORIZED, you are authorized to view all authorization information.
	if authorization == AUTHORIZED || authorization == NOTAUTHORIZED {

		// First of all, I need to check whether the username on which uuid wants to do the action exists.
		fixedUsername, errUsername := db.CheckUserPresence(username)

		// Check whether the fixedUsername I am trying to delete, does not exists.
		if errors.Is(errUsername, ErrUserDoesNotExist) {
			log.Println("Err: The fixedUsername I am trying to get the profile of, does not exists.")
			return User{}, ErrUserDoesNotExist
		}

		// Check if strange errors occurs.
		if !errors.Is(errUsername, nil) && !errors.Is(errUsername, Okay_Error_Inverse) {
			log.Println("Err: Strange error during the Check of User Presence")
			return User{}, errUsername
		}

		// The User Exists.
		log.Println("The User Exists in the WASAPhoto Platform!")

		// If the User was not "Authorized", i.e. it is not the Profile Owner, it must be checked whether you are banned or not.
		// 0.1) First of all, I need to check whether the username that wants to know the Followers of exists.
		fixedUsernameRequester, errfixedUsernameRequester := db.GetFixedUsername(uuid)

		// Check if strange errors occurs.
		if !errors.Is(errfixedUsernameRequester, nil) {
			log.Println("Err: Unexpected Error in the Username Requester Retrieval ")
			return User{}, errfixedUsernameRequester
		}

		// If we arrive here, we have correclty retrieved the Requester Username.
		// Proceed to check whether it is Banned or not.
		errBanRetrieval := db.CheckBanPresence(fixedUsername, fixedUsernameRequester)
		if errors.Is(errBanRetrieval, Okay_Error_Inverse) {
			log.Println("Err: The Ban exists. You cannot get its Profile.")
			return User{}, ErrUserNotAuthorized
		}

		// Check if strange errors occurs.
		if !errors.Is(errBanRetrieval, nil) && !errors.Is(errBanRetrieval, ErrBanDoesNotExist) {
			log.Println("Err: Strange error during the Check of Ban Presence")
			return User{}, errBanRetrieval
		}

		// If we arrive here, the user is not Banned and we can retrieve all the Profile Data from the DB.
		err := db.c.QueryRow(`SELECT *
							FROM Users
							WHERE username == ?`, username).Scan(&user.FixedUsername, &user.Uuid, &user.Username, &user.Biography, &user.DateOfCreation, &user.NumberOfPhotos, &user.NumberFollowers, &user.NumberFollowing, &user.Name, &user.Surname, &user.DateOfBirth, &user.Email, &user.Nationality, &user.Gender)

		// Check for the error during the Query.
		if err != nil {

			// If we have encountered some errors in the Query retrieval.
			log.Println("Err: Unexpected Error! During the Query Retrieval!")
			return User{}, err
		}

		// Otherwise we have retrieved the User Profile Correctly
		log.Println("User Profile retrieved correctly!")
		return user, nil

	} else if authorization == NOTVALID {

		log.Println("Err: The Uuid you are providing or the fixedUsername is not present.")
		// If we arrive here, the Uuid is not present in the DB.
		return User{}, ErrUserNotAuthorized
	} else {

		// If we arrive here, we encountered other types of problems.
		return User{}, errAuth
	}
}
