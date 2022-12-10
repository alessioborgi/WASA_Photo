package api

type Comment struct {
	Commentid         int      `json:"commentid"`
	CommenterUsername string   `json:"commenterUsername"`
	Phrase            string   `json:"phrase"`
	UploadDate        DateTime `json:"uploadDate"`
}

// ----- FINAL COMMENT FUNCTION -----

// Function Method used to check for the User Validity.
func ValidComment(comment Comment) bool {
	return comment.Commentid >= 0 &&
		regex_username.MatchString(comment.CommenterUsername) &&
		len(comment.Phrase) >= 5 && len(comment.Phrase) <= 1000 &&
		comment.UploadDate.ValidUploadDate()
}

// -----                -----
