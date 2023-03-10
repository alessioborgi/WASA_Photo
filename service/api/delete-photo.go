package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/alessioborgi/WASA_Photo/service/api/reqcontext"
	api "github.com/alessioborgi/WASA_Photo/service/api/structs"
	"github.com/alessioborgi/WASA_Photo/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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
	log.Println("The User that owns the Photo is: ", username.Name)

	if username.Name == "" {

		// If the username is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Retrieval cannot be done because it has received an Empty username.")
		return
	}

	if !regex_username.MatchString(username.Name) {

		// If the username does not respect its Regex, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Retrieval cannot be done because it has received a not valid username.")
		return
	}

	// If I arrive here, I am receiving a correct Username.
	log.Println("I can proceed since I have received a valid Username.")

	// We can now take from the path the photoid of the Username's Profile.
	photoid := ps.ByName("photoid")
	log.Println("The Photoid that we want to delete is: ", photoid)

	if photoid == "" {

		// If the username is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Deletion cannot be done because it has received an Empty username.")
		return
	}

	// Arrived here, I got also a valid Photo.
	log.Println("I can proceed since I have received a Valid Photo.")

	// Call the DB action and wait for its response.
	filename, err := rt.db.DeletePhoto(username.Name, photoid, authorization_token)
	if errors.Is(err, database.ErrUserDoesNotExist) {

		// In this case, we have that the Username that requested to delete its photo, is not in the WASAPhoto Platform.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The fixedUsername that requested to delete the Photo, is not a WASAPhoto User.")
		return
	} else if errors.Is(err, database.ErrPhotoDoesNotExist) {

		// In this case, we have that the Photo that was requested to be deleted, is not in the WASAPhoto Platform.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo that requested to be deleted the Photo, is not a WASAPhoto Photo.")
		return
	} else if errors.Is(err, database.ErrUserNotAuthorized) {

		// In this case, we have that the Uuid is not the same as the Profile Owner, thus it cannot proceed.
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Uuid that requested to delete the Photo, is not the Profile Owner.")
		return
	} else if !errors.Is(err, nil) {

		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user.
		// Moreover, we add the error and an additional field (`Username`) to the log entry, so that we will receive
		// the Username of the User that triggered the error.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).WithField("Photo", photoid).Error("Err: Can't delete photoId of Username!")
		return
	}

	// If we arrive here, it means that the User Profile, has been correctly eliminated by the DB.
	// We can also proceed to elimnate it from the folder photos.

	// Getting the photoName from the filename gived back from the DB.
	photoName := filename[strings.LastIndex(filename, "/")+1:]

	// Constructing the path that will indicate what photo to delete.
	path := fmt.Sprint("./tmp/", photoName)

	// Proceed in the photo Deletion also in the Photos.
	e := os.Remove(path)
	if !errors.Is(e, nil) {

		ctx.Logger.WithError(err).WithField("Photo", photoid).Error("Err: Can't delete photoId of Username from the Folder!")
	}
	log.Println("Photo Correctly Deleted from the Photos Folder!")

	// Returning back the NoContent.
	w.WriteHeader(http.StatusNoContent)
	log.Println("The Photo has been totally correctly Deleted!")
}
