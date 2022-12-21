package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/alessioborgi/WASA_Photo/service/api/reqcontext"
	"github.com/alessioborgi/WASA_Photo/service/database"
	"github.com/julienschmidt/httprouter"
)

const BEARER = "Bearer"

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Username variable declaration.
	var username Username

	// Getting the Username from the Body JSON.
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Body was not a Parseable JSON!")
		return
	}
	log.Println("The Username that will be added is: ", username)

	// First check whether we have encountered some error in the Body Retrieval.
	if !errors.Is(err, nil) {

		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Body was not a Parseable JSON!")
		return
	}

	// We can then check whether the Username we are providing is currently a Valid Username respecting the Regex.
	if !username.ValidUsername(*regex_username) {

		w.WriteHeader(http.StatusBadRequest)
		log.Println("Err: The Username inserted is not Valid (Does not respect its Regex)!")
		return
	}

	// If we arrive here, the Regex is Validated, and threfore we can proceed to give back User or create it.
	// Note that here we only send the Username to the doLogin DB function, because I am going to create a standard User.
	newUid, err := rt.db.DoLogin(string(username.Name))

	// First of all, check whether there is an error (on our side. If yes, notify the user). Note that I pass through the error also whether we have a created or already present user (not so clean).
	if !errors.Is(err, nil) && !errors.Is(err, database.Creation_Error_Inverse) && !errors.Is(err, database.Okay_Error_Inverse) {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Error During User Logging. Can't log in!")
		return
	}

	// If I arrive here, either the User has been "Created" or it was already in the Db "Ok".
	// Thus, set the header as "Created" or "OK" accordingly.
	result := err
	if errors.Is(result, database.Creation_Error_Inverse) {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	// Here, we can finally send back the Uuid to the User, using the JSON.
	w.Header().Set("Content-Type", "application/json")
	log.Println("The User Uuid is returned to the WebSite")
	log.Println("...")
	_ = json.NewEncoder(w).Encode(newUid)
}
