package database

import "fmt"

func (db *appdbimpl) FollowUser(follow Follow, uuid string) (Follow, error) {

	// Addition of the User profile. Here we can distinguish two cases:
	//1) We have that the User can add the follow since the follow action has been requested by the username that is requesting the action.
	//2) We have that the User cannot add the follow since the follow action has NOT been requested by the username that is requesting the action.
	//	 Here, you have also to check whether the User that is requesting the profile, has been Banned by the fixedUsername.
	authorization, errAuth := db.CheckAuthorizationOwner(follow.FixedUsername, uuid)

	//Check for the error during the Query.
	if errAuth != nil {
		return follow, errAuth
	} else {
		// Go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).

		if authorization == "Authorized" {
			//Check first whether the (uuid)user that is requesting the action has been banned by the fixedUsernameFollowing.
			ban, errBan := db.CheckBan(follow.FixedUsernameFollowing, uuid)
			//Check for the error during the Query.
			if errBan != nil {
				return follow, errBan
			} else {

				//If no error occurs, checking whether the user was banned by the fixedUsername.
				if ban == "Not Banned" {
					//If Not Banned, you can add the follow object without any problem.
					_, err := db.c.Exec(`INSERT INTO Follows (fixedUsername, fixedUsernameFollowing, uploadDate) VALUES (?, ?, ?, ?)`,
						follow.FixedUsername, follow.FixedUsernameFollowing, follow.UploadDate)
					if err != nil {
						return follow, err
					} else {
						return follow, nil
					}
				} else {
					//If the Use was Banned instead, returns nothing.
					fmt.Println("You cannot have the PhotoList you are requiring!")
					return follow, ErrUserNotAuthorized
				}
			}
		} else {
			//If the User was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
			return follow, ErrUserNotAuthorized
		}
	}
}
