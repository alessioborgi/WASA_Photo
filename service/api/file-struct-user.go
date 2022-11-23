package api

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	DESCRIPTION:
	This file contains the user structure and its relative functions.
*/

// Declaration of Customized type used after for checking their validity.(User)
type Uuid string
type Username string

// Declaration of Customized type used after for checking validity (User Profile Personal Info).
type Name string
type Surname string
type Date string
type Email string
type Nationality string

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
	Name        Name    `json:"name"`
	Surname     Surname `json:"surname"`
	DateOfBirth Date  `json:"dateOfBirth"`
	Email       Email  `json:"email"`
	Nationality Nationality  `json:"nationality"`
	Gender      string  `json:"gender"`
}

// Declaring a Method for checking the uuid validty w.r.t. its length.
func (u Uuid) ValidUuid() bool {
	uuid := strings.SplitAfter(u, "-")
	return len(u) == 36 && len(uuid[0]) == 8 && len(uuid[1]) == 4 &&
		len(uuid[2]) == 4 && len(uuid[3]) == 4 && len(uuid[0]) == 12
}

// Declaring a Method for checking the Username validity w.r.t. its length.
func (u Username) ValidUsername() bool {
	return 1 <= len(u) && len(u) <= 31
}

// Declaring a Method for checking the Name validity w.r.t. its length.
func (u Name) ValidName() bool {
	return 1 <= len(u)
}

// Declaring a Method for checking the Surname validity w.r.t. its length.
func (s Surname) ValidSurname() bool {
	return 1 <= len(s)
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

// Function Method used to check for the User Validity.
func (f User) ValidUser() bool {
	return f.Uuid.ValidUuid() && f.Username.ValidUsername() && f.PersonalInfo.Name.ValidName() &&
		f.PersonalInfo.Surname.ValidSurname() && f.PersonalInfo.DateOfBirth.ValidDateofBirth() &&
		f.PersonalInfo.email.ValidEmail() && f.PersonalInfo.Nationality.ValidNationality() &&
		f.PersonalInfo.Gender.ValidGender()
}
