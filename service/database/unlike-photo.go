package database

func (db *appdbimpl) UnlikePhoto(fixedUsername string, photoId int, fixedUsernameLiker string, uuid string) error {
	// Deletion of a User Photo's Like. Here we can distinguish two cases:
	//1) We have that the User's Photo Like can can be deleted since the user is requesting the action is the user owner.
	//2) We have that the User's Photo Like cannot can be deleted since the user is requesting the action is NOT the user owner.
	authorization, errAuth := db.CheckAuthorizationOwner(fixedUsername, uuid)

	//Check for the error during the Query.
	if errAuth != nil {
		return errAuth
	} else {
		if authorization == "Authorized" {
			res, errDeletion := db.c.Exec(`DELETE FROM Likes WHERE fixedUsername="?" AND photoid = "?" AND likeid = "?"`, fixedUsername, photoId, fixedUsernameLiker)
			if errDeletion != nil {
				return errDeletion
			}

			affected, err := res.RowsAffected()
			if err != nil {
				return err
			} else if affected == 0 {
				// If we didn't delete any row, then the User didn't exist.
				return ErrUserDoesNotExist
			}
			return nil

		} else {
			//If the User was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
			return ErrUserNotAuthorized
		}
	}
}