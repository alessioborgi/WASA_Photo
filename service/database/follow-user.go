package database

func (db *appdbimpl) FollowUser(follow Follow, uuid string) (Follow, error) {

	// Addition of the User profile. Here we can distinguish two cases:
	//1) We have that the User can add the follow since the follow action has been requested by the username that is requesting the action.
	//2) We have that the User cannot add the follow since the follow action has NOT been requested by the username that is requesting the action.
	//	 Here, you have also to check whether the User that is requesting the profile, has been Banned by the fixedUsername.
	authorization, errAuth := db.CheckAuthorizationOwner(fixedUsername, uuid)

	//Check for the error during the Query.
	if errAuth != nil {
		return errAuth
	} else {
		// Go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).

		// HERE I NEED TO RETURN THE FIXEDUSERNAME OF THE COMMENTER!!!!!!

		// TO DOOOOOO
		if authorization == "Authorized" {
			res, errUpdate := db.c.Exec(`UPDATE Comments SET phrase="?" WHERE fixedUsername = "?" AND photoid = "?" AND commentId = "?"`, newComment, fixedUsername, photoId, commentId)
			if errUpdate != nil {
				return errUpdate
			}

			affected, err := res.RowsAffected()
			if err != nil {
				return err
			} else if affected == 0 {
				// If we didn't modified any row, then the User's Comment didn't exist.
				return ErrPhotoDoesNotExist
			}
			return nil

		} else {
			//If the User was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
			return ErrUserNotAuthorized
		}
	}
}
