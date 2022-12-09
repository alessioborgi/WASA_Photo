package database

import "fmt"

func (db *appdbimpl) LikePhoto(like Like, uuid string) (Like, error) {

	// Addition of a User's Photo Like. Here we can distinguish two cases:
	//1) We have that the User can add the like on the photo since we have that it is NOT the photo owner, of course, and it has not been banned by the actual photo owner.  since the like action has been requested by the username that is requesting the action.
	//2) We have that the User cannot add the like to the photo since either it is the profile owner or it has been banned by the photo owner.
	//	 Here, you have also to check whether the User that is requesting the profile, has been Banned by the fixedUsername.
	authorization, errAuth := db.CheckAuthorizationOwner(like.Likeid, uuid)

	//Check for the error during the Query.
	if errAuth != nil {
		return like, errAuth
	} else {
		// Go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).

		if authorization == "Authorized" {
			//In the case the user result as "authorized", that in my project means to be the owner of the object on which we are oding the action, we cannot add the like.
			fmt.Println("You cannot add the Like Object you are requiring!")
			return like, ErrUserProfileOwner
		} else {
			//If the User was not "Authorized", i.e. it is not the Profile Owner, we can proceed instead to control whether it has been banned by the user owner.
			ban, errBan := db.CheckBan(like.FixedUsername, uuid)
			//Check for the error during the Query.
			if errBan != nil {
				return like, errBan
			} else {

				//If no error occurs, checking whether the user(uuid) was banned by the fixedUsername.
				if ban == "Not Banned" {
					//If Not Banned, you can add the like object without any problem.
					_, err := db.c.Exec(`INSERT INTO Likes (likeid, photoid, fixedUsername, uploadDate) VALUES (?, ?, ?, ?)`,
						like.Likeid, like.PhotoId, like.FixedUsername, like.UploadDate)
					if err != nil {
						return like, err
					} else {
						return like, nil
					}
				} else {
					//If the User was Banned instead, returns nothing.
					fmt.Println("You cannot add the Like Object you are requiring!")
					return like, ErrUserNotAuthorized
				}
			}
		}
	}
}
