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

func (rt *_router) GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Variables Declaration.
	var (
		err   error
		users []string
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
	log.Println("The Bearer Authentication we have inserted can proceed in the API.")

	// We can now proceed with our api, getting the users from the DB.
	users, err = rt.db.GetUsers(authorization_token)

	// If we receive an error diverse from nil and ErrNoContent, we have an error in the DB Retrieval, in our side. Log the error.
	if !errors.Is(err, nil) && !errors.Is(err, database.ErrNoContent) && !errors.Is(err, sql.ErrNoRows) {

		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("It can't list users")
		return
	} else if errors.Is(err, database.ErrNoContent) || errors.Is(err, sql.ErrNoRows) {

		// If, instead, we have that the error is No Content, we return it, meaning that we haven't found any other User in the platform.
		w.WriteHeader(http.StatusNoContent)
		log.Println("Err: We have no User in the Platform.")
		return
	} else {
		// If we arrive here, it means that we have no errors, and we can proceed to correctly return the list to the user.
		w.WriteHeader(http.StatusOK)
		log.Println("We can correctly return the list of WASAPhoto's users.")
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(users)
	}
}
