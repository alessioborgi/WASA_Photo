package api

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/alessioborgi/WASA_Photo/service/api/reqcontext"
	"github.com/alessioborgi/WASA_Photo/service/database"
	"github.com/julienschmidt/httprouter"
)

// Add user to logged user's banned users list
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Variable Declaration
	var username Username
	var usernameFollowing Username

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

	// We can take now the Username that is requesting the action, i.e., that wants to unFollow  another user.
	username.Name = ps.ByName("username")
	log.Println("The username that wants to unFollow another user is: ", username.Name)

	if username.Name == "" {

		// If the username is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The unFollow cannot be done because it has received an Empty username.")
		return
	} else if !regex_username.MatchString(username.Name) {

		// If the username does not respect its Regex, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The unFollow cannot be done because it has received a not valid username.")
		return
	}

	// We can take now the usernameFollowing that is going to be unFollow.
	usernameFollowing.Name = ps.ByName("usernameFollowing")
	log.Println("The username that is unFollowed is: ", usernameFollowing.Name)

	if usernameFollowing.Name == "" {

		// If the usernameFollowing is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The unFollow cannot be done because it has received an Empty usernameFollowing.")
		return
	} else if !regex_username.MatchString(usernameFollowing.Name) {

		// If the usernameFollowing does not respect its Regex, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The unFollow cannot be done because it has received a not valid usernameFollowing.")
		return
	}

	// If we arrive here, there is no error and the username can unFollow the usernameFollowing since both are respecting their regex.
	// We can therefore proceed in the Ban by calling the DB action and wait for its response.
	err := rt.db.UnfollowUser(username.Name, usernameFollowing.Name, authorization_token)
	if errors.Is(err, database.ErrFollowDoesNotExist) {

		// In this case, we have that the Follow is not present in the WASAPhoto Platform.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Follow that was requested to be deleted, is not in WASAPhoto Platform.")
		return
	} else if errors.Is(err, database.ErrUserDoesNotExist) {

		// In this case, we have that the User (either Username or usernameFollowing) is not present in the WASAPhoto Platform.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Username or the UsernameFollowing that was requested to be deleted, is not in WASAPhoto Platform.")
		return
	} else if errors.Is(err, database.ErrUserNotAuthorized) {

		// In this case, we have that the Uuid is not the same as the Profile Owner, thus it cannot proceed.
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Err: The Uuid that requested to remove the Ban, is not the Profile Owner.")
		return
	} else if errors.Is(err, database.ErrBadRequest) {

		// In this case, we have that the profile is trying to self-follow.
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Uuid that requested to eliminate a self-Follow.")
		return
	} else if !errors.Is(err, nil) {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user.
		// Moreover, we add the error and an additional field (`Username`) to the log entry, so that we will receive
		// the Username of the User that triggered the error.
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).WithField("username", username.Name).Error("User not present in WASAPhoto. Can't delete the Follow.")
		return
	} else {

		// If we arrive here, it means that the Ban has been correclty deleted.
		w.WriteHeader(http.StatusNoContent)
		log.Println("The Follow has been correclty deleted.")
		return
	}
}
