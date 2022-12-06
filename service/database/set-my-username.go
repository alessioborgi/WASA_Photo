package database

// DOUBT: Do I have to pass it the entire User Object? Do I have to pass to it the new Username?

func (db *appdbimpl) SetMyUsername(fixedUsername string, newUsername string, uuid string) error {
	// Selection of the User profile. Here we can distinguish two cases:
	//1) We have that the User Profile must return sort of restricted data(i.e., without Personal Data if the User is requesting the action is not the owner.
	//2) We have that the User Profile must return all data of the profile if the User that is requesting the action is the owner of the profile.
	//	 Here, you have also to check whether the User that is requesting the profile, has been Banned by the fixedUsername.
	authorization, errAuth := db.CheckAuthorizationOwner(fixedUsername, uuid)

	//Check for the error during the Query.
	if errAuth != nil {
		return errAuth
	} else {
		// Go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
		if authorization == "Authorized" {
			res, errUpdate := db.c.Exec(`UPDATE Users SET username="?" WHERE fixedUsername = "?"`, newUsername, fixedUsername)
			if errUpdate != nil {
				return errUpdate
			}

			affected, err := res.RowsAffected()
			if err != nil {
				return err
			} else if affected == 0 {
				// If we didn't modified any row, then the User didn't exist.
				return ErrUserDoesNotExist
			}
			return nil

		} else {
			//If the Use was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
			return ErrUserNotAuthorized
		}
	}
}
