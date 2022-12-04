package database

import "fmt"

func (db *appdbimpl) GetPhotoComments(fixedUsername string, photoId int, uuid string) ([]Comment, error) {
	// Variable for returning the slice of Photos.
	var commentList []Comment

	// Selection of the Comment List. Here we can distinguish two cases:
	//1) When the User that is requesting the action is the profile owner or when the user that is requesting the action is not Banned. Return all the list of Photo's Comments.
	//2) When the User that is requesting the action is NOT the profile owner and it has been Banned by the fixedUsername. Return an empty list.
	authorization, errAuth := db.CheckAuthorizationOwner(fixedUsername, uuid)

	//Check for the error during the Query.
	if errAuth != nil {
		return nil, errAuth
	} else {

		// Go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).
		// Here we don't need to check whether you are banned since it is not possible to have a self-ban, of course.
		if authorization == "Authorized" {

			//If you are the Owner of the Profile, select all the photo's comments.
			comments, err := db.c.Query(`SELECT * 
			FROM Comments 
			WHERE fixedUsername == '?' AND photoId == '?'
			ORDER BY uploadDate DESC`, fixedUsername, photoId)

			//Check for the error during the Query.
			if err != nil {
				return nil, err
			} else {
				//Defer the comments closure. This is a Best-Practice.
				defer func() { _ = comments.Close() }()

				// Here we read the resultset and we build the list to be returned.
				for comments.Next() {
					var c Comment
					err = comments.Scan(&c.Commentid, &c.PhotoId, &c.FixedUsername, &c.Phrase, &c.CommenterFixedUsername, &c.UploadDate)
					if err != nil {
						return nil, err
					}

					//Append to the commentList if no error occurs.
					commentList = append(commentList, c)
				}
				if comments.Err() != nil {
					return nil, err
				} else {
					return commentList, nil
				}
			}
		} else {
			//In the case you are not the profile owner, i.e. you result as "unauthorized", you must be Not Banned if you want to see the User photos.

			//Check first whether the user that is requesting the action has been banned by the fixedUsername.
			ban, errBan := db.CheckBan(fixedUsername, uuid)
			//Check for the error during the Query.
			if errBan != nil {
				return nil, errBan
			} else {

				//If no error occurs, checking whether the user was banned by the fixedUsername.
				if ban == "Not Banned" {
					//If you are not banned, select all the photo's comments.
					comments, err := db.c.Query(`SELECT * 
					FROM Comments 
					WHERE fixedUsername == '?' AND photoId == '?'
					ORDER BY uploadDate DESC`, fixedUsername, photoId)

					//Check for the error during the Query.
					if err != nil {
						return nil, err
					} else {
						//Defer the comments closure. This is a Best-Practice.
						defer func() { _ = comments.Close() }()

						// Here we read the resultset and we build the list to be returned.
						for comments.Next() {
							var c Comment
							err = comments.Scan(&c.Commentid, &c.PhotoId, &c.FixedUsername, &c.Phrase, &c.CommenterFixedUsername, &c.UploadDate)
							if err != nil {
								return nil, err
							}

							//Append to the commentList if no error occurs.
							commentList = append(commentList, c)
						}
						if comments.Err() != nil {
							return nil, err
						} else {
							return commentList, nil
						}
					}
				} else {
					//If the Use was Banned instead, returns nothing.
					fmt.Println("You cannot have the commentList you are requiring!")
					return nil, nil
				}
			}
		}
	}
}
