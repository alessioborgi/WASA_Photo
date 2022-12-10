package database

func (db *appdbimpl) UploadPhoto(photo Photo) (Photo, error) {

	// Adding a User Photo. Here we can distinguish two cases:
	//1) We have that the User can add the photo since the photo insertion action has been requested by the profile owner's coinciding with the action requester.
	//2) We have that the User cannot add the photo since the photo insertion action has NOT been requested by the profile owner's coinciding with the action requester.

	authorization, errAuth := db.CheckAuthorizationOwner(photo.FixedUsername, uuid)

	//Check for the error during the Query.
	if errAuth != nil {
		return photo, errAuth
	} else {
		// Go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).

		if authorization == "Authorized" {

			//If Authorized, you can proceed to add up the Photo Object without any problem.
			_, err := db.c.Exec(`INSERT INTO Photos (photoid, fixedUsername, filename, uploadDate, phrase, numberLikes, numberComments, latitude, longitude) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				photo.Photoid, photo.FixedUsername, photo.Filename, photo.UploadDate, photo.Phrase, photo.NumberLikes, photo.NumberComments, photo.Latitude, photo.Longitude)
			if err != nil {
				return photo, err
			} else {
				return photo, nil
			}

		} else {
			//If the User was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
			return photo, ErrUserNotAuthorized
		}
	}
}
