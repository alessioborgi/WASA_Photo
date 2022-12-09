package database

import "fmt"

func (db *appdbimpl) CommentPhoto(comment Comment, uuid string) (Comment, error) {

	// Addition of a User's Photo Comment. Here we can distinguish two cases:
	//1) We have that the User can add the comment on the photo if the fixedUsername of the uuid(commenter) has not been banned by the photo owner.
	//2) We have that the User cannot add the comment on the photo if the fixedUsername of the uuid(commenter) has been banned by the photo owner.
	ban, errBan := db.CheckBan(comment.FixedUsername, uuid)
	//Check for the error during the Query.
	if errBan != nil {
		return comment, errBan
	} else {

		//If no error occurs, checking whether the user(uuid) was banned by the fixedUsername.
		if ban == "Not Banned" {
			//If Not Banned, you can add the like object without any problem.
			_, err := db.c.Exec(`INSERT INTO Comments (commentid, photoid, fixedUsername, phrase, commenterFixedUsername, uploadDate) VALUES (?, ?, ?, ?, ?, ?)`,
				comment.Commentid, comment.PhotoId, comment.FixedUsername, comment.Phrase, comment.CommenterFixedUsername, comment.UploadDate)
			if err != nil {
				return comment, err
			} else {
				return comment, nil
			}
		} else {
			//If the User was Banned instead, returns nothing.
			fmt.Println("You cannot add the Like Object you are requiring!")
			return comment, ErrUserNotAuthorized
		}
	}
}
