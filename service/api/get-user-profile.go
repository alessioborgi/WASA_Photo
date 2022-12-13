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

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error
	var profile database.User
	var user User

	// Get the Username of the user I am searching for from the URL.
	username_search := ps.ByName("username")
	username_search = strings.TrimPrefix(username_search, ":username=")
	log.Println("You are searching for the Username: ", username_search)

	// Getting the entire User Profile from the DB.
	profile, err = rt.db.GetUserProfile(username_search)

	if err != nil && err != database.ErrUserDoesNotExist {

		// Error on our side. Log the error (so we can be notified) and send a 500 to the user.
		ctx.Logger.WithError(err).Error("Can't provide User profile!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if err == database.ErrUserDoesNotExist {

		// User Does not Exists.
		log.Fatalf("The User you are searching for does not Exists in WASAPhoto")
	} else {

		// The User Exists.
		log.Println("Profile Retrieval Succedeed")
		user.FromDatabase(profile)
		w.WriteHeader(http.StatusOK)
		// Send the user profile to the user
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(user)
	}
}
