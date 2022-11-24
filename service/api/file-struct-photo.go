package api

import (
	"strings"
)

type Phrase string
type Photo struct {
	Photoid        int      `json:"photoid"`
	Filename       byte     `json:"filename"` //Note here it will be in bynary.
	UploadDate     string   `json:"uploadDate"`
	Location       Location `json:"location"`
	Phrase         Phrase   `json:"phrase"`
	NumberLikes    int      `json:"numberLikes"`
	NumberComments int      `json:"numberComments"`
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

func (p Photo) ValidPhoto() bool {
	return p.Location.ValidLocation() && p.Phrase.ValidPhrase() && p.Photoid >= 0 &&
		p.NumberLikes >= 0 && p.NumberComments >= 0
}

type Comment struct {
	Commentid         int    `json:"commentid"`
	CommenterUsername string `json:"commenterUsername"`
	Phrase            Phrase `json:"phrase"`
	UploadDate        string `json:"uploadDate"`
}

func (c Comment) ValidComment() bool {
	return c.Commentid >= 0 && c.Phrase.ValidPhrase()
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

func (b Ban) ValidBan() bool {
	b_lower := strings.ToLower(b.Motivation)
	return b_lower == "spam" || b_lower == "bad behaviour" || b_lower == "threats"
}

type Follow struct {
	Followid   string `json:"followid"` //This corresponds to the Username of the Follower.
	UploadDate string `json:"uploadDate"`
}

// Create a JSON Error Message Structure.
type JSONErrorMsg struct {
	Message string
}
