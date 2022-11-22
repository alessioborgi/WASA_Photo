package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

// listFountains is returning a JSON list of Fountains.
func (rt *_router) listUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	//Create a slice of Fountains in such a way to have some examples to test.
	var fountains = []User{}
	uuid := uuid.New()
	fmt.Println(uuid.String())
	// 	{
	// 		ID:        1,
	// 		Latitude:  12.45,
	// 		Longitude: 56.78,
	// 		Status:    "faulty",
	// 	},
	// 	{
	// 		ID:        2,
	// 		Latitude:  87.4,
	// 		Longitude: 12.6,
	// 		Status:    "good",
	// 	},
	// 	{
	// 		ID:        3,
	// 		Latitude:  22.3,
	// 		Longitude: 11.9,
	// 		Status:    "good",
	// 	},
	// }

	//Remember to set the Header. Here should be changed from "text/plain" to "application/json".
	//Notice that the Header is set before the encoding because the encoding should be set before w.
	w.Header().Set("content-type", "application/json")

	//Create a new Encoder and then Encode the fountains.
	err := json.NewEncoder(w).Encode(fountains)

	//Handling the Error.
	if err != nil {
		//Logging the error using the BaseLogger of rt. This is used for logging the error.
		rt.baseLogger.WithError(err).Warning("listFountains returned an error on Encode()")

		//Send Back an Internal Server Error.
		w.WriteHeader(http.StatusInternalServerError)

		//If an error occur here, means that the error has been found during the Encoding.
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: "Internal Server Error"})
		//However, if we have an error but it is not from above, this means that the error is due to
		//maybe some Network Error (e.g. when you arre in the Train and you have 3G connection etc...)
	}

}
