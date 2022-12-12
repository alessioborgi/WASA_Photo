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

	// First take from the path the username that is requested to be deleted.
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
		// Let's therefore take the json from the body.
		var newUsername Username
		err := json.NewDecoder(r.Body).Decode(&newUsername)
		log.Println("The newUsername received is: ", newUsername)
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
			err := rt.db.SetMyUsername(username, string(newUsername.Name))
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

// 	// The Fountain ID in the path is a 64-bit unsigned integer. Let's parse it.
// 	id, err := strconv.ParseUint(ps.ByName("id"), 10, 64)
// 	if err != nil {
// 		// The value was not uint64, reject the action indicating an error on the client side.
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Read the new content for the fountain from the request body.
// 	var updatedFountain Fountain
// 	err = json.NewDecoder(r.Body).Decode(&updatedFountain)
// 	if err != nil {
// 		// The body was not a parseable JSON, reject it
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	} else if !updatedFountain.IsValid() {
// 		// Here we validated the fountain structure content (e.g., location coordinates in correct range, etc.), and we
// 		// discovered that the fountain data are not valid.
// 		// Note: the IsValid() function skips the ID check (see below).
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// The client is not supposed to send us the ID in the body, as the fountain ID is already specified in the path,
// 	// and it's immutable. So, here we overwrite the ID in the JSON with the `id` variable (that comes from the URL).
// 	updatedFountain.ID = id

// 	// Update the fountain in the database.
// 	err = rt.db.UpdateFountain(updatedFountain.ToDatabase())
// 	if errors.Is(err, database.ErrFountainDoesNotExist) {
// 		// The fountain (indicated by `id`) does not exist, reject the action indicating an error on the client side.
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	} else if err != nil {
// 		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
// 		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
// 		// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
// 		// the identifier of the fountain that triggered the error.
// 		ctx.Logger.WithError(err).WithField("id", id).Error("can't update the fountain")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }
