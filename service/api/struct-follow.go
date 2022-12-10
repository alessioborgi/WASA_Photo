package api

type Follow struct {
	Followid   string   `json:"followid"` //This corresponds to the Username of the Follower.
	UploadDate DateTime `json:"uploadDate"`
}

// ----- FINAL FOLLOW FUNCTION -----

// Function Method used to check for the User Validity.
func ValidFollow(follow Follow) bool {
	return regex_username.MatchString(follow.Followid) &&
		follow.UploadDate.ValidUploadDate()
}

// -----                -----
