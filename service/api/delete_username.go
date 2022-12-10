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

func (rt *_router) deleteUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// First take from the path the username that is requested to be deleted.
	user := ps.ByName("username")
	user = strings.TrimPrefix(user, ":username=")
	log.Println("The Username that will be deleted is: ", user)

	if user == "" {

		// If the Username is empty, there is a bad request.
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {

		// If we arrive here, a non-empty Username has been requested to be deleted.
		// Call the DB action and wait for its response.
		err := rt.db.DeleteUsername(user)
		if errors.Is(err, database.ErrUserDoesNotExist) {

			// In this case, we have that the Username that was requested to be deleted, is not in the WASAPhoto Platform.
			w.WriteHeader(http.StatusNotFound)
			log.Fatalf("The Username that was requested to be deleted, is not a WASAPhoto User.")
			return
		} else if err != nil {
			// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user.
			// Moreover, we add the error and an additional field (`Username`) to the log entry, so that we will receive
			// the Username of the User that triggered the error.
			ctx.Logger.WithError(err).WithField("Username", user).Error("User not present in WASAPhoto. Can't delete the user profile.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {

			// If we arrive here, it means that the User Profile, has been correctly eliminated.
			log.Println("The Username has been correctly Deleted!")
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
