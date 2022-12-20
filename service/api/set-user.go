package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/alessioborgi/WASA_Photo/service/api/reqcontext"
	"github.com/alessioborgi/WASA_Photo/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Variable Declaration
	var username Username

	// Getting the Authorization Token.
	authorization_header := strings.Split(r.Header.Get("Authorization"), " ")
	authorization_type, authorization_token := authorization_header[0], authorization_header[1]
	log.Println("The authorization Type is:", authorization_type, "and the Authorization Token is:", authorization_token)

	// We first need to check whether the authorization we have been providing is the Bearer Authentication.
	if authorization_type != BEARER {

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
	// Let's therefore take the json from the body in order to see what is the newUsername.

	// Read the new content for the User from the request body.
	var newUser User

	// Getting the Username from the JSON.
	errBody := json.NewDecoder(r.Body).Decode(&newUser)

	if !errors.Is(errBody, nil) {

		// The body was not a parseable JSON, reject it.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Body was not a Parseable JSON!")
		return
	}

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
	err := rt.db.SetUser(username.Name, newUser.ToDatabase(), authorization_token)
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

		// If we arrive here, it means that the Username, has been correctly updated.
		w.WriteHeader(http.StatusNoContent)
		log.Println("The User has been correctly Updated!")
	}
}
