package api

import (
	"errors"
	"time"

	"github.com/alessioborgi/WASA_Photo/service/database"
)

type (
	DateTime   string
	Motivation string
)

var current_hour, current_minute, current_second = time.Now().Clock()

type Photo struct {
	Photoid        int64  `json:"photoid"`
	FixedUsername  string `json:"fixedUsername"`
	Username       string `json:"username"`
	Filename       string `json:"filename"`
	UploadDate     string `json:"uploadDate"`
	Phrase         string `json:"phrase"`
	NumberLikes    int64  `json:"numberLikes"`
	NumberComments int64  `json:"numberComments"`
}

// ----- FINAL PHOTO FUNCTION -----

// Function Method used to check for the User Validity.
func ValidPhoto(photo Photo) bool {
	return photo.Photoid >= 0 &&
		regex_username.MatchString(photo.Username) &&
		len(photo.Phrase) >= 5 && len(photo.Phrase) <= 1000 &&
		photo.NumberLikes >= 0 &&
		photo.NumberComments >= 0
}

var errphoto error

func (p *Photo) FromDatabase(photo database.Photo, db database.AppDatabase) error {
	p.Photoid = photo.Photoid
	p.Username, errphoto = db.GetUsername(photo.FixedUsername)
	if !errors.Is(errphoto, nil) {
		return errphoto
	}
	p.Filename = photo.Filename
	p.UploadDate = photo.UploadDate
	p.Phrase = photo.Phrase
	p.NumberLikes = photo.NumberLikes
	p.NumberComments = photo.NumberComments
	return nil
}

func (p *Photo) ToDatabase() database.Photo {
	return database.Photo{
		Photoid:        p.Photoid,
		FixedUsername:  p.FixedUsername,
		Filename:       p.Filename,
		UploadDate:     p.UploadDate,
		Phrase:         p.Phrase,
		NumberLikes:    p.NumberLikes,
		NumberComments: p.NumberComments,
	}
}
