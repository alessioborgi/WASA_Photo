package database

// DOUBT: Do I have to pass it the entire User Object? Do I have to pass to it the new Username?

func (db *appdbimpl) SetPhoto(fixedUsername string, photoId int, newPhrase string, uuid string) error {
	// Selection of the User profile. Here we can distinguish two cases:
	//1) We have that the User Profile can be modified since the photo pertains to the username that is requesting the action.
	//2) We have that the User Profile cannot be modified since the photo DOES NOT pertains to the username that is requesting the action.
	//	 Here, you have also to check whether the User that is requesting the profile, has been Banned by the fixedUsername.
	authorization, errAuth := db.CheckAuthorizationOwner(fixedUsername, uuid)

	//Check for the error during the Query.
	if errAuth != nil {
		return errAuth
	} else {
		// Go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
		if authorization == "Authorized" {
			res, errUpdate := db.c.Exec(`UPDATE Photos SET phrase="?" WHERE fixedUsername = "?" AND photoid = "?"`, newPhrase, fixedUsername, photoId)
			if errUpdate != nil {
				return errUpdate
			}

			affected, err := res.RowsAffected()
			if err != nil {
				return err
			} else if affected == 0 {
				// If we didn't modified any row, then the User's Photo didn't exist.
				return ErrUserDoesNotExist
			}
			return nil

		} else {
			//If the User was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
			return ErrUserNotAuthorized
		}
	}
}
