package database

import "fmt"

func (db *appdbimpl) GetMyStream(fixedUsername string, uuid string) ([]Photo, error) {
	// Variable for returning the slice of Photos.
	var myStream []Photo

	// Selection of the MyStream Photo List. Here we can distinguish two cases:
	//1) When the User that is requesting the action is the profile owner. Return all the list of Followings' Photos.
	//2) When the User that is requesting the action is NOT the profile owner. Return an empty list.
	authorization, errAuth := db.CheckAuthorizationOwner(fixedUsername, uuid)

	//Check for the error during the Query.
	if errAuth != nil {
		return nil, errAuth
	} else {

		// Go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
		// Here we don't need to check whether you are banned since it is not possible to have a self-ban, of course.
		if authorization == "Authorized" {

			//If you are the Owner of the Profile, select all the photos of you Followings.
			photos, err := db.c.Query(`SELECT * 
			FROM Photos
			WHERE fixedUsername IN
			(SELECT fixedUsername 
			FROM Follows
			WHERE fixedUsernameFollowing == '?')
			ORDER BY uploadDate DESC`, fixedUsername)

			//Check for the error during the Query.
			if err != nil {
				return nil, err
			} else {
				//Defer the Photos closure. This is a Best-Practice.
				defer func() { _ = photos.Close() }()

				// Here we read the resultset and we build the list to be returned.
				for photos.Next() {
					var p Photo
					err = photos.Scan(&p.Photoid, &p.FixedUsername, &p.Filename, &p.UploadDate, &p.Phrase, &p.NumberLikes, &p.NumberComments, &p.Latitude, &p.Longitude)
					if err != nil {
						return nil, err
					}

					//Append to the myStream if no error occurs.
					myStream = append(myStream, p)
				}
				if photos.Err() != nil {
					return nil, err
				} else {
					return myStream, nil
				}
			}
		} else {
			//In the case you are not the profile owner, i.e. you result as "unauthorized", you must not see any other's Streams.
			fmt.Println("You cannot have the PhotoList you are requiring!")
			return nil, ErrUserNotAuthorized
		}
	}
}
