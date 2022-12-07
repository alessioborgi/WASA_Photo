package database

func (db *appdbimpl) BanUser(ban Ban, uuid string) (Ban, error) {

	// Adding a User Ban. Here we can distinguish two cases:
	//1) We have that the User can add the ban since the ban action has been requested by the profile owner's coinciding with the action requester.
	//2) We have that the User cannot add the ban since the ban action has NOT been requested by the profile owner's coinciding with the action requester.

	authorization, errAuth := db.CheckAuthorizationOwner(ban.FixedUsernameBanner, uuid)

	//Check for the error during the Query.
	if errAuth != nil {
		return ban, errAuth
	} else {
		// Go checking whether you are authorized or not(i.e., whether you are the owner of the profile or not).

		if authorization == "Authorized" {

			//If Authorized, you can proceed to add up the Ban Object without any problem.
			_, err := db.c.Exec(`INSERT INTO Bans (fixedUsernameBanner, fixedUsernameBanned, uploadDate, Motivation) VALUES (?, ?, ?, ?)`,
				ban.FixedUsernameBanner, ban.FixedUsernameBanned, ban.UploadDate, ban.Motivation)
			if err != nil {
				return ban, err
			} else {
				return ban, nil
			}

		} else {
			//If the User was not "Authorized", i.e. it is not the Profile Owner, it must not be able to do this operation.
			return ban, ErrUserNotAuthorized
		}
	}
}
