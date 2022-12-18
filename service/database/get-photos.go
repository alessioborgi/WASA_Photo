package database

// import "fmt"

// func (db *appdbimpl) GetPhotos(fixedUsername string, uuid string) ([]Photo, error) {
// 	// Variable for returning the slice of Photos.
// 	var photoList []Photo

// 	// Selection of the Photo List. Here we can distinguish two cases:
// 	//1) When the User that is requesting the action is the profile owner or when the user that is requesting the action is not Banned. Return all the list.
// 	//2) When the User that is requesting the action is NOT the profile owner and it has been Banned by the fixedUsername. Return an empty list
// 	authorization, errAuth := db.CheckAuthorizationOwner(fixedUsername, uuid)

// 	//Check for the error during the Query.
// 	if errAuth != nil {
// 		return nil, errAuth
// 	} else {

// 		// Go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
// 		// Here we don't need to check whether you are banned since it is not possible to have a self-ban, of course.
// 		if authorization == "Authorized" {

// 			//If you are the Owner of the Profile, select all the photos.
// 			photos, err := db.c.Query(`SELECT *
// 			FROM Photos
// 			WHERE fixedUsername == '?'
// 			ORDER BY uploadDate DESC`, fixedUsername)

// 			//Check for the error during the Query.
// 			if err != nil {
// 				return nil, err
// 			} else {
// 				//Defer the Photos closure. This is a Best-Practice.
// 				defer func() { _ = photos.Close() }()

// 				// Here we read the resultset and we build the list to be returned.
// 				for photos.Next() {
// 					var p Photo
// 					err = photos.Scan(&p.Photoid, &p.FixedUsername, &p.Filename, &p.UploadDate, &p.Phrase, &p.NumberLikes, &p.NumberComments, &p.Latitude, &p.Longitude)
// 					if err != nil {
// 						return nil, err
// 					}

// 					//Append to the photoList if no error occurs.
// 					photoList = append(photoList, p)
// 				}
// 				if photos.Err() != nil {
// 					return nil, err
// 				} else {
// 					return photoList, nil
// 				}
// 			}
// 		} else {
// 			//In the case you are not the profile owner, i.e. you result as "unauthorized", you must be Not Banned if you want to see the User photos.

// 			//Check first whether the user that is requesting the action has been banned by the fixedUsername.
// 			ban, errBan := db.CheckBan(fixedUsername, uuid)
// 			//Check for the error during the Query.
// 			if errBan != nil {
// 				return nil, errBan
// 			} else {

// 				//If no error occurs, checking whether the user was banned by the fixedUsername.
// 				if ban == "Not Banned" {
// 					//If Not Banned, you can have the list of photos as in the above case.
// 					photos, err := db.c.Query(`SELECT *
// 					FROM Photos
// 					WHERE fixedUsername == '?'
// 					ORDER BY uploadDate DESC`, fixedUsername)

// 					//Check for the error during the Query.
// 					if err != nil {
// 						return nil, err
// 					} else {
// 						//Defer the Photos closure. This is a Best-Practice.
// 						defer func() { _ = photos.Close() }()

// 						// Here we read the resultset and we build the list to be returned.
// 						for photos.Next() {
// 							var p Photo
// 							err = photos.Scan(&p.Photoid, &p.FixedUsername, &p.Filename, &p.UploadDate, &p.Phrase, &p.NumberLikes, &p.NumberComments, &p.Latitude, &p.Longitude)
// 							if err != nil {
// 								return nil, err
// 							}

// 							//Append to the photoList if no error occurs.
// 							photoList = append(photoList, p)
// 						}
// 						if photos.Err() != nil {
// 							return nil, err
// 						} else {
// 							return photoList, nil
// 						}
// 					}

// 				} else {
// 					//If the Use was Banned instead, returns nothing.
// 					fmt.Println("You cannot have the PhotoList you are requiring!")
// 					return nil, nil
// 				}
// 			}
// 		}
// 	}
// }
