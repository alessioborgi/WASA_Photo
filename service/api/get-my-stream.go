package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/alessioborgi/WASA_Photo/service/api/reqcontext"
	api "github.com/alessioborgi/WASA_Photo/service/api/structs"
	"github.com/alessioborgi/WASA_Photo/service/database"
	"github.com/julienschmidt/httprouter"
)

// Get all media of a user with userid in the path
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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
	log.Println("The User that I am going to get the Stream is: ", username.Name)

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

	// Getting the list of photos
	photoListDB, err := rt.db.GetMyStream(username.Name, authorization_token)
	if errors.Is(err, database.ErrUserDoesNotExist) {

		// In this case, we have that the Username that was requested to get the list of photos, is not in the WASAPhoto Platform.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Username that was requested to get the list of stream photos, is not a WASAPhoto User.")
		return
	} else if errors.Is(err, database.ErrUserNotAuthorized) {

		// In this case, we have that the Uuid is not the same as the Profile Owner and that is has been banned, thus it cannot proceed.
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Uuid that requested to get the Stream Photo List is not the profile Owner.")
		return
	} else if errors.Is(err, database.ErrNoContent) {

		// In this case we have no Photo in the stream.
		w.WriteHeader(http.StatusNoContent)
		log.Println("There is no Photos in the Stream for this Username.")
		return
	} else if !errors.Is(err, nil) {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user.
		// Moreover, we add the error and an additional field (`Username`) to the log entry, so that we will receive
		// the Username of the User that triggered the error.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).WithField("Stream", username.Name).Error("Err: Can't get the stream of Photos.")
		return
	}

	// If we arrive here, it means that we have no errors, and we can proceed to correctly return the list to the user.
	var photoList []api.Photo
	for i := 0; i < len(photoListDB); i++ {
		var photo api.Photo
		err = photo.FromDatabase(photoListDB[i], rt.db)
		if !errors.Is(err, nil) {
			ctx.Logger.WithError(err).Error("error: Can't map photo from database to API")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		photoList = append(photoList, photo)
	}

	w.WriteHeader(http.StatusOK)
	log.Println("We can correctly return the list of Photos.")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photoList)

}
