package api

import "strings"

type Ban struct {
	Banid      string     `json:"banid"` //This corresponds to the Username of the User Banned.
	UploadDate DateTime   `json:"uploadDate"`
	Motivation Motivation `json:"motivation"`
}

func (m Motivation) ValidMotivation() bool {
	m_lower := strings.ToLower(string(m))
	return m_lower == "spam" || m_lower == "bad behaviour" || m_lower == "threats"
}

// ----- FINAL BAN FUNCTION -----

// Function Method used to check for the User Validity.
func ValidBan(ban Ban) bool {
	return regex_username.MatchString(ban.Banid) &&
		ban.UploadDate.ValidUploadDate() &&
		ban.Motivation.ValidMotivation()
}

// -----                -----
