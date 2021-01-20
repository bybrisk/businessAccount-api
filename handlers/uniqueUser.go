
package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bybrisk/businessAccount-api/data"
)

// swagger:route GET /account/check/{username} businessAccount getUserPresenceStatus
// Check username is present in the database or not
//
// responses:
//	200: usernamePresenceResponse
//  422: errorValidation
//  501: errorResponse

func (p *Account) CheckAvailableUsername (w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> businessAccount-api Module")
	
	vars := mux.Vars(r)
	id := vars["username"]

	lp := data.GetAvailabilityStatus(id)

	err := lp.UserPresentOrNotStructToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}