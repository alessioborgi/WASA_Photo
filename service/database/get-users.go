package database

import (
	"log"
)

func (db *appdbimpl) GetUsers(uuid string) ([]string, error) {

	// Variable for returning the slice of usernames.
	var usernameList []string

	// Selection of the usernameList available in WASAPhoto.

	// Check that the fixedUsername list should not include the fixedUsername of the person which is asking this action.
	// Check that the users that have banned the User that is requesting this action, must not be present in the list.
	// Returns the list in Descending Order(i.e., New Users first).

	// First of all, check whether we have the Uuid Authorization in the DB to do this.
	// Note that, here, we don't need to check the authorization owner since we are simply doing a get.
	var presence_auth string
	presence_auth, err := db.CheckAuthorizationPresence(uuid)
	if err != nil {

		// We check first whether the users retrieval caused an error.
		log.Fatalf("Error encountered during the Check of the Uuid Presence.")
		return nil, err
	} else if presence_auth == PRESENT {

		// If we arrive here, we have that the Uuid is present in the DB. We can therefore check for the list of users he/she can see.
		// 1) Select all the Usernames in the WASAPhoto Platform except the Username that is actually requesting this action.

		// TO DO:
		// 2) From this, also remove those users that have banned the uuid that is requesting the action.
		users, err := db.c.Query(`SELECT username FROM Users WHERE username NOT IN 
				(SELECT username FROM Users WHERE uuid == ?)
			ORDER BY dateOfCreation DESC`, uuid)

		// users, err := db.c.Query(`WITH Req_User (fusername) AS (SELECT fixedUsername FROM Users WHERE uuid == ?),
		// Tuser (tfusername) AS ( SELECT fixedUsername FROM Users EXCEPT SELECT * FROM Req_User),
		// Buser (bfusername) AS (SELECT b.fixedUsernameBanner from Users AS u JOIn Bans AS b On u.fixedUsername = b.fixedUsernameBanner WHERE b.fixedUsernameBanned == (SELECT * FROM Req_User)),
		// Suser (sfusername) AS ( SELECT tfusername FROM Tuser EXCEPT SELECT bfusername FROM Buser)

		// SELECT u.username
		// FROM Suser as s JOIN Users as u On s.sfusername = u.fixedUsername
		// ORDER By u.dateOfCreation DESC`, uuid)

		if err != nil {

			// We check first whether the users retrieval caused an error.
			log.Fatalf("Error encountered during the Query in the DB.")
			return nil, err
		}

		// If no error occur, we can proceed on elaborating the DB Response.
		log.Println("No error encountered during the Query in the DB.")
		defer func() { _ = users.Close() }()

		// Here we read the resultset and we build the list of Usernames to be returned.
		var username string
		for users.Next() {
			err = users.Scan(&username)
			if err != nil {
				log.Fatalf("Error encountered during the scan.")
				return nil, err
			}

			// Add up to the UsernameList the Username.
			usernameList = append(usernameList, username)
		}

		if users.Err() != nil {
			log.Fatalf("Error encountered on Users")
			return nil, err
		} else {

			// First, check whether the returned list has length equal to zero. Return No Content if so.
			if len(usernameList) == 0 {
				return usernameList, ErrNoContent
			}

			// If we arrive here, we have a result with at least one Username.
			return usernameList, nil
		}

	} else {

		// If we arrive here, it means that the Uuid that was inserted was not valid since there is no user with this Uuid.
		return nil, ErrUserNotAuthorized
	}
}
