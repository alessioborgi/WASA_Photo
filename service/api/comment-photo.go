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
	"github.com/julienschmidt/httprouter"
)

// Add user to logged user's banned users list
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Variable Declaration
	var (
		username api.Username
		photoid  string
	)

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

	// We can take now the Username that is requesting the action, i.e., that wants to Follow another user.
	username.Name = ps.ByName("username")
	log.Println("The username that owns the photo is: ", username.Name)

	if username.Name == "" {

		// If the username is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Comment cannot be added because it has received an Empty username.")
		return
	} else if !regex_username.MatchString(username.Name) {

		// If the username does not respect its Regex, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Comment cannot be added because it has received a not valid username.")
		return
	}

	// We can now take from the path the photoid of the Username's Profile.
	photoid = ps.ByName("photoid")
	log.Println("The Photoid that will be Commented is: ", photoid)

	if photoid == "" {

		// If the username is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Comment cannot be put because it has received an Empty photoid.")
		return
	}

	// Read the Phrase Content from the request body.
	var comment api.Comment

	// Getting the Username from the JSON.
	errBody := json.NewDecoder(r.Body).Decode(&comment)

	if !errors.Is(errBody, nil) {

		// The body was not a parseable JSON, reject it.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Body was not a Parseable JSON!")
		return
	}

	// We have read correctly the body too.
	// We can therefore proceed in the Comment by calling the DB action and wait for its response.
	commentid, err := rt.db.CommentPhoto(username.Name, photoid, comment.ToDatabase(rt.db), authorization_token)
	if errors.Is(err, database.ErrUserDoesNotExist) {

		// In this case, we have that the Username that requested the action, is not present.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Username that requested the action or the username that is going to receive a comment, is not a WASAPhoto User. ")
		return
	} else if errors.Is(err, database.ErrUserNotAuthorized) {

		// In this case, we have that the Uuid is not authorized, thus it cannot proceed.
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Uuid that requested to add a comment, is not Authorized.")
		return
	} else if errors.Is(err, database.ErrPhotoDoesNotExist) {

		// In this case, we have that the Uuid is not authorized, thus it cannot proceed.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo that is going to receive a comment, is not in WASAPhoto.")
		return
	} else if !errors.Is(err, nil) {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user.
		// Moreover, we add the error and an additional field (`Username`) to the log entry, so that we will receive
		// the Username of the User that triggered the error.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).WithField("Comment", comment.Phrase).Error("User cannot put comment on the Username photo.")
		return
	} else {

		// If we arrive here, it means that the Photo's Comment has been correctly updated.
		log.Println("The Photo's Comment has been correctly Updated!")

		// Here, we can finally send back the commentid to the User, using the JSON.
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		log.Println("The Comment is returned to the WebSite")
		log.Println("...")
		_ = json.NewEncoder(w).Encode(commentid)
	}
}
