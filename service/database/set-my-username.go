package database

// DOUBT: Do I have to pass it the entire User Object? Do I have to pass to it the new Username?

func (db *appdbimpl) SetMyUsername(fixedUsername string, newUsername string) error {
	// Selection of the User profile. Here we can distinguish two cases:
	// 1) We have that the User Username can be modified since the username that is requesting the action is the Profile Owner.
	// 2) We have that the User Username cannot be modified since the username that is requesting the action is NOT the Profile Owner.

	// Here, you have also to check whether the User that is requesting the profile, has been Banned by the fixedUsername.
	// authorization, errAuth := db.CheckAuthorizationOwner(username, uuid)

	// Check for the error during the Query.
	// if errAuth != nil {

	// 	// Check whether we have received some errors during the Authentication.
	// 	return errAuth
	// } else if authorization == AUTHORIZED {

	// Go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).

	// First of all, I need to check whether the fixedUsername on which uuid wants to do the action exists.
	// username, errUsername := db.CheckFixedUserPresence(fixedUsername)

	// Check if strange errors occurs.
	// if errUsername != nil && errUsername != ErrUserDoesNotExist {
	// 	return errUsername
	// }

	// Check whether theUsername I am trying to update, does not exists.
	// if errUsername == ErrUserDoesNotExist {
	// 	return ErrUserDoesNotExist
	// }

	// Here I arrive if the Username I am trying to update exists. I have the fixedUsername passed in input.

	// Perform the Update of the Username.
	res, errUpdate := db.c.Exec(`UPDATE Users SET username=? WHERE fixedUsername = ?`, newUsername, fixedUsername)

	// Check if some strage error occurred during the update.
	if errUpdate != nil {
		return errUpdate
	}

	// Here arrives if no strange errors occurred.
	affected, err := res.RowsAffected()
	if err != nil {

		return err
	} else if affected == 0 {
		// If we didn't modified any row, then the User didn't exist.
		return ErrUserDoesNotExist
	}
	return nil

	// } else if authorization == NOTAUTHORIZED {
	// 	//If the Use was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
	// 	return ErrUserNotAuthorized
	// } else {

	// 	// Here we are in the case in which we have that the authorization = NOTVALID, i.e., we have that the requesting user
	// 	// has inserted a Not-Valid Uuid
	// 	return nil
	// }
}
