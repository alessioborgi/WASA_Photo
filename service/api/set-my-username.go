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
	jsonpatch "github.com/evanphx/json-patch"
	"github.com/julienschmidt/httprouter"
)

// Change the Username with the newUsername.
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Variable Declaration
	var username api.Username

	// Getting the Authorization Token.
	authorization_header := strings.Split(r.Header.Get("Authorization"), " ")
	authorization_type, authorization_token := authorization_header[0], authorization_header[1]
	log.Println("The authorization Type is:", authorization_type, "and the Authorization Token is:", authorization_token)

	// We first need to check whether the authorization we have been providing is the Bearer Authentication.
	if authorization_type != BEARER {

		ctx.Logger.Error("Err: The Authentication inserted is not the Bearer Authenticaton.")
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Authentication inserted is not the Bearer Authenticaton.")
		return
	}

	// We then need to check whether the Bearer Token we are passing mastched its regex.
	if !regex_uuid.MatchString(authorization_token) {

		ctx.Logger.Error("Err: The Bearer Authentication Token you have inserted does not respect the Uuid Regex.")
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Bearer Authentication Token you have inserted does not respect the Uuid Regex.")
		return
	}

	// If we arrive here, we get a Valid Uuid (that we need to, however, check whether its in the DB and so on).

	// We can take from now from the path the username of the Users to which will be changed the Username.
	username.Name = ps.ByName("username")
	log.Println("The username that will update its Username is: ", username.Name)

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

	// Read the new content for the fountain from the request body.
	var newUsername api.Username

	// Getting the Username from the JSON.
	err := json.NewDecoder(r.Body).Decode(&newUsername)
	log.Println("The Username will be Updated to: ", newUsername.Name)

	if !errors.Is(err, nil) {

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
	}

	// If we arrive here, there is no error and the newUsername respects its regex.

	// Here we proceed by getting the UserProfile.
	usernameProfile, err := rt.db.GetUserProfile(username.Name, authorization_token)
	if errors.Is(err, database.ErrUserNotAuthorized) {

		// In this case, we have that the Uuid is not the same as the Profile Owner, thus it cannot proceed.
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Uuid that requested to update the Username, is not the Profile Owner.")
		return
	} else if !errors.Is(err, nil) {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user.
		// Moreover, we add the error and an additional field (`Username`) to the log entry, so that we will receive
		// the Username of the User that triggered the error.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).WithField("username", username.Name).Error("User not present in WASAPhoto. Can't update the Username.")
		return
	}

	// If we arrive here, we have that the Profile Retrieval has got no trouble.
	// If no error occurs during the Profile retrieval, we proceed by Marshaling the usernameProfile.
	usernameProfileMarshalled, err := json.Marshal(&usernameProfile)
	if !errors.Is(err, nil) {

		// If we fail to Marshal the Profile.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Profile Retrieval is failed.")
		return
	}

	// We can now create the Patch Operation.
	patch_operation := []byte(`[{"op": "replace", "path": "/username", "value": "` + newUsername.Name + `"}]`)
	patch, err := jsonpatch.DecodePatch(patch_operation)

	if !errors.Is(err, nil) {

		// If we fail to Marshal the Patch, it's our fault. Thus store it in the Logger.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).WithField("patch", patch).Error("Err: can't decode the patch!")
		return
	}

	// We apply the patch with username profile previously marshalled.
	newProfileByte, err := patch.Apply(usernameProfileMarshalled)

	if !errors.Is(err, nil) {

		// If we fail to Apply the Patch, it's our fault. Thus store it in the Logger.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).WithField("newProfile", newProfileByte).Error("error: can't apply patch!")
		return
	}

	// We can now convert the new profile from array of byte to json and put it into a new profile struct.
	var newProfileJson api.User
	err = json.Unmarshal(newProfileByte, &newProfileJson)

	if !errors.Is(err, nil) {

		// If we fail to UnMarshal the Json, it's our fault. Thus store it in the Logger.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Error: can't Unmarshal New Profile!")
		return
	}

	// Finally, we can call the function that updates the UserProfile.
	_, err = rt.db.SetUser(username.Name, newProfileJson.ToDatabase(), authorization_token)
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
	}

	// If we arrive here, it means that the Username, has been correctly updated.
	w.WriteHeader(http.StatusNoContent)
	log.Println("The Username has been correctly Updated!")
}
