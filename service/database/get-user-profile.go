package database

func (db *appdbimpl) GetUserProfile(fixedUsername string, uuid string) (User, error) {
	// Variable for returning the slice of fixedUsernames.
	var user User

	// Selection of the User profile. Here we can distinguish two cases:
	//1) We have that the User Profile must return sort of restricted data(i.e., without Personal Data if the User is requesting the action is not the owner.
	//2) We have that the User Profile must return all data of the profile if the User that is requesting the action is the owner of the profile.
	authorization, err1 := db.c.Query(`SELECT iif(uuid == '?', CAST("Authorized" AS TEXT), CAST("Not Authorized" AS TEXT))
	FROM Users
	WHERE fixedUsername == '?'`, uuid, fixedUsername)
	if err1 != nil {
		return user, err1
	}

	// Go checking whether you are authorized or not.
	for authorization.Next() {
		var auth string
		err1 = authorization.Scan(&auth)
		if err1 != nil {
			return user, err1
		} else {

			//If no errors occurs and you are Authorized, simply select all the User Row since the one that is requesting the row is the owner of the profile.
			if auth == "Authorized" {
				rows, err := db.c.Query(`SELECT * 
				FROM Users
				WHERE fixedUsername == '?'`, fixedUsername)

				// Here we read the resultset and we build the list to be returned.
				for rows.Next() {
					err = rows.Scan(&user.FixedUsername, &user.Uuid, &user.Username, &user.PhotoProfile, &user.Biography, &user.DateOfCreation, &user.NumberOfPhotos, &user.TotNumberLikes, &user.TotNumberComments, &user.NumberFollowers, &user.NumberFollowing, &user.Name, &user.Surname, &user.DateOfBirth, &user.Email, &user.Nationality, &user.Gender)
					if err != nil {
						return user, err
					}
				}
				if rows.Err() != nil {
					return user, err
				}
				return user, nil
			} else {
				//In the case no errors occurs but we have an "unauthorized" access, means that the user that is requesting the action is not the owner, thus must have a restricted View of the User profile (i.e, it must not see the Personal Data).
				//Notice that here, I could have also used a View instead of a query with all these selected columns. I have just opted for the worst choice, since it is very verbose.
				rows, err := db.c.Query(`SELECT fixedusername, username, photoprofile, biography, dateOfCreation, numberOfPhotos, totNumberLikes, totNumberComments, numberFollowers, numberFollowing
				FROM Users
				WHERE fixedUsername == '?'`, fixedUsername)

				// Here we read the resultset and we build the list to be returned.
				for rows.Next() {
					err = rows.Scan(&user.FixedUsername, &user.Uuid, &user.Username, &user.PhotoProfile, &user.Biography, &user.DateOfCreation, &user.NumberOfPhotos, &user.TotNumberLikes, &user.TotNumberComments, &user.NumberFollowers, &user.NumberFollowing)
					if err != nil {
						return user, err
					}
				}
				if rows.Err() != nil {
					return user, err
				}
				return user, nil
			}
		}
	}
	return user, nil
}
