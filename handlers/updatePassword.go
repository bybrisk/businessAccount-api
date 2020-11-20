
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/businessAccount-api/data"
)

// swagger:route POST /accountUpdate/password businessAccount updatePassword
// Update password of an existing Business Account with ID
//
// responses:
//	200: businessAccountPostResponse
//  422: errorValidation
//  501: errorResponse

func (p *Account) UpdateAccountPassword (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> businessAccount-api Module")
	account := &data.UpdatePasswordRequest{}

	err:=account.FromJSONUpdatePasswordRequest(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data 
	err = account.ValidateUpdatePasswordRequest()
	if err!=nil {
		p.l.Println("Validation error in POST request -> businessAccount-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//update account password
	result := data.UpdatePassword(account)

	//writing to the io.Writer
	err = result.ResultToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}