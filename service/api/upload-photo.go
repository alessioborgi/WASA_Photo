package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/alessioborgi/WASA_Photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// 1 - get userid from path
	mediaID := ps.ByName("media")
	mediaID = strings.TrimPrefix(mediaID, ":mediaid=")
	if mediaID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2 - get media from RB
	var media structs.Media
	err := json.NewDecoder(r.Body).Decode(&media)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !media.IsValid() {
		// Here we validated the fountain structure content (e.g., location coordinates in correct range, etc.), and we
		// discovered that the fountain data are not valid.
		// Note: the IsValid() function skips the ID check (see below).
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3 - chiamare metodo DB
	newMediaID, err := rt.db.UploadPhoto(mediaID, media.ToDatabase())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't log you in")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(newMediaID)
}
