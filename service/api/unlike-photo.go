package api

import (
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
func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Variable Declaration
	var (
		username      api.Username
		usernameLiker api.Username
		photoid       string
	)

	// Getting the Authorization Token.
	authorization_header := strings.Split(r.Header.Get("Authorization"), " ")
	authorization_type, authorization_token := authorization_header[0], authorization_header[1]
	log.Println("The authorization Type is:", authorization_type, "and the Authorization Token is:", authorization_token)

	// We first need to check whether the authorization we have been providing is the Bearer Authentication.
	if authorization_type != BEARER {

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

	// We can take now the Username that is requesting the action, i.e., that wants to Follow another user.
	username.Name = ps.ByName("username")
	log.Println("The username that owns the photo is: ", username.Name)

	if username.Name == "" {

		// If the username is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Like cannot be deleted because it has received an Empty username.")
		return
	} else if !regex_username.MatchString(username.Name) {

		// If the username does not respect its Regex, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Like cannot be deleted because it has received a not valid username.")
		return
	}

	// We can now take from the path the photoid of the Username's Profile.
	photoid = ps.ByName("photoid")
	log.Println("The Photoid that will be removed the Like is: ", photoid)

	if photoid == "" {

		// If the username is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Like cannot be removed because it has received an Empty photoid.")
		return
	}

	// We can take now the usernameLiker that is going to add the like.
	usernameLiker.Name = ps.ByName("usernameLiker")
	log.Println("The username that will delete the like is: ", usernameLiker.Name)

	if usernameLiker.Name == "" {

		// If the usernameLiker is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Like cannot be deleted because it has received an Empty usernameLiker.")
		return
	} else if !regex_username.MatchString(usernameLiker.Name) {

		// If the usernameLiker does not respect its Regex, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo Like cannot be deleted because it has received a not valid usernameLiker.")
		return
	}

	// If we arrive here, there is no error and the usernameLiker can delete the like to the username's Photoid since are all respecting their regex.
	// We can therefore proceed in the Like Deletion by calling the DB action and wait for its response.
	err := rt.db.UnlikePhoto(username.Name, photoid, usernameLiker.Name, authorization_token)
	if errors.Is(err, database.ErrUserDoesNotExist) {

		// In this case, we have that the Username that requested the action, is not present.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Username that requested the action or the username that is going to receive the like deletion, is not a WASAPhoto User. ")
		return
	} else if errors.Is(err, database.ErrUserNotAuthorized) {

		// In this case, we have that the Uuid is not authorized, thus it cannot proceed.
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Uuid that requested to delete the Like, is not Authorized.")
		return
	} else if errors.Is(err, database.ErrPhotoDoesNotExist) {

		// In this case, we have that the Uuid is not authorized, thus it cannot proceed.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Photo that is going to have a like to be deleted, is not in WASAPhoto.")
		return
	} else if errors.Is(err, database.ErrLikeDoesNotExists) {

		// In this case, we have that the  like that we are trying to delete in not present.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The like that we are trying to delete in not present in WASAPhoto.")
		return
	} else if errors.Is(err, database.ErrBadRequest) {

		// In this case, we have that the profile is trying to remove a self-like.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Uuid that requested to delete a self-like on one of its photos!.")
		return
	} else if !errors.Is(err, nil) {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user.
		// Moreover, we add the error and an additional field (`Username`) to the log entry, so that we will receive
		// the Username of the User that triggered the error.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).WithField("Like", usernameLiker.Name).Error("User cannot delete like on the Username photo.")
		return
	} else {

		// If we arrive here, it means that the usernameLiker, has correctly deleted the like.
		w.WriteHeader(http.StatusNoContent)
		log.Println("The usernameLiker has correclty deleted a Like to username.")
		return
	}
}
