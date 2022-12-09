package api

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/alessioborgi/WASA_Photo/service/database"
)

/*
	DESCRIPTION:
	This file contains the user structure and its relative functions.
*/

// Declaration of Customized type used after for checking their validity.(User)

// type Uuid string

// Declaration of Customized type used after for checking validity (User Profile Personal Info).
type Username string
type Date string
type Email string
type Gender string

// Variables Declaration.
var (
	regex_username = regexp.MustCompile(`^[a-zA-Z0-9._]{5,20}$`)
	// regex_uuid     = regexp.MustCompile(`^[0-9a-fA-F-]{36}`)                //123e4567-e89b-12d3-a456-426614174000
	// regex_date     = regexp.MustCompile(`^[0-9]{1,2}/[0-9]{1,2}/[0-9]{4}$`) //Without February Check.

	current_year, c_month, current_day = time.Now().Date()
	current_month                      = int(c_month)
)

// Create a User structure.
// (Option + 9 on mac for putting the ` character). They are used for allowing to put the name as JSON RFC Standard Specifications.
type User struct {
	FixedUsername     Username     `json:"fixedUsername"`
	Username          Username     `json:"username"`
	PhotoProfile      byte         `json:"photoProfile"`
	Biography         string       `json:"biography"`
	PersonalInfo      PersonalInfo `json:"personalInfo"`
	DateOfCreation    Date         `json:"dateOfCreation"`
	NumberOfPhotos    int64        `json:"numberOfPhotos"`
	TotNumberLikes    int64        `json:"totNumberLikes"`
	TotNumberComments int64        `json:"totNumberComments"`
	NumberFollowers   int64        `json:"numberFollowers"`
	NumberFollowing   int64        `json:"numberFollowing"`
}

// Ceation of a sub-Structure that handles the Personal Information of the User.
type PersonalInfo struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	DateOfBirth Date   `json:"dateOfBirth"`
	Email       Email  `json:"email"`
	Nationality string `json:"nationality"`
	Gender      Gender `json:"gender"`
}

// Declaring a Method for checking the Username validity w.r.t. its regex.
func (username Username) ValidUsername(regex regexp.Regexp) bool {
	return regex.MatchString(string(username))
}

// Declaring a Method for checking the DateOfBirth validity w.r.t. its validity.
func (d Date) ValidDateofBirth() bool { //yyyy/mm/dd
	date := strings.Split(string(d), "-") //here find the way to also include "/"
	year, erry := strconv.Atoi(date[0])
	month, errm := strconv.Atoi(date[1])
	day, errd := strconv.Atoi(date[2])

	if erry != nil && errm != nil && errd != nil {
		//Here assume that the user should have been born more than 10 year ago.
		if len(d) == 10 && year >= 1900 && year <= current_year-10 && month >= 1 && month <= 12 && day >= 1 {
			if month == 2 && (day != 29 && day != 30 && day != 31) {
				if (month == 4 || month == 6 || month == 9 || month == 11) && (day != 31) {
					fmt.Println("Correct Date of Birth Inserted", d)
					return true
				} else {
					fmt.Println("InCorrect Date of Birth Inserted", d)
					return false
				}
			} else {
				fmt.Println("InCorrect Date of Birth Inserted", d)
				return false
			}
		} else {
			fmt.Println("InCorrect Date of Birth Inserted", d)
			return false
		}
	} else {
		fmt.Println("Error encountered!")
		return false
	}
}

// Declaring a Method for checking the Email validity w.r.t. its length.
func (e Email) ValidEmail() bool {
	email_at := strings.Split(string(e), "-")
	pre_at := email_at[0]
	post_at := email_at[1]
	if len(pre_at) >= 1 && len(post_at) >= 1 && len(e) <= 100 {
		fmt.Println("Valid Email Inserted!")
		return true
	} else {
		fmt.Println("Not Valid Email Inserted!")
		return false
	}
}

// Declaring a Method for checking the Gender validity w.r.t. it belongs to an "enum" of values.
func (g Gender) ValidGender() bool {
	g_lower := strings.ToLower(string(g))
	return g_lower == "male" || g_lower == "female" || g_lower == "do not specify"
}

