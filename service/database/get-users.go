package database

func (db *appdbimpl) GetUsers(uuid string) ([]string, error) {
	// Variable for returning the slice of fixedUsernames.
	var ret []string

	// Selection of the fixedUsernames available in WASAPhoto.
	//Check that the fixedUsername list should not include the fixedUsername of the person which is asking this action.
	//Check that the users that have banned the User that is requesting this action, must not be present in the list.
	rows, err := db.c.Query(`WITH Req_User (fixedUsername) AS ( 
		SELECT fixedUsername 
		FROM Users
		WHERE uuid == '123')
	SELECT fixedUsername FROM Users
	EXCEPT
	SELECT * FROM Req_User 
	EXCEPT
	SELECT fixedUsernameBanner
	FROM Bans
	WHERE fixedUsernameBanned == (SELECT * FROM Req_User)`)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	for rows.Next() {
		var f Fountain
		err = rows.Scan(&f.ID, &f.Latitude, &f.Longitude, &f.Status)
		if err != nil {
			return nil, err
		}

		ret = append(ret, f)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil
}
