package api

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/alessioborgi/WASA_Photo/service/database"
)

/*
	DESCRIPTION:
	This file contains the user structure and its relative functions.
*/

// Declaration of Customized type used after for checking their validity.(User)

// Declaration of Customized type used after for checking validity (User Profile Personal Info).
type Date string
type Email string
type Gender string
type Uuid string

// Variables Declaration.
var (
	regex_username      = regexp.MustCompile(`^[a-zA-Z0-9._]{5,20}$`)
	regex_fixedUsername = regexp.MustCompile(`^[u0-9]{2,31}$`)
	regex_uuid          = regexp.MustCompile(`^[0-9a-fA-F-]{36}`)
)

// Create a User structure.
// (Option + 9 on mac for putting the ` character). They are used for allowing to put the name as JSON RFC Standard Specifications.
type User struct {
	FixedUsername   string `json:"fixedUsername"`
	Username        string `json:"username"`
	PhotoProfile    string `json:"photoProfile"`
	Biography       string `json:"biography"`
	DateOfCreation  Date   `json:"dateOfCreation"`
	NumberOfPhotos  int64  `json:"numberOfPhotos"`
	NumberFollowers int64  `json:"numberFollowers"`
	NumberFollowing int64  `json:"numberFollowing"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	DateOfBirth     Date   `json:"dateOfBirth"`
	Email           Email  `json:"email"`
	Nationality     string `json:"nationality"`
	Gender          Gender `json:"gender"`
}

type Username struct {
	Name string `json:"username"`
}

// Declaring a Method for checking the Username validity w.r.t. its regex.
func (username Username) ValidUsername(regex regexp.Regexp) bool {
	return regex.MatchString(username.Name)
}

// Declaring a Method for checking the uuid validity w.r.t. its regex.
func (uuid Uuid) ValidUuid(regex regexp.Regexp) bool {
	return regex.MatchString(string(uuid))
}

// Declaring a Method for checking the DateOfBirth validity w.r.t. its validity.
// yyyy/mm/dd
func (d Date) ValidDateofBirth() bool {
	date := strings.Split(string(d), "-")
	year, erry := strconv.Atoi(date[0])
	month, errm := strconv.Atoi(date[1])
	day, errd := strconv.Atoi(date[2])
	log.Println("The date of Birth is: ", year, month, day)

	if !errors.Is(erry, nil) || !errors.Is(errm, nil) || !errors.Is(errd, nil) {
		log.Println("Error encountered in the Date of Birth!")
	}
	log.Println("Correct Date of Birth Inserted", d)

	return false
}

// Declaring a Method for checking the Email validity w.r.t. its length.
func (e Email) ValidEmail() bool {
	email_at := strings.Split(string(e), "@")
	pre_at := email_at[0]
	post_at := email_at[1]
	if len(pre_at) >= 1 && len(post_at) >= 1 && len(e) <= 100 {
		log.Println("Valid Email Inserted!")
		return true
	} else {
		log.Println("Not Valid Email Inserted!")
		return false
	}
}

// Declaring a Method for checking the Gender validity w.r.t. it belongs to an "enum" of values.
func (g Gender) ValidGender() bool {
	g_lower := strings.ToLower(string(g))
	return g_lower == "male" || g_lower == "female" || g_lower == "do not specify"
}

// ----- FINAL FUNCTION -----

// Function Method used to check for the User Validity.
func (user *User) ValidUser() bool {
	return regex_username.MatchString(user.Username) && // Checking for the Username Regex.
		len(user.Biography) <= 1000 && // Checking for the Biography.
		len(user.Name) >= 2 && len(user.Name) <= 31 && // Checking for the Name.
		len(user.Surname) >= 2 && len(user.Surname) <= 31 && // Checking for the Surname.
		// user.DateOfBirth.ValidDateofBirth() &&
		user.Email.ValidEmail() &&
		len(user.Nationality) >= 3 && len(user.Nationality) <= 100 &&
		user.Gender.ValidGender() &&
		user.NumberOfPhotos >= 0 &&
		user.NumberFollowers >= 0 &&
		user.NumberFollowing >= 0
}

// -----                -----

// Function for handling the population of the struct with data from the DB.
func (u *User) FromDatabase(user database.User, db database.AppDatabase) {
	u.FixedUsername = user.FixedUsername
	u.Username = user.Username
	u.PhotoProfile = user.PhotoProfile
	u.Biography = user.Biography
	u.DateOfCreation = Date(user.DateOfCreation)
	u.NumberOfPhotos = user.NumberOfPhotos
	u.NumberFollowers = user.NumberFollowers
	u.NumberFollowing = user.NumberFollowing

	// Also personalInfo Struct
	u.Name = user.Name
	u.Surname = user.Surname
	u.DateOfBirth = Date(user.DateOfBirth)
	u.Email = Email(user.Email)
	u.Nationality = user.Nationality
	u.Gender = Gender(user.Gender)
}

// ToDatabase returns the User in a Database-Compatible Representation.

func (u *User) ToDatabase() database.User {
	return database.User{
		FixedUsername:   u.FixedUsername,
		Username:        u.Username,
		Biography:       u.Biography,
		PhotoProfile:    u.PhotoProfile,
		DateOfCreation:  string(u.DateOfCreation),
		NumberOfPhotos:  u.NumberOfPhotos,
		NumberFollowers: u.NumberFollowers,
		NumberFollowing: u.NumberFollowing,
		Name:            u.Name,
		Surname:         u.Surname,
		DateOfBirth:     string(u.DateOfBirth),
		Email:           string(u.Email),
		Nationality:     u.Nationality,
		Gender:          string(u.Gender),
	}
}
