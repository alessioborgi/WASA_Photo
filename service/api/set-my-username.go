package api

// import (
// 	"encoding/json"
// 	"errors"
// 	"log"
// 	"net/http"
// 	"strings"

// 	"github.com/alessioborgi/WASA_Photo/service/api/reqcontext"
// 	"github.com/alessioborgi/WASA_Photo/service/database"
// 	jsonpatch "github.com/evanphx/json-patch"
// 	"github.com/julienschmidt/httprouter"
// )

// // Change the Username with the newUsername.
// func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

// 	// Variable Declaration
// 	var username Username

// 	// Getting the Authorization Token.
// 	authorization_header := strings.Split(r.Header.Get("Authorization"), " ")
// 	authorization_type, authorization_token := authorization_header[0], authorization_header[1]
// 	log.Println("The authorization Type is:", authorization_type, "and the Authorization Token is:", authorization_token)

// 	// We first need to check whether the authorization we have been providing is the Bearer Authentication.
// 	if authorization_type != "Bearer" {

// 		w.WriteHeader(http.StatusUnauthorized)
// 		log.Println("Err: The Authentication inserted is not the Bearer Authenticaton.")
// 		return
// 	}

// 	// We then need to check whether the Bearer Token we are passing mastched its regex.
// 	if !regex_uuid.MatchString(authorization_token) {

// 		w.WriteHeader(http.StatusUnauthorized)
// 		log.Println("Err: The Bearer Authentication Token you have inserted does not respect the Uuid Regex.")
// 		return
// 	}

// 	// If we arrive here, we get a Valid Uuid (that we need to, however, check whether its in the DB and so on).

// 	// We can take from now from the path the username of the Users to which will be changed the Username.
// 	username.Name = ps.ByName("username")
// 	log.Println("The username that will update its Username is: ", username.Name)

// 	if username.Name == "" {

// 		// If the username is empty, there is a bad request.
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Println("Err: The Update cannot be done because it has received an Empty username.")
// 		return
// 	} else if !regex_username.MatchString(username.Name) {

// 		// If the username does not respect its Regex, there is a bad request.
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Println("Err: The Update cannot be done because it has received a not valid username.")
// 		return
// 	}

// 	// If we arrive here, a non-empty Username has been requested to be updated.
// 	// Let's therefore take the json from the body in order to see what is the newUsername.

// 	// Read the new content for the fountain from the request body.
// 	var newUsername Username

// 	// Getting the Username from the JSON.
// 	err := json.NewDecoder(r.Body).Decode(&newUsername)
// 	log.Println("The Username will be Updated to: ", newUsername.Name)

// 	if !errors.Is(err, nil) {

// 		// The body was not a parseable JSON, reject it.
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Println("Err: The Body was not a Parseable JSON!")
// 		return
// 	}

// 	// Check whether we have that the newUsername inserted respect its Regex.
// 	if !newUsername.ValidUsername(*regex_username) {

// 		// If no error occurs, check whether the Username is a Valid User using the regex.
// 		// In this case it is not. Thus reject it.
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Println("Err: The newUsername received does not respect the Regex.")
// 		return
// 	}

// 	// If we arrive here, there is no error and the newUsername respects its regex.

// 	// notice that in this case I don't check if the requested user exists because the only case in which he updates
// 	// the username is if it correspond to the logged user which I take as an assumption to exist

// 	// here only ifs logged user is trying to change his username

// 	// Here we proceed by getting the UserProfile.
// 	usernameProfile, err := rt.db.GetUserProfile(username.Name, authorization_token)
// 	if err != nil {

// 		// If we fail to retrieve the Profile.
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Println("Err: The Profile Retrieval is failed.")
// 		return
// 	}

// 	// If no error occurs during the Profile retrieval, we proceed by Marshaling the usernameProfile.
// 	usernameProfileMarshalled, err := json.Marshal(&usernameProfile)
// 	if err != nil {

// 		// If we fail to Marshal the Profile.
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Println("Err: The Profile Retrieval is failed.")
// 		return
// 	}

// 	// We can now create the Patch Operation.
// 	patch_operation := []byte(`[{"op": "replace", "path": "/username", "value": "` + newUsername.Name + `"}]`)
// 	patch, err := jsonpatch.DecodePatch(patch_operation)

// 	if err != nil {

// 		// If we fail to Marshal the Patch, it's our fault. Thus store it in the Logger.
// 		w.WriteHeader(http.StatusInternalServerError)
// 		ctx.Logger.WithError(err).WithField("patch", patch).Error("Err: can't decode the patch!")
// 		return
// 	}

// 	// We apply the patch with username profile previously marshalled.
// 	newProfileByte, err := patch.Apply(usernameProfileMarshalled)

// 	if err != nil {

// 		// If we fail to Aply the Patch, it's our fault. Thus store it in the Logger.
// 		w.WriteHeader(http.StatusInternalServerError)
// 		ctx.Logger.WithError(err).WithField("newProfile", newProfileByte).Error("error: can't apply patch!")
// 		return
// 	}

// 	// We can now convert the new profile from array of byte to json and put it into a new profile struct.
// 	var newProfileJson User
// 	err = json.Unmarshal(newProfileByte, &newProfileJson)

// 	if err != nil {

// 		// If we fail to UnMarshal the Json, it's our fault. Thus store it in the Logger.
// 		w.WriteHeader(http.StatusInternalServerError)
// 		ctx.Logger.WithError(err).Error("Error: can't Unmarshal New Profile!")
// 		return
// 	}

// 	// Finally, we can call the function that updates the UserProfile.
// 	_, err = rt.db.SetUser(newProfileJson.ToDatabase(), authorization_token)
// 	if errors.Is(err, database.ErrUserDoesNotExist) {
// 		// user profile does not exist -> return error
// 		// should never happen since it is already checked in the API
// 		ctx.Logger.WithError(err).WithField("username", newProfileJson.Username).Error("error: cannot change username because user profile does not exist")
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	} else if err != nil {
// 		ctx.Logger.WithError(err).WithField("userID", userID).Error("Can't update user profile")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	// 14 - return updated profile
// 	w.WriteHeader(http.StatusCreated)
// 	w.Header().Set("Content-Type", "application/json")
// 	_ = json.NewEncoder(w).Encode(newProfileJson)
// }
