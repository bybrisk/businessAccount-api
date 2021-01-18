
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/businessAccount-api/data"
)

// swagger:route POST /account/getBybID businessAccount getBybID
// get BybID of Business Account using password and username
//
// responses:
//	200: businessAccountPostResponse
//  422: errorValidation
//  501: errorResponse

func (p *Account) GetAccountID (w http.ResponseWriter, r *http.Request){
	enableCORS(&w)

	p.l.Println("Handle POST request -> businessAccount-api Module")
	account := &data.PasswordAndUsername{}

	err:=account.FromJSONToPAsswordAndUsernameStruct(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = account.ValidatePasswordAndUsernameRequest()
	if err!=nil {
		p.l.Println("Validation error in POST request -> businessAccount-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//create account
	res := data.GetAccountIDByUUID(account)

	//writing to the io.Writer
	err = res.ResultToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}

//allow CORS request
func enableCORS (w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin","*")
}