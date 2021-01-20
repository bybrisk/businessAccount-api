package data

func GetData (docID string) *BusinessAccountResponse {
	account := getBusinessAccount(docID)
	return account
}

func AddData (d *BusinessAccountRequest) *BusinessAccountPostSuccess{

	var response BusinessAccountPostSuccess
	//save data to database and return ID
	if IsUserPresent(d.UserName) == false {
		id := createBusinessAccount(d)

		//get and set ProfileConfig based on Business Plan
		res:=getProfileConfig(d)
		_ = setProfileConfig(res,id)

		//set deliveryPending, deliveryDelivered, deliveryCancelled and deliveryTransit
		_ = setDeliveryStats(id)

		//save bybID to postgreSQL
		AddBybIDToPostgre(id,d)
		
		//sending response
			response = BusinessAccountPostSuccess{
			BybID: id,
			Message: "200_OK_SUCCESS",
		}
	} else {
			response = BusinessAccountPostSuccess{
			BybID: "Error",
			Message: "Username exist",
		}
	}

	return &response
}

func UpdateData (d *UpdateBusinessAccountRequest) *BusinessAccountPostSuccess {
	res := updateBusinessAccount(d)

	var response BusinessAccountPostSuccess
	//sending response
	if res == 1 {
		response = BusinessAccountPostSuccess{
			BybID: d.BybID,
			Message: "Update Done Successfully",
		}
	}

	return &response

}

func UpdatePassword (d *UpdatePasswordRequest) *BusinessAccountPostSuccess {
	var response BusinessAccountPostSuccess

	password:= getPassword(d.BybID)
	if password==d.OldPassword {
		res := updatePassword(d)
		//sending response
		if res == 1 {
		response = BusinessAccountPostSuccess{
			BybID: d.BybID,
			Message: "Password updated Successfully",
			}
		}
	}else {
		response = BusinessAccountPostSuccess{
			BybID: d.BybID,
			Message: "Authentication denied!",
		}
	}

	return &response

}

func GetAccountIDByUUID(d *PasswordAndUsername) *BusinessAccountPostSuccess {
	var response BusinessAccountPostSuccess	

	id:= GetID(d)

	if id==""{
		response = BusinessAccountPostSuccess{
			BybID: "Denied",
			Message: "Credentials donot match",
		}
	} else {
		response = BusinessAccountPostSuccess{
			BybID: id,
			Message: "Authenticated with BybID",
		}
	}

	return &response
}

func GetAvailabilityStatus (username string) *UserPresentOrNot {
	var status UserPresentOrNot

	res := IsUserPresent(username)
	status = UserPresentOrNot{
		IsPresent : res,
	}
	return &status
}