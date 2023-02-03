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
func (rt *_router) getPhotoComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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

		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Bearer Authentication Token you have inserted does not respect the Uuid Regex.")
		return
	}

	// If we arrive here, we get a Valid Uuid (that we need to, however, check whether its in the DB and so on).

	// We can take now from the path the username of the Users to which will be changed the Username.
	username.Name = ps.ByName("username")
	log.Println("The User that owns the Photos I am going to get its comments of is: ", username.Name)

	if username.Name == "" {

		// If the username is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Comments List cannot be acquired because it has received an Empty username.")
		return
	}

	if !regex_username.MatchString(username.Name) {

		// If the username does not respect its Regex, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Like List cannot be acquired because it has received a not valid username.")
		return
	}

	// If I arrive here, I am receiving a correct Username.
	log.Println("I can proceed since I have received a valid Username.")

	// We can now take from the path the photoid of the Username's Profile.
	photoid := ps.ByName("photoid")
	log.Println("The Photoid that for which we want to know its List of Comments is: ", photoid)

	if photoid == "" {

		// If the username is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Comment List cannot be acquired because it has received an Empty username.")
		return
	}

	// Arrived here, I got also a valid Photo.
	log.Println("I can proceed since I have received a valid Photo.")

	// If we arrive here, there is no error and we can proceed on retrieving the list of users that have put likes to Username's photoid.
	// We can therefore proceed in the Photo Likes Retrieval by calling the DB action and wait for its response.
	comments, err := rt.db.GetPhotoComments(username.Name, photoid, authorization_token)

	// If we receive an error diverse from nil and ErrNoContent, we have an error in the DB Retrieval, in our side. Log the error.
	if errors.Is(err, database.ErrUserDoesNotExist) {

		// In this case, we have that the Username that owns the photo, is not in the WASAPhoto Platform.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Username that was requested to know its photo's Comments List, is not a WASAPhoto User.")
		return
	} else if errors.Is(err, database.ErrPhotoDoesNotExist) {

		// In this case, we have that the Photo that was requested to get its Comments, is not in the WASAPhoto Platform.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo for which we would like to know its Comments Lists, is not a WASAPhoto Photo.")
		return
	} else if errors.Is(err, database.ErrUserNotAuthorized) {

		// In this case, we have that the Uuid is not the same as the Profile Owner, thus it cannot proceed.
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Uuid that requested to update the Username, is not the Profile Owner.")
		return
	} else if errors.Is(err, database.ErrNoContent) {

		// In this case we have no Username in the list of Banned Usernames.
		w.WriteHeader(http.StatusNoContent)
		log.Println("There is no Like for this Username's photoid.")
		return
	} else if !errors.Is(err, nil) {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user.
		// Moreover, we add the error and an additional field (`Username`) to the log entry, so that we will receive
		// the Username of the User that triggered the error.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).WithField("get photo", username.Name).Error("Err: Can't retrieve the Photo's Like List.")
		return
	} else {
		// If we arrive here, it means that we have no errors, and we can proceed to correctly return the list to the user.
		w.WriteHeader(http.StatusOK)
		log.Println("We can correctly return the list of Comments.")
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(comments)
	}
}
