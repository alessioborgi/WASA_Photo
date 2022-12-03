package database

func (db *appdbimpl) GetUsers(uuid string) ([]string, error) {
	// Variable for returning the slice of fixedUsernames.
	var fixedUsernameList []string

	// Selection of the fixedUsernames available in WASAPhoto.
	//Check that the fixedUsername list should not include the fixedUsername of the person which is asking this action.
	//Check that the users that have banned the User that is requesting this action, must not be present in the list.
	rows, err := db.c.Query(`WITH Req_User (fixedUsername) AS ( 
		SELECT fixedUsername 
		FROM Users
		WHERE uuid == '?')
	SELECT fixedUsername FROM Users
	EXCEPT
	SELECT * FROM Req_User 
	EXCEPT
	SELECT fixedUsernameBanner
	FROM Bans
	WHERE fixedUsernameBanned == (SELECT * FROM Req_User)`, uuid)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list of fixedUsername to be returned
	for rows.Next() {
		var f_username string
		err = rows.Scan(&f_username)
		if err != nil {
			return nil, err
		}

		fixedUsernameList = append(fixedUsernameList, f_username)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return fixedUsernameList, nil
}
