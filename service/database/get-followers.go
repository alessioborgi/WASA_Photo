package database

import "fmt"

func (db *appdbimpl) GetFollowers(fixedUsername string, uuid string) ([]Follow, error) {
	// Variable for returning the slice of Followers.
	var followersList []Follow

	// Selection of the Follow List. Here we can distinguish two cases:
	//1) When the User that is requesting the action is the profile owner or when the user that is requesting the action is not Banned. Return all the list of Followers.
	//2) When the User that is requesting the action is NOT the profile owner and it has been Banned by the fixedUsername. Return an empty list.
	authorization, errAuth := db.CheckAuthorizationOwner(fixedUsername, uuid)

	//Check for the error during the Query.
	if errAuth != nil {
		return nil, errAuth
	} else {

		// Go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
		// Here we don't need to check whether you are banned since it is not possible to have a self-ban, of course.
		if authorization == "Authorized" {

			//If you are the Owner of the Profile, select all the follow objects where the fixedUsername(person that is followed) is the user.
			//Remember always that here the fixedUsername is the person that is followed, while fixedUsernameFollowing is the person that is Following.
			followers, err := db.c.Query(`SELECT * 
			FROM Follows
			WHERE fixedUsername == '?'
			ORDER BY uploadDate DESC`, fixedUsername)

			//Check for the error during the Query.
			if err != nil {
				return nil, err
			} else {
				//Defer the followers closure. This is a Best-Practice.
				defer func() { _ = followers.Close() }()

				// Here we read the resultset and we build the list to be returned.
				for followers.Next() {
					var f Follow
					err = followers.Scan(&f.FixedUsername, &f.FixedUsernameFollowing, &f.UploadDate)
					if err != nil {
						return nil, err
					}

					//Append to the followersList if no error occurs.
					followersList = append(followersList, f)
				}
				if followers.Err() != nil {
					return nil, err
				} else {
					return followersList, nil
				}
			}
		} else {
			//In the case you are not the profile owner, i.e. you result as "unauthorized", you must be Not Banned if you want to see the User's Followers.

			//Check first whether the user that is requesting the action has been banned by the fixedUsername.
			ban, errBan := db.CheckBan(fixedUsername, uuid)
			//Check for the error during the Query.
			if errBan != nil {
				return nil, errBan
			} else {

				//If no error occurs, checking whether the user was banned by the fixedUsername.
				if ban == "Not Banned" {
					//If you are not banned, select all the User's Followers.
					followers, err := db.c.Query(`SELECT * 
					FROM Follows
					WHERE fixedUsername == '?'
					ORDER BY uploadDate DESC`, fixedUsername)

					//Check for the error during the Query.
					if err != nil {
						return nil, err
					} else {
						//Defer the followers closure. This is a Best-Practice.
						defer func() { _ = followers.Close() }()

						// Here we read the resultset and we build the list to be returned.
						for followers.Next() {
							var f Follow
							err = followers.Scan(&f.FixedUsername, &f.FixedUsernameFollowing, &f.UploadDate)
							if err != nil {
								return nil, err
							}

							//Append to the followersList if no error occurs.
							followersList = append(followersList, f)
						}
						if followers.Err() != nil {
							return nil, err
						} else {
							return followersList, nil
						}
					}
				} else {
					//If the User was Banned instead, returns nothing.
					fmt.Println("You cannot have the Followers List you are requiring!")
					return nil, nil
				}
			}
		}
	}
}
