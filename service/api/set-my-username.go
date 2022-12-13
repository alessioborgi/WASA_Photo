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

func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// As first thing, we take the authorization token Uuid.
	// authorization_header := strings.Split(r.Header.Get("Authorization"), " ")
	// authorization_type, authorization_token := authorization_header[0], authorization_header[1]

	// log.Println("The authorization Type is:", authorization_type, "and the Authorization Token is:", authorization_token)
	// if authorization_type != "Bearer" {

	// 	// If the Authorization we have inserted is not the Bearer ones, stop.
	// 	log.Fatalf("The Authentication inserted is not the Bearer Authenticaton.")
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// } else if !regex_uuid.MatchString(authorization_token) {

	// 	// If the Authorization we have inserted does not respect its Regex, stop.
	// 	log.Fatalf("The Bearer Authentication Token you have inserted does not respect the Uuid Regex.")
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// } else {

	// 	// Here, we can proceed with our api.
	// 	// First, get the users from the DB.
	// 	log.Println("The Bearer Authentication we have inserted can proceed in the API. ")

	// 	// Secondly, we take from the path the username that is requested to be deleted.
	// 	// Get the Username of the user I am searching for from the URL.

	username := ps.ByName("username")
	username = strings.TrimPrefix(username, ":username=")
	log.Println("The Username that will be updated is: ", username)

	if username == "" {

		// If the Username is empty, there is a bad request.
		log.Fatalf("The Update has received an Empty username.")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {

		// If we arrive here, a non-empty Username has been requested to be updated.
		// Let's therefore take the json from the body in order to see what is the newUsername.

		// Read the new content for the fountain from the request body.
		var newUsername Username

		// Getting the Username from the JSON.
		err := json.NewDecoder(r.Body).Decode(&newUsername)
		log.Println("The Username that will be Updated is: ", newUsername)

		if err != nil {

			// The body was not a parseable JSON, reject it.
			w.WriteHeader(http.StatusBadRequest)
			log.Fatalf("The Body was not a Parseable JSON!")
			return

		} else if !newUsername.ValidUsername(*regex_username) {

			// If no error occurs, check whether the Username is a Valid User using the regex.
			// In this case it is not. Thus reject it.
			w.WriteHeader(http.StatusBadRequest)
			log.Fatalf("The newUsername received does not respect the Regex.")
			return
		} else {

			// If we arrive here, there is no error and the newUsername respects its regex.
			// We can therefore proceed in the Username Update.
			// Call the DB action and wait for its response.
			err := rt.db.SetMyUsername(username, newUsername.Name)
			if err == database.ErrUserDoesNotExist {

				// In this case, we have that the Username that was requested to be deleted, is not in the WASAPhoto Platform.
				w.WriteHeader(http.StatusNotFound)
				log.Fatalf("The Username that was requested to be deleted, is not a WASAPhoto User.")
				return
			} else if err != nil {
				// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user.
				// Moreover, we add the error and an additional field (`Username`) to the log entry, so that we will receive
				// the Username of the User that triggered the error.

				ctx.Logger.WithError(err).WithField("Username", username).Error("User not present in WASAPhoto. Can't update the Username.")
				return
			} else {

				// If we arrive here, it means that the Username, has been correctly updated.
				log.Println("The Username has been correctly Updated!")
				w.WriteHeader(http.StatusNoContent)
			}
		}
	}
}

// }
