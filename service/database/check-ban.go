package database

import "errors"

func (db *appdbimpl) CheckBan(fixedUsername string, uuid string) (string, error) {

	// Check whether the uuid that is requesting the action has been banned by the fixedUsername.
	var ban bool
	err := db.c.QueryRow(`SELECT fixedUsernameBanned == (SELECT fixedUsername FROM Users WHERE uuid == '?')
	FROM Bans
	WHERE fixedUsernameBanner == '?'`, uuid, fixedUsername).Scan(&ban)

	// Check for the error during the Query.
	if !errors.Is(err, nil) {
		return "", err
	} else {
		// Returning "Banned" if the Uuid has been banned by fixedUsername, "Not Banned" otherwise.
		if ban == true {
			return "Banned", nil
		} else {
			return "Not Banned", nil
		}
	}
}
