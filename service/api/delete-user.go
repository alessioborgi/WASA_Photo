package api

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/alessioborgi/WASA_Photo/service/api/reqcontext"
	api "github.com/alessioborgi/WASA_Photo/service/api/structs"
	"github.com/alessioborgi/WASA_Photo/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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

	// First take from the path the username that is requested to be deleted.
	username := ps.ByName("username")
	log.Println("The username that will be deleted is: ", username)

	// First check whether we have inserted an empty username.
	if username == "" {

		// If the Username is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The username that was requested to be deleted, is Empty.")
		return
	}

	// Check then if we have inserted a Valid username w.r.t. its regex.
	if !regex_username.MatchString(username) {

		// If the username does not respect its Regex, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Update cannot be done because it has received a not valid username.")
		return
	}

	// If we arrive here, a non-empty and respecting-regex Username has been requested to be deleted.

	// First of procedeing with deleting the User, we get the list of its photos, in such a way to delete all its photos saved in local.
	// Getting the list of photos.
	photoListDB, errGetPhotos := rt.db.GetPhotos(username, authorization_token)
	if errors.Is(errGetPhotos, database.ErrUserDoesNotExist) {

		// In this case, we have that the Username that was requested to get the list of photos, is not in the WASAPhoto Platform.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Username that was requested to get the list of photos, is not a WASAPhoto User.")
		return
	} else if errors.Is(errGetPhotos, database.ErrUserNotAuthorized) {

		// In this case, we have that the Uuid is not the same as the Profile Owner and that is has been banned, thus it cannot proceed.
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Uuid that requested to get the Followings List has been banned by the Username.")
		return
	} else if errors.Is(errGetPhotos, database.ErrNoContent) {

		// In this case we have no Username in the list of Followings Usernames.
		w.WriteHeader(http.StatusNoContent)
		log.Println("There is no Photos for this Username.")

	} else if !errors.Is(errGetPhotos, nil) {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user.
		// Moreover, we add the error and an additional field (`Username`) to the log entry, so that we will receive
		// the Username of the User that triggered the error.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(errGetPhotos).WithField("username", username).Error("User not present in WASAPhoto. Can't get the list of Followings.")
		return
	}

	// If we arrive here, it means that we have no errors, and we can proceed to correctly delete the list of photos.
	var photoList []api.Photo
	for i := 0; i < len(photoListDB); i++ {
		var photo api.Photo
		errPhoto := photo.FromDatabase(photoListDB[i], rt.db)
		if !errors.Is(errPhoto, nil) {
			ctx.Logger.WithError(errPhoto).Error("error: can't map photo from database to API")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		photoList = append(photoList, photo)
	}

	// If we arrive here, we can proceed deleting the User.

	// Call the DB action and wait for its response.
	err := rt.db.DeleteUser(username, authorization_token)
	if errors.Is(err, database.ErrUserDoesNotExist) {

		// In this case, we have that the Username that was requested to be deleted, is not in the WASAPhoto Platform.
		w.WriteHeader(http.StatusNotFound)
		log.Println("Err: The fixedUsername that was requested to be deleted, is not a WASAPhoto User.")
		return
	} else if errors.Is(err, database.ErrUserNotAuthorized) {

		// In this case, we have that the Uuid is not the same as the Profile Owner, thus it cannot proceed.
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Uuid that requested to update the Username, is not the Profile Owner.")
		return
	} else if !errors.Is(err, nil) {

		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user.
		// Moreover, we add the error and an additional field (`Username`) to the log entry, so that we will receive
		// the Username of the User that triggered the error.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).WithField("username", username).Error("User not present in WASAPhoto. Can't delete the user profile.")
		return
	}

	// If we arrive here, we correclty deleted the User, and all its dependencies from the DB. We just left with deleting the photos saved in local.
	// Getting the photoName from the filename gived back from the DB.
	for i := 0; i < len(photoList); i++ {

		// Proceed in the photo Deletion also in the Photos.
		e := os.Remove(photoList[i].Filename)
		if !errors.Is(e, nil) {

			// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user.
			// Moreover, we add the error and an additional field (`Username`) to the log entry, so that we will receive
			// the Username of the User that triggered the error.
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).WithField("Photo", i).Error("Err: Can't delete photoId of Username from the Folder!")
		}
	}

	// If we arrive here, it means that the User Profile, has been correctly eliminated, included its Photos saved in Local.
	w.WriteHeader(http.StatusNoContent)
	log.Println("The username has been correctly Deleted!")
}
