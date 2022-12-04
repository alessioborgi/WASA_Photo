package database

import "fmt"

func (db *appdbimpl) GetBannedUsers(fixedUsername string, uuid string) ([]Ban, error) {
	// Variable for returning the slice of Bans.
	var banList []Ban

	// Selection of the Ban List. Here we can distinguish two cases:
	//1) When the User that is requesting the action is the profile owner that is requesting the action is not Banned. Return its Banned List.
	//2) When the User that is requesting the action is NOT the profile owner. Return an empty list since it is not authorized.
	authorization, errAuth := db.CheckAuthorizationOwner(fixedUsername, uuid)

	//Check for the error during the Query.
	if errAuth != nil {
		return nil, errAuth
	} else {

		// Go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
		if authorization == "Authorized" {

			//If you are the Owner of the Profile, select your Banned Users' List.
			//Note that in the Query below, we do not need to check also whether the uuid that is requesting the action is the profile owner, since we already have this certainty.
			bans, err := db.c.Query(`SELECT * 
			FROM Bans
			WHERE fixedUsernameBanner == '?'
			ORDER BY uploadDate DESC`, fixedUsername)

			//Check for the error during the Query.
			if err != nil {
				return nil, err
			} else {
				//Defer the bans closure. This is a Best-Practice.
				defer func() { _ = bans.Close() }()

				// Here we read the resultset and we build the list to be returned.
				for bans.Next() {
					var b Ban
					err = bans.Scan(&b.FixedUsernameBanner, &b.FixedUsernameBanned, &b.UploadDate, &b.Motivation)
					if err != nil {
						return nil, err
					}

					//Append to the banList if no error occurs.
					banList = append(banList, b)
				}
				if bans.Err() != nil {
					return nil, err
				} else {
					return banList, nil
				}
			}
		} else {
			//In the case you are not the profile owner, i.e. you result as "unauthorized", you must receive any BanList information.
			fmt.Println("You cannot have the Ban List you are requiring!")
			return nil, nil

		}
	}
}
