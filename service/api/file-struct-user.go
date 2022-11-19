package api

import (
	"strings"
)

/*
	DESCRIPTION:
	This file contains the user structure and its relative functions.
*/

// Declaration of Customized type used after for checking their validity.
type Uuid string
type Username string

// Create a User structure.
// (Option + 9 on mac for putting the ` character). They are used for allowing to put the name as JSON RFC Standard Specifications.
type User struct {
	Uuid            Uuid         `json:"uuid"`
	Username        Username     `json:"username"`
	PersonalInfo    PersonalInfo `json:"personalInfo"`
	DateOfCreation  string       `json:"dateOfCreation"`
	NumberOfPhotos  int          `json:"numberOfPhotos"`
	TotNumberLikes  int          `json:"totNumberLikes"`
	NumberComments  int          `json:"numberComments"`
	NumberFollowers int          `json:"numberFollowers"`
	NumberFollowing int          `json:"numberFollowing"`
	ArrayPhotos     []Photo      `json:"arrayPhotos"`
}

// Ceation of a sub-Structure that handles the Personal Information of the User.
type PersonalInfo struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	DateOfBirth string `json:"dateOfBirth"`
	Email       string `json:"email"`
	Nationality string `json:"nationality"`
	Gender      string `json:"gender"`
}

// Declaring a Method for checking the uuid validty w.r.t. its length.
func (u Uuid) ValidUuid() bool {
	uuid := strings.SplitAfter(u, "-")
	return len(u) == 36 && len(uuid[0]) == 8 && len(uuid[1]) == 4 &&
		len(uuid[2]) == 4 && len(uuid[3]) == 4 && len(uuid[0]) == 12
}

// Declaring a Method for checking the Username validity w.r.t. its length.
func (u Username) ValidUsername() bool {
	return 3 <= len(u) && len(u) <= 31
}

// Declaring a Method for checking the Gender validity w.r.t. it belongs to an "enum" of values.
func (f User) ValidGender() bool {
	g_lower := strings.ToLower(f.PersonalInfo.Gender)
	return g_lower == "male" || g_lower == "female" || g_lower == "do not specify"
}

// Function Method used to check for the User Validity.
func (f User) ValidUser() bool {
	return f.Uuid.ValidUuid() && f.Username.ValidUsername() && f.ValidGender()
}
