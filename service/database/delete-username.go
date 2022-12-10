package database

import "log"

func (db *appdbimpl) DeleteUsername(username string) error {
	// Deletion of the User profile. Here we can distinguish two cases:
	//1) We have that the User Profile can can be deleted since the user is requesting the action is the user owner.
	//2) We have that the User Profile cannot can be deleted since the user is requesting the action is NOT the user owner.
	// authorization, errAuth := db.CheckAuthorizationOwner(fixedUsername, uuid)

	//Check for the error during the Query.
	// if errAuth != nil {
	// 	return errAuth
	// } else {
	// 	if authorization == "Authorized" {

	// Perform the actual Deletion of the User profile from the DB.
	res, errDeletion := db.c.Exec(`DELETE FROM Users WHERE username=?`, username)
	if errDeletion != nil {
		log.Fatalf("Error encountered during the User Deletion in the DB.")
		return errDeletion
	} else {

		// If we arrive here, no error occurred.
		// Check whether the deletion action is actually happened.
		affected, err := res.RowsAffected()
		if err != nil {
			return err
		} else if affected == 0 {

			// If we didn't delete any row, then the User didn't exist.
			return ErrUserDoesNotExist
		} else {

			// If we are here, the Deletion Action has actually deleted the User Profile.
			log.Println("User correctly deleted from the DB")
			return nil
		}

		// 	} else {
		// 		//If the User was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
		// 		return ErrUserNotAuthorized
		// 	}
		// }
	}
}
