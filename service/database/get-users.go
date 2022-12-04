package database

func (db *appdbimpl) GetUsers(uuid string) ([]string, error) {
	// Variable for returning the slice of fixedUsernames.
	var fixedUsernameList []string

	// Selection of the fixedUsernames available in WASAPhoto.
	//Check that the fixedUsername list should not include the fixedUsername of the person which is asking this action.
	//Check that the users that have banned the User that is requesting this action, must not be present in the list.
	//Returns the list in Descending Order(i.e., New Users first).
	users, err := db.c.Query(`WITH Req_User (fixedUsername) AS ( 
		SELECT fixedUsername 
		FROM Users
		WHERE uuid == '?'),
	Tuser (name) AS (
		SELECT fixedUsername FROM Users
		EXCEPT
		SELECT * FROM Req_User 
		EXCEPT
		SELECT fixedUsernameBanner
		FROM Bans
		WHERE fixedUsernameBanned == (SELECT * FROM Req_User))
	
	SELECT fixedUsername, dateOfCreation
	FROM Tuser AS t JOIN Users As u On t.name = u.fixedUsername
	ORDER By dateOfCreation`, uuid)

	if err != nil {
		return nil, err
	}
	defer func() { _ = users.Close() }()

	// Here we read the resultset and we build the list of fixedUsername to be returned
	for users.Next() {
		var f_username string
		err = users.Scan(&f_username)
		if err != nil {
			return nil, err
		}

		fixedUsernameList = append(fixedUsernameList, f_username)
	}
	if users.Err() != nil {
		return nil, err
	}

	return fixedUsernameList, nil
}
