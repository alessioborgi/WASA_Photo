package database

func (db *appdbimpl) GetUserProfile(fixedUsername string, uuid string) (User, error) {
	// Variable for returning the UserProfile.
	var user User

	// Selection of the User profile. Here we can distinguish two cases:
	//1) We have that the User Profile must return sort of restricted data(i.e., without Personal Data if the User is requesting the action is not the owner.
	//2) We have that the User Profile must return all data of the profile if the User that is requesting the action is the owner of the profile.
	//	 Here, you have also to check whether the User that is requesting the profile, has been Banned by the fixedUsername.
	authorization, errAuth := db.CheckAuthorizationOwner(fixedUsername, uuid)

	//Check for the error during the Query.
	if errAuth != nil {
		return User{}, errAuth
	} else {
		// Go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
		if authorization == "Authorized" {

			//If you are the Owner of the Profile.
			err := db.c.QueryRow(`SELECT * 
			FROM Users
			WHERE fixedUsername == '?'`, fixedUsername).Scan(&user.FixedUsername, &user.Uuid, &user.Username, &user.PhotoProfile, &user.Biography, &user.DateOfCreation, &user.NumberOfPhotos, &user.TotNumberLikes, &user.TotNumberComments, &user.NumberFollowers, &user.NumberFollowing, &user.Name, &user.Surname, &user.DateOfBirth, &user.Email, &user.Nationality, &user.Gender)

			//Check for the error during the Query.
			if err != nil {
				return User{}, err
			} else {
				return user, nil
			}
		} else {
			//In the case you are not the profile owner, i.e. you result as "unauthorized", you must have a restricted View of the User profile (i.e, it must not see the Personal Data).

			//Check first whether the user that is requesting the action has been banned by the fixedUsername.
			ban, errBan := db.CheckBan(fixedUsername, uuid)
			//Check for the error during the Query.
			if errBan != nil {
				return User{}, errBan
			} else {

				//Checking whether the user was banned by the fixedUsername.
				if ban == "Not Banned" {
					//If we are Not Banned, we proceed to return the User Data (restricted View).
					//Notice that here, I could have also used a View instead of a query with all these selected columns. I have just opted for the worst choice, since it is very verbose.
					err := db.c.QueryRow(`SELECT fixedusername, username, photoprofile, biography, dateOfCreation, numberOfPhotos, totNumberLikes, totNumberComments, numberFollowers, numberFollowing
					FROM Users
					WHERE fixedUsername == '?'`, fixedUsername).Scan(&user.FixedUsername, &user.Uuid, &user.Username, &user.PhotoProfile, &user.Biography, &user.DateOfCreation, &user.NumberOfPhotos, &user.TotNumberLikes, &user.TotNumberComments, &user.NumberFollowers, &user.NumberFollowing)

					//Check for the error during the Query.
					if err != nil {
						return User{}, err
					} else {
						return user, nil
					}
				} else {
					//If the Use was Banned instead, return the empty User.
					return User{}, nil
				}
			}
		}
	}
}