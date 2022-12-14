package api

import (
	"encoding/json"
	"log"
	"net/http"

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
	// 	log.Println("Err: The Authentication inserted is not the Bearer Authenticaton.")
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// } else if !regex_uuid.MatchString(authorization_token) {

	// 	// If the Authorization we have inserted does not respect its Regex, stop.
	// 	log.Println("Err: The Bearer Authentication Token you have inserted does not respect the Uuid Regex.")
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// } else {

	// 	// Here, we can proceed with our api.
	// 	// First, get the users from the DB.
	// 	log.Println("The Bearer Authentication we have inserted can proceed in the API. ")

	// 	// Secondly, we take from the path the username that is requested to be deleted.
	// 	// Get the Username of the user I am searching for from the URL.

	var fixedUsername Username
	fixedUsername.Name = ps.ByName("fixedUsername")
	log.Println("The fixedUsername that will update its Username is: ", fixedUsername.Name)

	if fixedUsername.Name == "" {

		// If the fixedUsername is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Update cannot be done because it has received an Empty fixedUsername.")
		return
	} else if !regex_fixedUsername.MatchString(fixedUsername.Name) {

		// If the fixedUsername does not respect its Regex, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Update cannot be done because it has received a not valid fixedUsername.")
		return
	} else {

		// If we arrive here, a non-empty Username has been requested to be updated.
		// Let's therefore take the json from the body in order to see what is the newUsername.

		// Read the new content for the fountain from the request body.
		var newUsername Username

		// Getting the Username from the JSON.
		err := json.NewDecoder(r.Body).Decode(&newUsername)
		log.Println("The Username will be Updated to: ", newUsername.Name)

		if err != nil {

			// The body was not a parseable JSON, reject it.
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Err: The Body was not a Parseable JSON!")
			return
		}

		// Check whether we have that the newUsername inserted respect its Regex.
		if !newUsername.ValidUsername(*regex_username) {

			// If no error occurs, check whether the Username is a Valid User using the regex.
			// In this case it is not. Thus reject it.
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Err: The newUsername received does not respect the Regex.")
			return
		} else {

			// If we arrive here, there is no error and the newUsername respects its regex.
			// We can therefore proceed in the Username Update.
			// Call the DB action and wait for its response.
			err := rt.db.SetMyUsername(fixedUsername.Name, newUsername.Name)
			if err == database.ErrUserDoesNotExist {

				// In this case, we have that the Username that was requested to be updated, is not in the WASAPhoto Platform.
				w.WriteHeader(http.StatusNotFound)
				log.Println("Err: The Username that was requested to be updated, is not a WASAPhoto User.")
				return
			} else if err != nil {
				// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user.
				// Moreover, we add the error and an additional field (`Username`) to the log entry, so that we will receive
				// the Username of the User that triggered the error.
				w.WriteHeader(http.StatusInternalServerError)
				ctx.Logger.WithError(err).WithField("fixedUsername", fixedUsername.Name).Error("User not present in WASAPhoto. Can't update the Username.")
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
