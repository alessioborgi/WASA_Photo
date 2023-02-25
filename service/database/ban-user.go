package database

import (
	"errors"
	"log"
)

func (db *appdbimpl) BanUser(username string, usernameBanned string, uuid string) (string, error) {

	// Adding a User Ban.
	// Here, you have 4 options, stored in the "authorization" variable:
	// 1) AUTHORIZED: The action requester is the Profile Owner. It can proceed to add the Ban only whether username and usernameBanned is not the same (we cannot self ban).
	// 2) UNAUTHORIZED: The action requester is NOT the Profile Owner. It cannot Ban.
	// 3) NOT VALID: The action requester has not inserted a valid Uuid, since it's not present in the DB.
	// 4) "": Returned if we have some errors.

	// 0.0) As a premature check, check whether the username that is requesting the action is going to self-ban.
	if username == usernameBanned {
		return "", ErrBadRequest
	}

	// 0.1) First of all, I need to check whether the username that wants to add the Ban exists (that must be also the uuid itself, check later).
	fixedUsernameBanner, errUsername := db.CheckUserPresence(username)

	// Check whether the Username I am trying to update, does not exists.
	if errors.Is(errUsername, ErrUserDoesNotExist) {
		log.Println("Err: The fixedUsernameBanner, does not exists.")
		return "", ErrUserDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errUsername, nil) && fixedUsernameBanner != "Not Exists" {
		log.Println("Err: Strange error during the Check of User Presence")
		return "", errUsername
	}

	// 0.2) Secondly, I need to check only whether the usernameBanned I passed exists in the DB.
	fixedUsernameBanned, errusernameBanned := db.CheckUserPresence(usernameBanned)

	// Check whether the usernameBanned I am trying to insert, already exists.
	if errors.Is(errusernameBanned, ErrUserDoesNotExist) {
		log.Println("Err: The errusernameBanned I am trying to update, does not exists.")
		return "", ErrUserDoesNotExist
	}

	// Check if strange errors occurs.
	if !errors.Is(errusernameBanned, nil) {
		log.Println("Err: Strange error during the Check of User Presence")
		return "", errusernameBanned
	}

	// 0.3) Thirdly, we should check whether there exists the same Ban already.
	ban_presence, errBanRetrieval := db.CheckBanPresence(fixedUsernameBanner, fixedUsernameBanned)
	if ban_presence == PRESENT {
		log.Println("Err: The Ban already exists.")
		return PRESENT, nil
	}

	// Check if strange errors occurs.
	if !errors.Is(errBanRetrieval, nil) && !errors.Is(errBanRetrieval, ErrBanDoesNotExist) {
		log.Println("Err: Strange error during the Check of Ban Presence")
		return "", errBanRetrieval
	}

	// If we arrive here, it means that the Ban is not present. Thus we can continue.
	// First of all, check the Authorization of the person who is asking the action.
	authorization, errAuth := db.CheckAuthorizationOwnerUsername(username, uuid)

	// Check for the error during the Query.
	if !errors.Is(errAuth, nil) {

		// Check whether we have received some errors during the Authentication.
		return "", errAuth
	}

	// We can now go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
	if authorization == AUTHORIZED {

		// If the uuid is requesting the action is the actual User Owner.
		// First of all you should proceed to eliminate the "follow" of the owner w.r.t. the banned and the "follow" of the banned w.r.t. the owner.

		// First, check whether there is a Follow from the fixedUsernameBanner and fixedUsernameBanned.
		errFollowRetrieval := db.CheckFollowPresence(fixedUsernameBanner, fixedUsernameBanned)
		if errors.Is(errFollowRetrieval, nil) {

			// Here I can proceed to eliminate the Follow.
			log.Println("The Follow between Banner and Banned exists.")

			_, errDeleteFollow1 := db.c.Exec(`DELETE FROM Follows WHERE fixedUsername = ? AND fixedUsernameFollowing = ?`, fixedUsernameBanner, fixedUsernameBanned)
			if !errors.Is(errDeleteFollow1, nil) {
				log.Println("Error Encountered during deletion of follow between Banner and Banned.")
				return "", errDeleteFollow1
			}

			// The Deletion went well.
			log.Println("Deletion between Banner and Banned correctly executed!")

		}

		// Check if strange errors occurs.
		if !errors.Is(errFollowRetrieval, nil) && !errors.Is(errFollowRetrieval, ErrFollowDoesNotExist) {
			log.Println("Err: Strange error during the Check of Follow Presence")
			return "", errFollowRetrieval
		}

		// Here I can now proceed to check if the converse Follow is present, i.e., if there is a follow from the fixedUsernameBanned to the fixedUsernameBanner
		errFollowRetrieval2 := db.CheckFollowPresence(fixedUsernameBanned, fixedUsernameBanner)
		if errors.Is(errFollowRetrieval2, nil) {

			// Here I can proceed to eliminate the Follow.
			log.Println("The Follow between Banned and Banner exists.")

			_, errDeleteFollow2 := db.c.Exec(`DELETE FROM Follows WHERE fixedUsername = ? AND fixedUsernameFollowing = ?`, fixedUsernameBanned, fixedUsernameBanner)
			if !errors.Is(errDeleteFollow2, nil) {
				log.Println("Error Encountered during deletion of follow between Banned and Banner.")
				return "", errDeleteFollow2
			}

			// The Deletion went well.
			log.Println("Deletion between Banned and Banner correctly executed!")
		}

		// Check if strange errors occurs.
		if !errors.Is(errFollowRetrieval2, nil) && !errors.Is(errFollowRetrieval2, ErrFollowDoesNotExist) {
			log.Println("Err: Strange error during the Check of Follow Presence")
			return "", errFollowRetrieval2
		}

		// Once we have eliminated follow, che can proceed in creating Ban.
		// If Authorized, you can proceed to add up the Ban Object without any problem.
		_, err := db.c.Exec(`INSERT INTO Bans (fixedUsernameBanner, fixedUsernameBanned) VALUES (?, ?)`, fixedUsernameBanner, fixedUsernameBanned)
		if err != nil {
			return "", err
		}

		// The Insertion went well.
		return NOTPRESENT, nil
	}

	// We can now see what to do if the Uuid that is requesting the action is not the User Owner.
	if authorization == NOTAUTHORIZED {

		// If the Use was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
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
