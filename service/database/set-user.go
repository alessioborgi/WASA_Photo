package database

import (
	"errors"
	"log"
)

func (db *appdbimpl) SetUser(username string, user User, uuid string) (string, error) {

	// Selection of the User profile.
	// Here, you have 4 options, stored in the "authorization" variable:
	// 1) AUTHORIZED: The action requester is the Profile Owner.
	// 2) UNAUTHORIZED: The action requester is NOT the Profile Owner.
	// 3) NOT VALID: The action requester has not inserted a valid Uuid, since it's not present in the DB.
	// 4) "": Returned if we have some errors.

	// 0.1) First of all, I need to check whether the username on which uuid wants to do the action exists (that must be also the uuid itself, check later).
	// This is for getting the fixedUsername from the username.
	fixedUsername, errFixedUsername := db.GetFixedUsername(uuid)

	// Check whether the Username I am trying to update with the newUsername, does not exists.
	if !errors.Is(errFixedUsername, nil) {
		log.Println("Err: Error!")
		return "", ErrBadRequest
	}

	// Let's now check whether the Username we want to insert is already present in the WASAPhoto Platform.
	user_presence, errUsername := db.CheckUserPresence(user.Username)

	// Check whether the Username I am trying to update with the newUsername, does not exists.
	// Here, the user must be first not null, then it must be not a present user, except for the case where it is actually the user owner.
	// This means that it can change also some profile things, without necessarily changing the username.
	if user_presence == "" {
		log.Println("Err: The newUsername is empty. Error!")
		return "", ErrBadRequest
	}

	if user_presence != NOTEXISTS && username != user.Username {
		// If the username is not-existing, we are ok.
		// We are ok also if the username is existing and it is the actual user owner. This means that the user is changing something and not the username in its profile.
		log.Println("Err: The newUsername is already a WASAPhoto Username. Error!")
		return "", ErrBadRequest
	}

	// Check if strange errors occurs.
	if !errors.Is(errUsername, nil) && !errors.Is(errUsername, ErrUserDoesNotExist) {
		log.Println("Err: Strange error during the Check of User Presence")
		return "", errUsername
	}

	// Set the fixedUsername of the new user.
	user.FixedUsername = fixedUsername

	// If both the Usernames are ok, check the Authorization of the person who is asking the action.
	authorization, errAuth := db.CheckAuthorizationOwnerUsername(username, uuid)

	// Check for the error during the Query.
	if !errors.Is(errAuth, nil) {

		// Check whether we have received some errors during the Authentication.
		return "", errAuth
	}

	// We can now go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
	if authorization == AUTHORIZED {

		// If the uuid is requesting the action is the actual User Owner.
		log.Println("The User is Authorized to change its own Profile Information.")

		// The first thing to do is the following: get the user.photoProfile.
		var oldPhotoPath string
		err := db.c.QueryRow(`SELECT photoProfile
							FROM Users
							WHERE username == ?`, username).Scan(&oldPhotoPath)

		// Check for the error during the Query.
		if err != nil {

			// If we have encountered some errors in the Query retrieval.
			log.Println("Err: Unexpected Error! During the Query Retrieval!")
			return "", err
		}

		// Otherwise we have retrieved the User Profile Correctly
		log.Println("User Profile retrieved correctly!")

		// Perform the Update of the User.
		_, errUpdate := db.c.Exec(`UPDATE Users SET username=?, photoProfile=?, biography=?, name=?, surname=?, dateOfBirth=?, email=?, nationality=?, gender=? WHERE fixedUsername=?`,
			user.Username, user.PhotoProfile, user.Biography, user.Name, user.Surname, user.DateOfBirth, user.Email, user.Nationality, user.Gender, fixedUsername)

		// Check if some strage error occurred during the update.
		if !errors.Is(errUpdate, nil) {
			log.Println("Err: Error during Update.")
			return "", errUpdate
		}

		return oldPhotoPath, nil
	}

	// We can now see what to do if the Uuid that is requesting the action is not the User Owner.
	if authorization == NOTAUTHORIZED {

		// If the User was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
		log.Println("Err: The Uuid you are providing is not Authorized to do this action.")
		return "", ErrUserNotAuthorized
	}

	// Check if we have a NOTVALID auth, i.e., the Uuid is not present in the DB.
	if authorization == NOTVALID {

		log.Println("Err: The Uuid you are providing is not present.")
		return "", ErrUserNotAuthorized
	}

	// If we arrive here, we encountered other types of problem.
	log.Println("Err: Unexpected Error.")
	return "", errAuth
}
