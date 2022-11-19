package api

import (
	"strings"
)

// Create a Fountain structure for having the Fountain struct.
// (Option + 9 on mac for putting the ` character). They are used for allowing to put the name as JSON RFC Standard Specifications.
type User struct {
	Uuid            string       `json:"uuid"`
	Username        string       `json:"username"`
	PersonalInfo    PersonalInfo `json:"personalInfo"`
	DateOfCreation  string       `json:"dateOfCreation"`
	NumberOfPhotos  int          `json:"numberOfPhotos"`
	TotNumberLikes  int          `json:"totNumberLikes"`
	NumberComments  int          `json:"numberComments"`
	NumberFollowers int          `json:"numberFollowers"`
	NumberFollowing int          `json:"numberFollowing"`
	ArrayPhotos     []Photo      `json:"arrayPhotos"`
}

type PersonalInfo struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	DateOfBirth string `json:"dateOfBirth"`
	Email       string `json:"email"`
	Nationality string `json:"nationality"`
	Gender      string `json:"gender"`
}

type Photo struct {
	Photoid        string   `json:"photoid"`
	Filename       string   `json:"filename"` //Note here it will be in bynary.
	UploadDate     string   `json:"uploadDate"`
	Location       Location `json:"location"`
	Phrase         string   `json:"phrase"`
	NumberLikes    int      `json:"numberLikes"`
	NumberComments int      `json:"numberComments"`
}

type Location struct {
	latitude  float64
	longitude float64
}

type Comment struct {
	Commentid         int    `json:"commentid"`
	CommenterUsername string `json:"commenterUsername"`
	Phrase            string `json:"phrase"`
	UploadDate        string `json:"uploadDate"`
}

type Like struct {
	Likeid     string `json:"likeid"` //This corresponds to the Username of the Liker.
	UploadDate string `json:"uploadDate"`
}

type Ban struct {
	Banid      string `json:"banid"` //This corresponds to the Username of the User Banned.
	UploadDate string `json:"uploadDate"`
	Motivation string `json:"motivation"`
}

type Follow struct {
	Followid   string `json:"followid"` //This corresponds to the Username of the Follower.
	UploadDate string `json:"uploadDate"`
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
