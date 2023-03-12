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

	"github.com/alessioborgi/WASA_Photo/service/api/reqcontext"
	api "github.com/alessioborgi/WASA_Photo/service/api/structs"
	"github.com/alessioborgi/WASA_Photo/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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
	log.Println("The User that will be updated is: ", username.Name)

	if username.Name == "" {

		// If the username is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Update cannot be done because it has received an Empty username.")
		return
	} else if !regex_username.MatchString(username.Name) {

		// If the username does not respect its Regex, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Update cannot be done because it has received a not valid username.")
		return
	}

	// If we arrive here, a non-empty Username has been requested to be updated.
	// Retrieve the photo from the Multipart-Form-Data.
	photo_body, header, errForm := r.FormFile("photoProfile")
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

	// If we arrive here it means it has errFixedUsername=nil.

	// Setting the photo id (that for all the photo Profiles is zero).
	photoid := 0

	// Constructing the photo path.
	imageDir := "/tmp"
	folderName := "images"

	photo_path := fmt.Sprintf("%s%s", fixedUsername+"-photo-"+fmt.Sprint(photoid), filepath.Ext(header.Filename))
	log.Println("The photo name is: ", photo_path)

	path := filepath.Join(imageDir, folderName, photo_path)

	// Read the new content for the User from the request body.
	var newUser api.User
	newUser.Username = r.FormValue("username")
	newUser.PhotoProfile = path
	newUser.Biography = r.FormValue("biography")
	newUser.Name = r.FormValue("name")
	newUser.Surname = r.FormValue("surname")
	newUser.DateOfBirth = api.Date(r.FormValue("dateOfBirth"))
	newUser.Email = api.Email(r.FormValue("email"))
	newUser.Nationality = r.FormValue("nationality")
	newUser.Gender = api.Gender(r.FormValue("gender"))

	// Check whether we have that the newUsername inserted respect its Regex.
	if !newUser.ValidUser() {

		// If no error occurs, check whether the User is a Valid User.
		// In this case it is not. Thus reject it.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The NewUser received does not respect the Validation Process.")
		return
	}

	// If we arrive here, there is no error and the newUser is Valid.
	log.Println("The NewUser received has passed the Validation Process.")

	// We can therefore proceed in the User Update.
	// Call the DB action and wait for its response.
	filename, err := rt.db.SetUser(username.Name, newUser.ToDatabase(), authorization_token)
	if errors.Is(err, database.ErrUserDoesNotExist) {

		// In this case, we have that the Username that was requested to be updated, is not in the WASAPhoto Platform.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Username that was requested to be updated, is not a WASAPhoto User.")
		return
	} else if errors.Is(err, database.ErrUserNotAuthorized) {

		// In this case, we have that the Uuid is not the same as the Profile Owner, thus it cannot proceed.
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Uuid that requested to update the Username, is not the Profile Owner.")
		return
	} else if errors.Is(err, database.ErrBadRequest) {

		// In this case we have already a user having the NewUsername as Username.
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
	}

	// If we arrive here, it means that the Username, has been correctly updated.

	// I need now to take the old photo and delete it from the folder.
	// Getting the photoName of the old photoProfile from the filename gived back from the DB.
	log.Println("The PhotoProfile to be deleted is: ", filename)

	// Check whether the photoProfile that we are going to delete is the diverse from the actual "NoPhoto" we find when we create a new Profile.
	if filename != "" {

		// Proceed in the photo Deletion in the folder.
		e := os.Remove(filename)
		if !errors.Is(e, nil) {

			ctx.Logger.WithError(err).WithField("Photo", photoid).Error("Err: Can't delete photoId of Username from the Folder!")
		}
		log.Println("Photo Correctly Deleted from the Photos Folder!")

	}

	// I can therefore save the PhotoProfile in local tmp folder.
	// Saving the photo in the Folder.
	f, errPathCreation := os.Create(path)

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

	// Set the result to NoContent if no error occurs.
	w.WriteHeader(http.StatusNoContent)
	log.Println("The User has been correctly Updated!")
}
