package data_test

import (
	"testing"
	"fmt"
	//"github.com/go-playground/validator/v10"
    "github.com/bybrisk/structs"
	"github.com/bybrisk/businessAccount-api/data"
)

func TestAddData(t *testing.T) {

	account := &data.BusinessAccountRequest{
		UserName: "User Cluster12",
		Email: "user@provider.com",
		BusinessName: "CSoL",
		BusinessCategory: "grocery",
		Password: "shashank",
		Address: "A.G Colony Chetna Samiti",
		DeliveryConfig: structs.DeliveryConfig{
			AutoScaling: false,
			AvgWorkingHours: 8,
			BybriskDelivery: false,
			DeliveryAgent: 6,
			InstantDelivery: true,
		},
		BusinessPlan: "1",
		PicURL: "img/pic.jpg",
	}

	res:= data.AddData(account) 

	fmt.Println(res)
	if res==nil{
		t.Fail()
	}
}

/*func TestGetID(t *testing.T) {
	payload:= &data.PasswordAndUsername {
		Password : "shashank",
		UserName : "Psy Patna",
	}
	res := data.GetAccountIDByUUID(payload)
	fmt.Println(res)
}*/

/*func TestUsername(t *testing.T) {
	res := data.IsUserPresent("Psy Patna")
	fmt.Println(res)
}*/

/*func TestUserPresentOrNot (t *testing.T) {
	res := data.GetAvailabilityStatus("shashank2")
	fmt.Println(res)
}*/