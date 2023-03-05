package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/alessioborgi/WASA_Photo/service/api/reqcontext"
	api "github.com/alessioborgi/WASA_Photo/service/api/structs"
	"github.com/alessioborgi/WASA_Photo/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Variable Declaration
	var username api.Username

	// Getting the Authorization Token.
	authorization_header := strings.Split(r.Header.Get("Authorization"), " ")
	authorization_type, authorization_token := authorization_header[0], authorization_header[1]
	log.Println("The authorization Type is:", authorization_type, "and the Authorization Token is:", authorization_token)

	// We first need to check whether the authorization we have been providing is the Bearer Authentication.
	if authorization_type != BEARER {

		ctx.Logger.Error("Err: The Authentication inserted is not the Bearer Authenticaton.")
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Authentication inserted is not the Bearer Authenticaton.")
		return
	}

	// We then need to check whether the Bearer Token we are passing mastched its regex.
	if !regex_uuid.MatchString(authorization_token) {

		ctx.Logger.Error("Err: The Bearer Authentication Token you have inserted does not respect the Uuid Regex.")
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
	}

	if !regex_username.MatchString(username.Name) {

		// If the username does not respect its Regex, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Update cannot be done because it has received a not valid username.")
		return
	}

	// If I arrive here, I am receiving a correct Username.
	log.Println("I can proceed since I have received a valid Username.")

	// Retrieve the photo from the Multipart-Form-Data.
	photo_body, header, errForm := r.FormFile("filename")
	if !errors.Is(errForm, nil) {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: Error encountered during the photo retrieval.")
		return
	}

	// If I arrive here, I have correctly retrieved the photo and I can proceed on saving it.
	defer photo_body.Close()

	// Make a new Photo buffer.
	photo_buffer := make([]byte, 512)

	// Read the Photo.
	_, errBody := photo_body.Read(photo_buffer)

	if !errors.Is(errBody, nil) {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(errBody).Error("Err: Error encountered during the photo reading!")
		return
	}

	// Check if the provided type of photo is valid.
	file_type := http.DetectContentType(photo_buffer)
	if file_type != "image/jpeg" && file_type != "image/png" && file_type != "image/jpg" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The image type is not one that is accepted!")
		return
	}

	// Starting to point the photo.
	_, errSeek := photo_body.Seek(0, io.SeekStart)
	if !errors.Is(errSeek, nil) {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(errSeek).Error("Err: Error encountered during the photo seeking!")
		return
	}

	// Getting the fixedUsername of the Username.
	fixedUsername, errFixedUsername := rt.db.CheckUserPresence(username.Name)
	if errors.Is(errFixedUsername, database.ErrUserDoesNotExist) {

		// The Username I am trying to get is not present in the DB.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Insertion cannot be done because it has received a not valid Username.")
		return
	} else if !errors.Is(errFixedUsername, nil) {

		// I got an error on getting the last fixedUsername.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(errFixedUsername).Error("Err: Err: The Photo Insertion cannot be done because it has encountered a strange problem")
		return
	}

	// Getting the last photo id.
	photoid, errPhotoId := rt.db.GetLastPhotoId(username.Name)
	if !errors.Is(errPhotoId, nil) {

		// I got an error on getting the last photoid.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(errPhotoId).Error("Err: I got an error on getting back the last photoId!")
		return
	}

	// Getting the Photo Phrase from the MultipartFormData.
	phrase := r.FormValue("phrase")

	// If I arrive here is all Ok. I can proceed to build up the path.
	photo_path := fixedUsername + "-photo-" + fmt.Sprint(photoid)
	log.Println("The photo name is: ", photo_path)

	// Creation of the Path URL. Save the photo in the "/tmp" folder.
	path := fmt.Sprint("./service/api/photos/", photo_path, filepath.Ext(header.Filename))
	// imageDir := "/var/tmp"
	// folderName := "images"
	// filenameImage := fmt.Sprintf("%s%s", photo_path, filepath.Ext(header.Filename))
	// path := filepath.Join(imageDir, folderName, filenameImage)
	// log.Println("The final path is: ", path)

	// Creation of a Photo and values assignment.
	var newPhoto api.Photo
	newPhoto.Photoid = photoid
	newPhoto.FixedUsername = fixedUsername
	newPhoto.Filename = path
	newPhoto.Phrase = phrase

	// Check whether we have that the newPhoto inserted respect its Regex.
	if !api.ValidPhoto(newPhoto) {

		// If no error occurs, check whether the newPhoto is a Valid Photo.
		// In this case it is not. Thus reject it.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The newPhoto received does not respect the Validation Process.")
		return
	}

	// If we arrive here, there is no error and the newUser is Valid.
	log.Println("The NewUser received has passed the Validation Process.")

	// Saving the photo in the Folder.
	f, errPathCreation := os.Create(path)

	// if !errors.Is(errPathCreation, nil) {
	// 	ctx.Logger.WithError(errPathCreation)
	// 	return
	// }

	if !errors.Is(errPathCreation, nil) {

		// I got an error on creating the path.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(errPathCreation).Error("Err: I got an error on Creating the path to the Image!")
		return
	}

	// If I arrive here, i have created the Path correctly.
	log.Println("Path created correctly!")

	// I can copy the photo.
	defer f.Close()
	_, errSaving := io.Copy(f, photo_body)

	if !errors.Is(errSaving, nil) {

		// I got an error on saving the Image.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(errSaving).Error("Err: I got an error on Saving the Image!")
		return
	}

	// Transforming it to a DB struct.
	photodb := newPhoto.ToDatabase(rt.db)

	// We can finally call the Upload of the photo.
	err := rt.db.UploadPhoto(username.Name, photodb, authorization_token)
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

		// If we arrive here, it means that the Photo has been correctly updated.
		log.Println("The Photo has been correctly Updated!")

		// Here, we can finally send back the photoid to the User, using the JSON.
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		log.Println("The PhotoId is returned to the WebSite")
		log.Println("...")
		_ = json.NewEncoder(w).Encode(newPhoto.Photoid)
	}
}
