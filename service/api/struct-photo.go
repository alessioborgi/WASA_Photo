package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Phrase string
type DateTime string
type Motivation string

var current_hour, current_minute, current_second = time.Now().Clock()

type Photo struct {
	Photoid        int      `json:"photoid"`
	Filename       byte     `json:"filename"` //Note here it will be in bynary.
	UploadDate     DateTime `json:"uploadDate"`
	Location       Location `json:"location"`
	Phrase         Phrase   `json:"phrase"`
	NumberLikes    int      `json:"numberLikes"`
	NumberComments int      `json:"numberComments"`
}

// Declaring a Method for checking the UploadDate validity w.r.t. its validity.
func (d DateTime) ValidUploadDate() bool { //yyyy/mm/dd
	date_time := strings.Split(string(d), "T") //here find the way to also include "/"
	date := strings.Split(string(date_time[0]), "-")
	time := strings.Split(string(date_time[1]), ":")
	year, erry := strconv.Atoi(date[0])
	month, errm := strconv.Atoi(date[1])
	day, errd := strconv.Atoi(date[2])
	hour, errh := strconv.Atoi(time[0])
	minute, errmi := strconv.Atoi(time[1])
	second, errs := strconv.Atoi(time[2][:2]) //For not taking the "Z" character.

	if erry != nil && errm != nil && errd != nil && errh != nil && errmi != nil && errs != nil {
		// Check whether the date is after the 2022-01-01T00:00:00Z and before the current date (range included).
		if len(d) == 20 && year <= 2022 && year <= current_year && month >= 1 && month <= 12 && month <= current_month && day >= 1 && current_day <= 31 && day <= current_day && hour >= 0 && hour <= 24 && hour <= current_hour && minute >= 0 && minute <= 60 && minute <= second && second >= 0 && second <= 60 && second <= current_second {
			if month == 2 && (day != 29 && day != 30 && day != 31) {
				if (month == 4 || month == 6 || month == 9 || month == 11) && (day != 31) {
					fmt.Println("Correct Photo's UploadDate Inserted", d)
					return true
				} else {
					fmt.Println("InCorrect Photo's UploadDate Inserted", d)
					return false
				}

			} else {
				fmt.Println("InCorrect  Photo's UploadDate Inserted", d)
				return false
			}
		} else {
			fmt.Println("InCorrect  Photo's UploadDate Inserted", d)
			return false
		}
	} else {
		fmt.Println("Error encountered!")
		return false
	}
}

type Location struct {
	latitude  float64
	longitude float64
}

func (l Location) ValidLocation() bool {
	return l.latitude <= 90 && l.latitude >= -90 && l.longitude <= 180 && l.longitude >= -180
}

func (p Phrase) ValidPhrase() bool {
	return len(p) >= 5 && len(p) <= 1000
}

// ----- FINAL PHOTO FUNCTION -----

// Function Method used to check for the User Validity.
func ValidPhoto(photo Photo) bool {
	return photo.Photoid >= 0 &&
		photo.UploadDate.ValidUploadDate() &&
		photo.Location.ValidLocation() &&
		photo.Phrase.ValidPhrase() &&
		photo.NumberLikes >= 0 &&
		photo.NumberComments >= 0
}

// -----                -----

type Comment struct {
	Commentid         int           `json:"commentid"`
	CommenterUsername FixedUsername `json:"commenterUsername"`
	Phrase            Phrase        `json:"phrase"`
	UploadDate        DateTime      `json:"uploadDate"`
}

// ----- FINAL PHOTO FUNCTION -----

// Function Method used to check for the User Validity.
func ValidComment(comment Comment) bool {
	return comment.Commentid >= 0 &&
		comment.CommenterUsername.ValidFixedUsername(regex_fixed_username) &&
		comment.Phrase.ValidPhrase() &&
		comment.UploadDate.ValidUploadDate()
}

// -----                -----

type Like struct {
	Likeid     FixedUsername `json:"likeid"` //This corresponds to the Username of the Liker.
	UploadDate DateTime      `json:"uploadDate"`
}

// ----- FINAL LIKE FUNCTION -----

// Function Method used to check for the User Validity.
func ValidLike(like Like) bool {
	return like.Likeid.ValidFixedUsername(regex_fixed_username) &&
		like.UploadDate.ValidUploadDate()
}

// -----                -----

type Ban struct {
	Banid      FixedUsername `json:"banid"` //This corresponds to the Username of the User Banned.
	UploadDate DateTime      `json:"uploadDate"`
	Motivation Motivation    `json:"motivation"`
}

func (m Motivation) ValidMotivation() bool {
	m_lower := strings.ToLower(string(m))
	return m_lower == "spam" || m_lower == "bad behaviour" || m_lower == "threats"
}

// ----- FINAL BAN FUNCTION -----

// Function Method used to check for the User Validity.
func ValidBan(ban Ban) bool {
	return ban.Banid.ValidFixedUsername(regex_fixed_username) &&
		ban.UploadDate.ValidUploadDate() &&
		ban.Motivation.ValidMotivation()
}

// -----                -----

type Follow struct {
	Followid   FixedUsername `json:"followid"` //This corresponds to the Username of the Follower.
	UploadDate DateTime      `json:"uploadDate"`
}

// ----- FINAL FOLLOW FUNCTION -----

// Function Method used to check for the User Validity.
func ValidFollow(follow Follow) bool {
	return follow.Followid.ValidFixedUsername(regex_fixed_username) &&
		follow.UploadDate.ValidUploadDate()
}

// -----                -----

// Create a JSON Error Message Structure.
type JSONErrorMsg struct {
	Message string
}
