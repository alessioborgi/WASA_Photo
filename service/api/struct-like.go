package api

type Like struct {
	Likeid     string   `json:"likeid"` //This corresponds to the Username of the Liker.
	UploadDate DateTime `json:"uploadDate"`
}

// ----- FINAL LIKE FUNCTION -----

// Function Method used to check for the User Validity.
func ValidLike(like Like) bool {
	return regex_username.MatchString(like.Likeid) &&
		like.UploadDate.ValidUploadDate()
}

// -----                -----
