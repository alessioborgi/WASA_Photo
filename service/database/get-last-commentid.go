package database

import (
	"log"
)

func (db *appdbimpl) GetLastCommentId() (int, error) {

	// Variable Declaration
	var commentsList []int

	// Getting all the fixedUsername.
	comments, err := db.c.Query(`SELECT commentid FROM Comments`)

	// Check if we have encountered some error in the retrieval of the query.
	if err != nil {

		// We check first whether the comments retrieval caused an error.
		log.Println("Err: Error encountered during the Query in the DB.")
		return 0, err
	}

	// If no error occur, we can proceed on elaborating the DB Response.
	log.Println("No error encountered during the Query in the DB.")
	defer func() { _ = comments.Close() }()

	// Here we read the resultset and we build the list of fixedUsernames to be returned.
	var comment int

	for comments.Next() {
		errCom := comments.Scan(&comment)
		if errCom != nil {

			log.Println("Err: Error encountered during the scan.")
			return 0, errCom
		}

		// Add up to the UsernameList the fixedUsername.
		commentsList = append(commentsList, comment)
	}

	// Check whether the list is empty.
	if len(commentsList) == 0 {
		return 1, nil
	}

	// Get the maximum value.
	max := commentsList[0]

	for i := 1; i < len(commentsList); i++ {
		if max < commentsList[i] {
			max = commentsList[i]
		}
	}

	return max + 1, nil
}
