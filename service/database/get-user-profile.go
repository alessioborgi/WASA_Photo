package database

import (
	"database/sql"
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

	//Check for the error during the Query.
	if errAuth != nil {

		//Returning empty user and error if some errors occurs.
		return User{}, errAuth
	}

	// If arrive here, you have encountered No Errors.

	// Variable for returning the UserProfile.
	var user User

	// 1) First, check the authorization. In the you are AUTHORIZED and NOTAUTHORIZED, you are authorized to view all authorization information.
	if authorization == AUTHORIZED || authorization == NOTAUTHORIZED {

		// First, check whether the username you are searching is present.
		var exists = 0
		err := db.c.QueryRow(`SELECT COUNT(fixedUsername) FROM Users WHERE username == ?`, username).Scan(&exists)

		// Check whether we have encountered some error during the Query.
		if err != nil && err != sql.ErrNoRows {
			log.Println("Err: Unexpected Error during the Query!")
			return User{}, err
		}

		// Check whether we have the username we are searching for in the DB.
		if exists == 1 {

			// The User Exists.
			log.Println("The User Exists in the WASAPhoto Platform!")

			// Retrieve all the Profile Data from the DB.
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
		} else {

			// The User Does not exists(exists = 0).
			log.Println("Err: User Does Not Exist!")
			return User{}, ErrUserDoesNotExist
		}
	} else if authorization == NOTVALID {

		log.Println("Err: The Uuid you are providing or the fixedUsername is not present.")
		// If we arrive here, the Uuid is not present in the DB.
		return User{}, ErrUserNotAuthorized
	} else {

		// If we arrive here, we encountered other types of problems.
		return User{}, errAuth
	}

	// }
	// } else {
	//In the case you are not the profile owner, i.e. you result as "unauthorized", you must have a restricted View of the User profile (i.e, it must not see the Personal Data).

	//Check first whether the user that is requesting the action has been banned by the fixedUsername.

	// ban, errBan := db.CheckBan(fixedUsername, uuid)

	//Check for the error during the Query.
	// if errBan != nil {
	// 	return User{}, errBan
	// } else {

	//Checking whether the user was banned by the fixedUsername.
	// if ban == "Not Banned" {
	//If we are Not Banned, we proceed to return the User Data (restricted View).
	//Notice that here, I could have also used a View instead of a query with all these selected columns. I have just opted for the worst choice, since it is very verbose.
	// err := db.c.QueryRow(`SELECT fixedusername, username, biography, dateOfCreation, numberOfPhotos, totNumberLikes, totNumberComments, numberFollowers, numberFollowing
	// FROM Users
	// WHERE username == '?'`, username).Scan(&user.FixedUsername, &user.Uuid, &user.Username, &user.Biography, &user.DateOfCreation, &user.NumberOfPhotos, &user.TotNumberLikes, &user.TotNumberComments, &user.NumberFollowers, &user.NumberFollowing)

	// //Check for the error during the Query.
	// if err != nil {
	// 	return User{}, err
	// } else {
	// 	return user, nil
	// }
	// 			} else {
	// 				//If the Use was Banned instead, return the empty User.
	// 				return User{}, nil
	// 			}
	// 		}
	// 	}
	// }
}
