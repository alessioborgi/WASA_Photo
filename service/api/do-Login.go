package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alessioborgi/WASA_Photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Read the new content for the fountain from the request body.

	//Step 1: Verify that is a Valid User.
	var username Username

	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		fmt.Println("Error Encountered")
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !username.ValidUsername(*regex_username) {

		fmt.Println("The Username inserted is Not Valid")
		// Here we validated the fountain structure content (e.g., location coordinates in correct range, etc.), and we
		// discovered that the fountain data are not valid.
		// Note: the IsValid() function skips the ID check (see below).
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("The Username inserted is:", username)

	//Step 2: If it is valid and it does exists, Login. Return Uuid.

	//Step 3: If it is valid and does not exists. Create it. Return Uuid.

	//Step 4: If it is not valid, return Error.

	// Create the fountain in the database. Note that this function will return a new instance of the fountain with the
	// same information, plus the ID.
	newUid, err := rt.db.DoLogin(string(username.Name))
	// dbfountain, err := rt.db.CreateFountain(fountain.ToDatabase())
	if err != nil {
		fmt.Println("Error Encountered in the newuuid")
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't log you in")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		fmt.Println("All ok")
		// Send the output to the user.
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(newUid)
	}
}
