package api

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/julienschmidt/httprouter"
// 	"github.com/nexusfb/WASAPhoto/service/api/reqcontext"
// )

// func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

// 	// 1 - get bannerID from path
// 	bannerID := ps.ByName("userid")
// 	bannerID = strings.TrimPrefix(bannerID, ":userid=")
// 	if len(bannerID) == 0 {
// 		// bannerID is empty -> return error
// 		ctx.Logger.Error("error: bannerID is empty")
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// 2 - get logged user
// 	token := r.Header.Get("Authorization")
// 	token = strings.TrimPrefix(token, "Bearer ")

// 	// 3 - logged user can ban only from his own profile, check if it is his profile
// 	if token != bannerID {
// 		ctx.Logger.Error("error: could not change username because you are not authorized ")
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	// notice that in this case I don't check if the requested user exists because the only case in which he bans
// 	// the other user is if it correspond to the logged user which I take as an assumption to exist

// 	// here only if logged user is trying to ban another user

// 	// 4 - get bannedID from path
// 	bannedID := ps.ByName("bannedid")
// 	bannedID = strings.TrimPrefix(bannedID, ":bannedid=")
// 	if len(bannedID) == 0 {
// 		// bannedID is empty -> return error
// 		ctx.Logger.Error("error: bannedID is empty")
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// 5 - check if user is valid
// 	if !rt.db.ExistenceCheck(bannedID, "user") {
// 		// banned user does not exist
// 		ctx.Logger.Error("error: banned user does not exist")
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// 6 - check validity of the ban
// 	if bannerID == bannedID {
// 		ctx.Logger.Error("error: you cannot ban yourself")
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// 7 - check if logged user already banned the specified user
// 	if rt.db.Check("ban", "bannerid", "bannedid", bannerID, bannedID) {
// 		// logged user has already banned the specified user
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	}

// 	// 8 - call ban user database function
// 	err := rt.db.BanUser(bannerID, bannedID)
// 	if err != nil {
// 		// ban user database function returned error -> return error
// 		ctx.Logger.WithError(err).Error("error: can't follow user")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	// 8 - return success (no content)
// 	w.WriteHeader(http.StatusCreated)
// }
