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

// Get the Banned users list of a Username.
func (rt *_router) getBannedUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Variable Declaration
	var username api.Username

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

	// We can take now the Username that is requesting the action, i.e., that wants to Ban another user.
	username.Name = ps.ByName("username")
	log.Println("The username that want to know the Bans is: ", username.Name)

	if username.Name == "" {

		// If the username is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Banned User Retrieval cannot be done because it has received an Empty username.")
		return
	} else if !regex_username.MatchString(username.Name) {

		// If the username does not respect its Regex, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Banned User Retrieval cannot be done because it has received a not valid username.")
		return
	}

	// If we arrive here, there is no error and we can proceed on retrieving the banned users of Username.
	// We can therefore proceed in the Bans Retrieval by calling the DB action and wait for its response.
	users, err := rt.db.GetBannedUsers(username.Name, authorization_token)

	// If we receive an error diverse from nil and ErrNoContent, we have an error in the DB Retrieval, in our side. Log the error.
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
	} else if errors.Is(err, database.ErrNoContent) {

		// In this case we have no Username in the list of Banned Usernames.
		w.WriteHeader(http.StatusOK)
		log.Println("There is no Username in the list of Banned Usernames. ")
		return
	} else if !errors.Is(err, nil) {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user.
		// Moreover, we add the error and an additional field (`Username`) to the log entry, so that we will receive
		// the Username of the User that triggered the error.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).WithField("username", username.Name).Error("User not present in WASAPhoto. Can't update the Username.")
		return
	} else {
		// If we arrive here, it means that we have no errors, and we can proceed to correctly return the list to the user.
		w.WriteHeader(http.StatusOK)
		log.Println("We can correctly return the list of Banned Usernames.")
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(users)
	}
}
