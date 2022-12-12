package api

import (
	"encoding/json"
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

	authorization_header := strings.Split(r.Header.Get("Authorization"), " ")
	authorization_type, authorization_token := authorization_header[0], authorization_header[1]

	log.Println("The authorization Type is:", authorization_type, "and the Authorization Token is:", authorization_token)
	if authorization_type != "Bearer" {

		// If the Authorization we have inserted is not the Bearer ones, stop.
		log.Fatalf("The Authentication inserted is not the Bearer Authenticaton.")
		w.WriteHeader(http.StatusForbidden)
		return
	} else if !regex_uuid.MatchString(authorization_token) {

		// If the Authorization we have inserted does not respect its Regex, stop.
		log.Fatalf("The Bearer Authentication Token you have inserted does not respect the Uuid Regex.")
		w.WriteHeader(http.StatusForbidden)
		return
	} else {

		// Here, we can proceed with our api.
		// First, get the users from the DB.
		log.Println("The Bearer Authentication we have inserted can proceed in the API. ")
		users, err = rt.db.GetUsers(authorization_token)

		if err != nil && err != database.ErrNoContent {

			// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
			// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
			log.Fatalf("We have encountered a problem in the DB GetUsers retrieval.")
			ctx.Logger.WithError(err).Error("can't list users")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if err == database.ErrNoContent {

			// In this case, the User has found any other User in the platform.
			log.Fatalf("We have no User in the Platform.")
			w.WriteHeader(http.StatusNoContent)
			return
		} else {

			// If we arrive here, it is all ok.
			// We only need to send back the list to the user.
			log.Println("We can correctly return the list of WASAPhoto's users.")
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(users)
		}
	}
}
