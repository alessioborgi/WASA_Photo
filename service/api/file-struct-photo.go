package api

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

// Create a JSON Error Message Structure.
type JSONErrorMsg struct {
	Message string
}
