package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
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

	// We can now get the Username(from query) of the user I am searching for from the URL.
	// If I have no username parameter, I will have an error BadRequest.
	if !r.URL.Query().Has("username") {

		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Username Parameter has not been inserted.")
		return
	}

	// Since no error has been detected, we can get the Username.
	username_search := r.URL.Query().Get("username")

	// We arrive here, we can get the fixedUsername.
	log.Println("You are searching for profile retrieval of the Username: ", username_search)

	// Getting the entire User Profile from the DB.
	profile, err = rt.db.GetUserProfile(username_search, authorization_token)

	// If we receive an error diverse from nil and ErrNoContent, we have an error in the DB Retrieval, in our side. Log the error.
	if !errors.Is(err, nil) && !errors.Is(err, database.ErrNoContent) && !errors.Is(err, database.ErrUserNotAuthorized) && !errors.Is(err, sql.ErrNoRows) && !errors.Is(err, database.ErrUserDoesNotExist) {

		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("It can't provide User Profile!")
		return

		// If error is then equal to ErrNoContent, we can return this Status.
	} else if errors.Is(err, database.ErrUserDoesNotExist) {

		// If, instead, we have that the error is No Content, we return it, meaning that we haven't found any User with the fixedUsername passed in the platform.
		w.WriteHeader(http.StatusNotFound)
		log.Println("Err: We have no User with that fixedUsername in the Platform.")
		return

	} else if errors.Is(err, database.ErrUserNotAuthorized) {

		// If we arrive here, the Uuid we have inserted, is not a Uuid present in the DB. Thus it is not Authorized.
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: We have not the requester Uuid in the Platform.")
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