// Declaring a Method for checking the DateOfCreation validity w.r.t. its validity.
func (d Date) ValidDateofCreation() bool { //yyyy/mm/dd
	date := strings.Split(string(d), "-") //here find the way to also include "/"
	year, erry := strconv.Atoi(date[0])
	month, errm := strconv.Atoi(date[1])
	day, errd := strconv.Atoi(date[2])

	if erry != nil && errm != nil && errd != nil {
		// Check whether the date is after the 2022-01-01 and before the current date (range included).
		if len(d) == 10 && year <= 2022 && year <= current_year && month >= 1 && month <= 12 && month <= current_month && day >= 1 && current_day <= 31 && day <= current_day {
			if month == 2 && (day != 29 && day != 30 && day != 31) {
				if (month == 4 || month == 6 || month == 9 || month == 11) && (day != 31) {
					fmt.Println("Correct Date of Creation Inserted", d)
					return true
				} else {
					fmt.Println("InCorrect Date of Creation Inserted", d)
					return false
				}

			} else {
				fmt.Println("InCorrect Date of Creation Inserted", d)
				return false
			}
		} else {
			fmt.Println("InCorrect Date of Creation Inserted", d)
			return false
		}
	} else {
		fmt.Println("Error encountered!")
		return false
	}
}

// ----- FINAL FUNCTION -----

// Function Method used to check for the User Validity.
func (user *User) ValidUser() bool {
	return user.FixedUsername.ValidUsername(*regex_username) && //Checking for the FixedUsername Regex.
		user.Username.ValidUsername(*regex_username) && //Checking for the Username Regex.
		len(user.Biography) >= 0 && len(user.Biography) <= 1000 && //Checking for the Biography.
		len(user.PersonalInfo.Name) >= 2 && len(user.PersonalInfo.Name) <= 31 && //Checking for the Name.
		len(user.PersonalInfo.Surname) >= 2 && len(user.PersonalInfo.Surname) <= 31 && //Checking for the Surname.
		user.PersonalInfo.DateOfBirth.ValidDateofBirth() &&
		user.PersonalInfo.Email.ValidEmail() &&
		len(user.PersonalInfo.Nationality) >= 3 && len(user.PersonalInfo.Nationality) <= 100 &&
		user.PersonalInfo.Gender.ValidGender() &&
		user.DateOfCreation.ValidDateofCreation() &&
		user.NumberOfPhotos >= 0 &&
		user.TotNumberLikes >= 0 &&
		user.TotNumberComments >= 0 &&
		user.NumberFollowers >= 0 &&
		user.NumberFollowing >= 0
}

// -----                -----

// Function for handling the population of the struct with data from the DB.
func (u *User) FromDatabase(user database.User) {
	u.FixedUsername = Username(user.FixedUsername)
	u.Username = Username(user.Username)
	u.PhotoProfile = user.PhotoProfile
	u.Biography = user.Biography
	u.DateOfCreation = Date(user.DateOfCreation)
	u.NumberOfPhotos = user.NumberOfPhotos
	u.TotNumberLikes = user.TotNumberLikes
	u.TotNumberComments = user.TotNumberComments
	u.NumberFollowers = user.NumberFollowers
	u.NumberFollowing = user.NumberFollowing

	//Also personalInfo Struct
	u.PersonalInfo.Name = user.Name
	u.PersonalInfo.Surname = user.Surname
	u.PersonalInfo.DateOfBirth = Date(user.DateOfBirth)
	u.PersonalInfo.Email = Email(user.Email)
	u.PersonalInfo.Nationality = user.Nationality
	u.PersonalInfo.Gender = Gender(user.Gender)
}

// ToDatabase returns the User in a Database-Compatible Representation.

// DOUBT: What to do with the UUID?
func (u *User) ToDatabase() database.User {
	return database.User{
		FixedUsername:     string(u.FixedUsername),
		Username:          string(u.Username),
		PhotoProfile:      u.PhotoProfile, //Maybe without byte()?
		Biography:         u.Biography,
		DateOfCreation:    string(u.DateOfCreation),
		NumberOfPhotos:    u.NumberOfPhotos,
		TotNumberLikes:    u.TotNumberLikes,
		TotNumberComments: u.TotNumberComments,
		NumberFollowers:   u.NumberFollowers,
		NumberFollowing:   u.NumberFollowing,
		Name:              u.PersonalInfo.Name,
		Surname:           u.PersonalInfo.Surname,
		DateOfBirth:       string(u.PersonalInfo.DateOfBirth),
		Email:             string(u.PersonalInfo.Email),
		Nationality:       u.PersonalInfo.Nationality,
		Gender:            string(u.PersonalInfo.Gender),
	}
}
