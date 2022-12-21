package database

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func (db *appdbimpl) GetLastFixedUsername() (string, error) {

	// Variable Declaration
	var fixedUsernameList []int

	// Getting all the fixedUsername.
	users, err := db.c.Query(`SELECT fixedUsername FROM Users`)

	// Check if we have encountered some error in the retrieval of the query.
	if !errors.Is(err, nil) {

		// We check first whether the users retrieval caused an error.
		log.Println("Err: Error encountered during the Query in the DB.")
		return "", err
	}

	// If no error occur, we can proceed on elaborating the DB Response.
	log.Println("No error encountered during the Query in the DB.")
	defer func() { _ = users.Close() }()

	// Here we read the resultset and we build the list of fixedUsernames to be returned.
	var fixedUsername string

	for users.Next() {
		errFixedUsername := users.Scan(&fixedUsername)
		if !errors.Is(errFixedUsername, nil) {

			log.Println("Err: Error encountered during the scan.")
			return "", errFixedUsername
		}

		fixedUsername = strings.Replace(fixedUsername, "u", "", 1)
		intFixedUsername, errConvert := strconv.Atoi(fixedUsername)
		if !errors.Is(errConvert, nil) {

			log.Println("Err: Error encountered during the string-to-integer convertion.")
			return "", errConvert
		}

		// Add up to the UsernameList the fixedUsername.
		fixedUsernameList = append(fixedUsernameList, intFixedUsername)
	}

	// If we have encountered some error in the users variable.
	if users.Err() != nil {
		log.Println("Err: Error encountered on users")
		return "", err
	}

	max := fixedUsernameList[0]

	for i := 1; i < len(fixedUsernameList); i++ {
		if max < fixedUsernameList[i] {
			max = fixedUsernameList[i]
		}
	}

	finalFixedUsername := "u" + fmt.Sprint(max+1)
	log.Println("The new finalFixedUsername to be added is: ", finalFixedUsername)

	// We can assume that there is always at least one fixedUsername since there is the root User "alessioborgi01".
	return finalFixedUsername, nil

}
