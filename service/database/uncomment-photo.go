package database

func (db *appdbimpl) UncommentPhoto(fixedUsername string, photoId int, commentId int, uuid string) error {
	// Deletion of a User Photo's Comment. Here we can distinguish two cases:
	//1) We have that the User's Photo Comment can can be deleted since the user is requesting the action is the user owner.
	//2) We have that the User's Photo Comment cannot can be deleted since the user is requesting the action is NOT the user owner.
	authorization, errAuth := db.CheckAuthorizationOwner(fixedUsername, uuid)

	//Check for the error during the Query.
	if errAuth != nil {
		return errAuth
	} else {
		if authorization == "Authorized" {

			//Here I need first to obtain the fixedusername of the Uuid is requesting the action, in such a way to see whether it the maker of the comment.
			var commenterFixedUsername string
			err := db.c.QueryRow(`SELECT fixedUsername 
			FROM Users 
			WHERE uuid == '?'`, uuid).Scan(&commenterFixedUsername)

			if err != nil {
				return err
			} else {

				//Here we proceed in the deletion of the Comment.
				res, errDeletion := db.c.Exec(`DELETE FROM Comments WHERE fixedUsername="?" AND photoid = "?" AND commentId = "?" AND commenterFixedUsername = "?"`, fixedUsername, photoId, commentId, commenterFixedUsername)
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
			}

		} else {
			//If the User was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
			return ErrUserNotAuthorized
		}
	}
}
