package database

func (db *appdbimpl) CheckAuthorizationOwner(fixedUsername string, uuid string) (string, error) {

	//Check whether the uuid that is requesting the action is the owner of the profile.
	var auth bool
	err := db.c.QueryRow(`SELECT uuid == '?' 
	FROM Users 
	WHERE fixedUsername == '?'`, uuid, fixedUsername).Scan(&auth)

	//Check for the error during the Query.
	if err != nil {
		return "", err
	} else {
		//Returning "Authorized" if it is the owner, "Not Authorized" otherwise.
		if auth == true {
			return "Authorized", nil
		} else {
			return "Not Authorized", nil
		}
	}
}

// (`SELECT iif(uuid == '?', "Authorized", "Not Authorized" )
// 	FROM Users
// 	WHERE fixedUsername == '?'`, uuid, fixedUsername)
