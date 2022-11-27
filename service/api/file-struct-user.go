package api

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"regexp"
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

// Variables Declaration.
var regex_fixed_username string = "[a-zA-Z0-9]+$"
var regex_uuid string = "^[0-9a-fA-F-]{36}"		//123e4567-e89b-12d3-a456-426614174000
var regex_date = "^[0-9]{1,2}/[0-9]{1,2}/[0-9]{4}$"	//Without February Check.



// Create a User structure.
// (Option + 9 on mac for putting the ` character). They are used for allowing to put the name as JSON RFC Standard Specifications.
type User struct {
	Uuid            Uuid         `json:"uuid"`
	FixedUsername   FixedUsername`json:"fixedUsername"`
	Username        Username     `json:"username"`
	PhotoProfile	byte		 `json:"photoProfile"`
	PersonalInfo    PersonalInfo `json:"personalInfo"`
	DateOfCreation  string       `json:"dateOfCreation"`
	NumberOfPhotos  int          `json:"numberOfPhotos"`
	TotNumberLikes  int          `json:"totNumberLikes"`
	NumberComments  int          `json:"numberComments"`
	NumberFollowers int          `json:"numberFollowers"`
	NumberFollowing int          `json:"numberFollowing"`
}

// Ceation of a sub-Structure that handles the Personal Information of the User.
type PersonalInfo struct {
	Name        Name    `json:"name"`
	Surname     Surname `json:"surname"`
	DateOfBirth Date  `json:"dateOfBirth"`
	Email       Email  `json:"email"`
	Nationality Nationality  `json:"nationality"`
	Gender      string  `json:"gender"`
}

// Declaring a Method for checking the uuid validty w.r.t. its Regex.
func (u Uuid) ValidUuid(regex string) bool {
	match, err := regexp.MatchString(regex, string(u))
	if err == nil {
		correct_spaces := string(u[8]) == "-" && string(u[13]) == "-" && string(u[18]) == "-" && string(u[23]) == "-"
		if match == true && correct_spaces == true{
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
		if err == true && len(string(fu)) >= 3 && len(string(fu)) <= 31{
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
func (d Date) ValidDateofBirth() bool {
	date := strings.SplitAfter(d, "-")
	year, erry := strconv.ParseInt(date[0], 10, 32)
	month, errm := strconv.ParseInt(date[1], 10, 32)
	day, errd := strconv.ParseInt(date[1], 10, 32)
	if erry != nil && errm != nil && errd != nil {
		return year >= 1900 && year <= currentyear &&
			month >= 1 && month <= currentmonth &&
			day >= 1 && day <= currentday
	}else if erry == nil && errm != nil && errd != nil {
		fmt.Println("Error on Year of the Inserted Date")
		return
	}else if erry != nil && errm == nil && errd != nil {
		fmt.Println("Error on Month of the Inserted Date")
		return
	}else if erry != nil && errm != nil && errd == nil {
		fmt.Println("Error on Day of the Inserted Date")
		return
	}else{
		fmt.Println("Error on more part of the Date Inserted!")
		return
	}

// ----- FINAL FUNCTION -----

//Function Method used to check for the User Validity.
func ValidUser(user User, regex string) bool {
	return user.Uuid.ValidUuid(regex_uuid) && 
	user.FixedUsername.ValidFixedUsername(regex_fixed_username) &&
	user.Username.ValidUsername() &&
	user.PersonalInfo.Name.ValidName() &&
	user.PersonalInfo.Surname.ValidSurname()
}


















// Declaring a Method for checking the Email validity w.r.t. its length.
func (e Email) ValidEmail() bool {
	email_at := strings.SplitAfter(d, "-")
	pre_at, err_pre := strconv.ParseInt(date[0], 10, 32)
	post_at, err_post := strconv.ParseInt(date[1], 10, 32)
	if err_pre != nil && err_post != nil {
		return 1 <= len(e) && 1 <= len(pre_at) && 1 <= len(post_at)
	}else if err_pre == nil && err_post != nil {
		fmt.Println("Empty part before the 'at' Character! Wrong Email inserted!") 
		return
	}else if err_pre != nil && err_post == nil {
		fmt.Println("Empty part after the 'at' Character! Wrong Email inserted!") 
		return
	}else if err_pre == nil && err_post != nil {
		fmt.Println("Empty Email! Wrong Email inserted!") 
		return
	}
}

// Declaring a Method for checking the Nationality validity w.r.t. its length.
func (n Nationality) ValidNationality() bool {
	return 1 <= len(n)
}

 
// Declaring a Method for checking the Gender validity w.r.t. it belongs to an "enum" of values.
func (f User) ValidGender() bool {
	g_lower := strings.ToLower(f.PersonalInfo.Gender)
	return g_lower == "male" || g_lower == "female" || g_lower == "do not specify"
}

