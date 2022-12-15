package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/alessioborgi/WASA_Photo/service/api/reqcontext"
	"github.com/alessioborgi/WASA_Photo/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Variable declaration for returning then the User Profile.
	var (
		err     error
		profile database.User
		user    User
	)

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
	// We can take from the path the username of the User I want to know info of.
	username_search := ps.ByName("username")
	log.Println("The username we want to know info is: ", username_search)

	if username_search == "" {

		// If the username is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The User Profile Retrieval cannot be done because it has received an Empty username.")
		return
	} else if !regex_username.MatchString(username_search) {

		// If the username does not respect its Regex, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The User Profile Retrieval cannot be done because it has received a not valid username.")
		return
	}

	// If we arrive here, a non-empty Username has been requested.

	// Getting the entire User Profile from the DB.
	profile, err = rt.db.GetUserProfile(username_search, authorization_token)
	if errors.Is(err, database.ErrUserDoesNotExist) || errors.Is(err, sql.ErrNoRows) || errors.Is(err, database.ErrNoContent) {

		// In this case, we have that the Username that was requested to be updated, is not in the WASAPhoto Platform.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Username that was requested, is not a WASAPhoto User.")
		return
	} else if errors.Is(err, database.ErrUserNotAuthorized) {

		// In this case, we have that the Uuid is not the same as the Profile Owner, thus it cannot proceed.
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Uuid that requested the Username, has not a valid Uuid.")
		return
	} else if !errors.Is(err, nil) {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user.
		// Moreover, we add the error and an additional field (`Username`) to the log entry, so that we will receive
		// the Username of the User that triggered the error.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).WithField("username", username_search).Error("User not present in WASAPhoto. Can't update the Username.")
		return
	} else {

		// If we arrive here, it means that no errors exists (err = nil), and that the User Exists.
		log.Println("Profile Retrieval Succedeed")

		// Construct the Profile Struct from the DB
		user.FromDatabase(profile)

		// Set the User Profile all to 200 ("OK"), and send the User Profile to the User.
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(user)
	}
}
