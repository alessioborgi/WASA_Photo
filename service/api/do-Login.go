package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alessioborgi/WASA_Photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var username Username

	// Getting the Username from the JSON.
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		// The body was not a parseable JSON, reject it.
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalf("The Body was not a Parseable JSON!")
		return
	} else if !username.ValidUsername(*regex_username) {
		// If no error occurs, check whether the Username is a Valid User using the regex.
		// In this case it is not. Thus reject it.
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalf("The Username inserted is not Valid (Does not respect its Regex)!")
		return
	} else {
		// Here the Regex is Validated, and threfore we can proceed to give back User or create it.
		newUid, err := rt.db.DoLogin(string(username.Name))

		// dbfountain, err := rt.db.CreateFountain(fountain.ToDatabase())

		if err != nil {
			// We have an error on our side. Log the error (so we can be notified) and send a 500 to the user
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("Error During User Logging. Can't log in!")
			return
		} else {
			// It is all fine. We can send back the uuid to the User.
			w.Header().Set("Content-Type", "application/json")
			log.Println("The User Uuid is returned to the WebSite...")
			_ = json.NewEncoder(w).Encode(newUid)
		}
	}
}
