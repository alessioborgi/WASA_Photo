package api

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
	DESCRIPTION:
	This file contains the user structure and its relative functions.
*/

// Declaration of Customized type used after for checking their validity.(User)
type Uuid string
type Username string
type FixedUsername string

// Declaration of Customized type used after for checking validity (User Profile Personal Info).
type Name string
type Surname string
type Date string
type Email string
type Nationality string
type Gender string

// Variables Declaration.
var regex_fixed_username string = "[a-zA-Z0-9]+$"
var regex_uuid string = "^[0-9a-fA-F-]{36}"         //123e4567-e89b-12d3-a456-426614174000
var regex_date = "^[0-9]{1,2}/[0-9]{1,2}/[0-9]{4}$" //Without February Check.

var current_year, c_month, current_day = time.Now().Date()
var current_month = int(c_month)

// Create a User structure.
// (Option + 9 on mac for putting the ` character). They are used for allowing to put the name as JSON RFC Standard Specifications.
type User struct {
	Uuid              Uuid          `json:"uuid"`
	FixedUsername     FixedUsername `json:"fixedUsername"`
	Username          Username      `json:"username"`
	PhotoProfile      byte          `json:"photoProfile"`
	PersonalInfo      PersonalInfo  `json:"personalInfo"`
	DateOfCreation    Date          `json:"dateOfCreation"`
	NumberOfPhotos    int           `json:"numberOfPhotos"`
	TotNumberLikes    int           `json:"totNumberLikes"`
	TotNumberComments int           `json:"totNumberComments"`
	NumberFollowers   int           `json:"numberFollowers"`
	NumberFollowing   int           `json:"numberFollowing"`
}

// Ceation of a sub-Structure that handles the Personal Information of the User.
type PersonalInfo struct {
	Name        Name        `json:"name"`
	Surname     Surname     `json:"surname"`
	DateOfBirth Date        `json:"dateOfBirth"`
	Email       Email       `json:"email"`
	Nationality Nationality `json:"nationality"`
	Gender      Gender      `json:"gender"`
}

// Declaring a Method for checking the uuid validty w.r.t. its Regex.
func (u Uuid) ValidUuid(regex string) bool {
	match, err := regexp.MatchString(regex, string(u))
	if err == nil {
		correct_spaces := string(u[8]) == "-" && string(u[13]) == "-" && string(u[18]) == "-" && string(u[23]) == "-"
		if match == true && correct_spaces == true {
			fmt.Println("Uuid Regex Matched")
			return true
		} else {
			fmt.Println("Uuid Regex UnMatched!")
			return false
		}
	} else {
		fmt.Println("Error:", err)
		return false
	}
}

// Declaring a Method for checking the fixedUsername validty w.r.t. its Regex.
func (fu FixedUsername) ValidFixedUsername(regex string) bool {
	match, err := regexp.MatchString(regex, string(fu))
	if err == nil {
		if match == true && len(string(fu)) >= 3 && len(string(fu)) <= 31 {
			fmt.Println("Fixed Username Regex Matched")
			return true
		} else {
			fmt.Println("Fixed Username Regex UnMatched!")
			return false
		}
	} else {
		fmt.Println("Error:", err)
		return false
	}
}

// Declaring a Method for checking the Username validity w.r.t. its length.
func (u Username) ValidUsername() bool {
	return len(string(u)) >= 3 && len(string(u)) <= 31
}

// Declaring a Method for checking the Name validity w.r.t. its length.
func (n Name) ValidName() bool {
	return len(string(n)) >= 2 && len(string(n)) <= 31
}

// Declaring a Method for checking the Surname validity w.r.t. its length.
func (s Surname) ValidSurname() bool {
	return len(string(s)) >= 2 && len(string(s)) <= 31
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

// Declaring a Method for checking the Nationality validity w.r.t. its length.
func (n Nationality) ValidNationality() bool {
	return len(n) >= 3 && len(n) <= 100
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
func ValidUser(user User, regex string) bool {
	return user.Uuid.ValidUuid(regex_uuid) &&
		user.FixedUsername.ValidFixedUsername(regex_fixed_username) &&
		user.Username.ValidUsername() &&
		user.PersonalInfo.Name.ValidName() &&
		user.PersonalInfo.Surname.ValidSurname() &&
		user.PersonalInfo.DateOfBirth.ValidDateofBirth() &&
		user.PersonalInfo.Email.ValidEmail() &&
		user.PersonalInfo.Nationality.ValidNationality() &&
		user.PersonalInfo.Gender.ValidGender() &&
		user.DateOfCreation.ValidDateofCreation() &&
		user.NumberOfPhotos >= 0 &&
		user.TotNumberLikes >= 0 &&
		user.TotNumberComments >= 0 &&
		user.NumberFollowers >= 0 &&
		user.NumberFollowing >= 0
}

// -----                -----
