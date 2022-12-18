package api

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alessioborgi/WASA_Photo/service/api/reqcontext"
	"github.com/alessioborgi/WASA_Photo/service/database"
	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Variable Declaration
	var username Username

	// Getting the Authorization Token.
	authorization_header := strings.Split(r.Header.Get("Authorization"), " ")
	authorization_type, authorization_token := authorization_header[0], authorization_header[1]
	log.Println("The authorization Type is:", authorization_type, "and the Authorization Token is:", authorization_token)

	// We first need to check whether the authorization we have been providing is the Bearer Authentication.
	if authorization_type != "Bearer" {

		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Authentication inserted is not the Bearer Authenticaton.")
		return
	}

	// We then need to check whether the Bearer Token we are passing mastched its regex.
	if !regex_uuid.MatchString(authorization_token) {

		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Bearer Authentication Token you have inserted does not respect the Uuid Regex.")
		return
	}

	// If we arrive here, we get a Valid Uuid (that we need to, however, check whether its in the DB and so on).

	// We can take from now from the path the username of the Users to which will be changed the Username.
	username.Name = ps.ByName("username")
	log.Println("The User to which will be added the Photo is: ", username.Name)

	if username.Name == "" {

		// If the username is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Update cannot be done because it has received an Empty username.")
		return
	} else if !regex_username.MatchString(username.Name) {

		// If the username does not respect its Regex, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Update cannot be done because it has received a not valid username.")
		return
	}

	// If we arrive here, a non-empty Username has been requested to be updated.
	// Let's therefore take the multipart-form-data in order to see what is the Photo.

	// type Form struct {
	// 	Value string
	// 	File  []*multipart.FileHeader
	// }

	// err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// var photo Form
	// fmt.Println(r.Form)
	// photo.Value = r.PostFormValue("phrase")
	// photo.File = r.MultipartForm.File["file"]
	// fmt.Println(photo)

	// fmt.Println(r.Form["name"])
	// //Access the name key - Second Approach
	// fmt.Println(r.PostForm["name"])
	// //Access the name key - Third Approach
	// fmt.Println(r.MultipartForm.Value["name"])
	// //Access the name key - Fourth Approach
	// fmt.Println(r.PostFormValue("name"))
	// //Access the age key - First Approach
	// fmt.Println(r.Form["age"])
	// //Access the age key - Second Approach
	// fmt.Println(r.PostForm["age"])
	// //Access the age key - Third Approach
	// fmt.Println(r.MultipartForm.Value["age"])
	// //Access the age key - Fourth Approach
	// fmt.Println(r.PostFormValue("age"))
	// //Access the photo key - First Approach
	photo_body, header, err := r.FormFile("photo")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer photo_body.Close()
	buff := make([]byte, 512)
	_, err = photo_body.Read(buff)
	if err != nil {
		ctx.Logger.WithError(err).Error("error: could not read photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//5 - check if photo is valid
	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/jpg" {
		ctx.Logger.WithError(err).Error("error: The provided file format is not allowed. Please upload a JPEG,JPG or PNG image")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = photo_body.Seek(0, io.SeekStart)
	if err != nil {
		ctx.Logger.WithError(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// create new mediaID
	rawMid, err := uuid.NewV4()
	if err != nil {
		// newV4 returned error -> return error
		ctx.Logger.WithError(err).Error("error encountered while creating new mediaID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Mid := rawMid.String()

	// 6 - save photo in photos folder and save with image id

	f, err := os.Create(fmt.Sprintf("./users/:"+username.Name+"/photos/%s%s", Mid, filepath.Ext(header.Filename)))
	if err != nil {
		ctx.Logger.WithError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, photo_body)
	if err != nil {
		ctx.Logger.WithError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 7 - create picture url
	picURL := fmt.Sprintf("http://localhost:3000/users/:"+username.Name+"/photos/%s%s", Mid, filepath.Ext(header.Filename))

	// 8 - take caption
	phrase := r.FormValue("phrase")

	// Getting the photoid from the DB.
	photoid, errphotoid := rt.db.GetLastPhotoId(username.Name)
	if errors.Is(errphotoid, database.ErrUserDoesNotExist) {

		// The Username we are adding the photo to is not present.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Insertion cannot be done because it has received a not valid Username.")
		return
	}

	if !errors.Is(errphotoid, nil) {

		// The Username we are adding the photo to is not present.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Insertion cannot be done because it has encountered a strange problem")
		return
	}

	// Getting the fixedUsername from the Username.
	fixedUsername, errFixedUsername := rt.db.CheckUserPresence(username.Name)
	if errors.Is(errFixedUsername, database.ErrUserDoesNotExist) {

		// The Username we are adding the photo to is not present.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Insertion cannot be done because it has received a not valid Username.")
		return
	}

	if !errors.Is(errFixedUsername, nil) && !errors.Is(errFixedUsername, database.Ok) {

		// The Username we are adding the photo to is not present.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Insertion cannot be done because it has encountered a strange problem")
		return
	}

	// If we arrive here, we have no problem during the fixedUsername retrieval.

	// 9 - create media object
	var photo Photo

	// If we arrive here, it is all ok. Thus we can continue.
	photo.Photoid = photoid
	photo.FixedUsername = fixedUsername
	photo.Filename = picURL
	photo.UploadDate = time.Now().String()
	photo.Phrase = phrase
	photo.NumberLikes = 0
	photo.NumberComments = 0

	photodb := photo.ToDatabase()
	// 5 - call upload photo database function with userID and converted media struct to database media
	err = rt.db.UploadPhoto(username.Name, photodb, authorization_token)
	if errors.Is(err, database.ErrUserNotAuthorized) {

		// In this case, we have that the Uuid is not the same as the Profile Owner, thus it cannot proceed.
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Uuid that requested to update the Username, is not the Profile Owner.")
		return
	} else if errors.Is(err, database.ErrBadRequest) {

		// In this case we have already a user having the NewUsername as Username
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The newUsername we are trying to insert is already present. ")
		return
	} else if !errors.Is(err, nil) {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user.
		// Moreover, we add the error and an additional field (`Username`) to the log entry, so that we will receive
		// the Username of the User that triggered the error.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).WithField("username", username.Name).Error("User not present in WASAPhoto. Can't update the Username.")
		return
	} else {

		// If we arrive here, it means that the Photo has been correclty updated.
		w.WriteHeader(http.StatusOK)
		log.Println("The Photo has been correctly Updated!")
	}
}

// 	// Read the new content for the User from the request body.
// 	var newUser User

// 	// 1 - get username from path
// 	username := ps.ByName("username")
// 	username = strings.TrimPrefix(username, ":username=")

// 	// 2- get the photo object from the request body.
// 	var photo Photo

// 	// Getting the Username from the JSON.
// 	err := json.NewDecoder(r.Body).Decode(&photo)

// 	if err != nil {
// 		// The body was not a parseable JSON, reject it.
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Fatalf("The Body was not a Parseable JSON!")
// 		return
// 	} else if !ValidPhoto(photo) {
// 		// If no error occurs, check whether the Username is a Valid User using the regex.
// 		// In this case it is not. Thus reject it.
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Fatalf("The Photo Pbject inserted is not Valid (Does not respect its Regex)!")
// 		return
// 	} else {
// 		// Here the Regex is Validated, and threfore we can proceed to give back User or create it.
// 		newPhotoId, err := rt.db.UploadPhoto(username, photo.ToDatabase())

// 		// dbfountain, err := rt.db.CreateFountain(fountain.ToDatabase())

// 		if err != nil {
// 			// We have an error on our side. Log the error (so we can be notified) and send a 500 to the user
// 			w.WriteHeader(http.StatusInternalServerError)
// 			ctx.Logger.WithError(err).Error("Error During User Logging. Can't log in!")
// 			return
// 		} else {
// 			// It is all fine. We can send back the uuid to the User.
// 			w.Header().Set("Content-Type", "application/json")
// 			log.Println("The User Uuid is returned to the WebSite...")
// 			_ = json.NewEncoder(w).Encode(newPhotoId)
// 		}
// 	}
// }
