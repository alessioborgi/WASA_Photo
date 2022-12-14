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

func (rt *_router) GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the query string part. To do that, we need to check whether the latitude, longitude and range exists.
	// If latitude and longitude are specified, we parse them, and we filter results for them. If range is specified,
	// the value will be parsed and used as a filter. If it's not specified, 10 will be used as default (as specified in
	// the OpenAPI file).
	// If one of latitude or longitude is not specified (or both), no filter will be applied.

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
		log.Fatalf("The Authentication inserted is not the Bearer Authenticaton.")
		return
	}

	// We then need to check whether the Bearer Token we are passing mastched its regex.
	if !regex_uuid.MatchString(authorization_token) {

		w.WriteHeader(http.StatusUnauthorized)
		log.Fatalf("The Bearer Authentication Token you have inserted does not respect the Uuid Regex.")
		return
	}

	// Once that we have checked that the Beare Token we have inserted is valid, we can proceed with our api.
	// Now, get the users from the DB.
	log.Println("The Bearer Authentication we have inserted can proceed in the API.")
	users, err = rt.db.GetUsers(authorization_token)

	// If we receive an error diverse from nil and ErrNoContent, we have an error in the DB Retrieval, in our side. Log the error.
	if !errors.Is(err, nil) && !errors.Is(err, database.ErrNoContent) {

		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("We have encountered a problem in the DB GetUsers retrieval.")
		ctx.Logger.WithError(err).Error("It can't list users")
		return
	} else if errors.Is(err, database.ErrNoContent) {

		// If, instead, we have that the error is No Content, we return it, meaning that we haven't found any other User in the platform.
		w.WriteHeader(http.StatusNoContent)
		log.Fatalf("We have no User in the Platform.")
		return
	} else {
		// If we arrive here, it means that we have no errors, and we can proceed to correctly return the list to the user.
		w.WriteHeader(http.StatusOK)
		log.Println("We can correctly return the list of WASAPhoto's users.")
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(users)
	}
}
