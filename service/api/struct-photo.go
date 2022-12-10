package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type (
	DateTime   string
	Motivation string
)

var current_hour, current_minute, current_second = time.Now().Clock()

type Photo struct {
	Photoid        int      `json:"photoid"`
	Filename       byte     `json:"filename"` //Note here it will be in bynary.
	UploadDate     DateTime `json:"uploadDate"`
	Location       Location `json:"location"`
	Phrase         string   `json:"phrase"`
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

// ----- FINAL PHOTO FUNCTION -----

// Function Method used to check for the User Validity.
func ValidPhoto(photo Photo) bool {
	return photo.Photoid >= 0 &&
		photo.UploadDate.ValidUploadDate() &&
		photo.Location.ValidLocation() &&
		len(photo.Phrase) >= 5 && len(photo.Phrase) <= 1000 &&
		photo.NumberLikes >= 0 &&
		photo.NumberComments >= 0
}

// -----                -----
