package database

/*
	DESCRIPTION:
	This file contains the User, Photo, Comment, Like, Ban and Follow Structures.
*/

// Create a User structure.
type User struct {
	FixedUsername   string //Primary Key (PK).
	Uuid            string
	Username        string
	Biography       string //Optional
	DateOfCreation  string
	NumberOfPhotos  int64
	NumberFollowers int64
	NumberFollowing int64
	Name            string
	Surname         string
	DateOfBirth     string
	Email           string
	Nationality     string
	Gender          string
}

type Photo struct {
	Photoid        int64  //Primary Key (PK).
	FixedUsername  string //Primary & Foreign Key (PK, FK).
	Filename       byte
	UploadDate     string
	Phrase         string //Optional.
	NumberLikes    int64
	NumberComments int64
	Latitude       float64
	Longitude      float64
}

type Comment struct {
	Commentid              int64  //Primary Key (PK).
	PhotoId                int64  //Primary & Foreign Key (PK, FK).
	FixedUsername          string //Primary & Foreign Key (PK, FK).
	CommenterFixedUsername string //Foreign Key (FK) (Not a PK though!).
	Phrase                 string
	UploadDate             string
}

type Like struct {
	Likeid        string //Primary Key (PK). This corresponds to the Username of the Liker.
	PhotoId       int64  //Primary & Foreign Key (PK, FK).
	FixedUsername string //Primary & Foreign Key (PK, FK).
	UploadDate    string
}

type Ban struct {
	FixedUsernameBanner string //Primary & Foreign Key (PK, FK). This corresponds to the fixedUsername of the User that wants to Ban the other fixedUsername).
	FixedUsernameBanned string //Primary & Foreign Key (PK, FK). This corresponds to the fixedUsername of the User Banned.
}

type Follow struct {
	FixedUsername          string //Primary & Foreign Key (PK, FK). This corresponds to the fixedUsername of the person that wants to follow someone.
	FixedUsernameFollowing string //Primary & Foreign Key (PK, FK). This corresponds to the fixedUsername of the person he wants to Follow.
}
