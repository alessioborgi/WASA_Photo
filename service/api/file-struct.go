package api

import (
	"strings"
)

// Create a Fountain structure for having the Fountain struct.
// (Option + 9 on mac for putting the ` character). They are used for allowing to put the name as JSON RFC Standard Specifications.
type User struct {
	Uuid            string  `json:"uuid"`
	Username        float64 `json:"username"`
	PersonalInfo    float64 `json:"personalInfo"`
	DateOfCreation  string  `json:"dateOfCreation"`
	NumberOfPhotos  float64 `json:"numberOfPhotos"`
	TotNumberLikes  float64 `json:"totNumberLikes"`
	NumberComments  float64 `json:"numberComments"`
	NumberFollowers float64 `json:"numberFollowers"`
	NumberFollowing float64 `json:"numberFollowing"`
	ArrayPhotos     float64 `json:"arrayPhotos"`
}

// Function method used to check for the Fountain Validity.
func (f User) Valid() bool {
	f1 := strings.ToLower(f.Status)
	return f.ID > 0 && f.Latitude <= 90 && f.Latitude >= -90 && f.Longitude <= 180 && f.Longitude >= 180 &&
		(f1 == "faulty" || f1 == "good")
}

// Create a JSON Error Message Structure.
type JSONErrorMsg struct {
	Message string
}
