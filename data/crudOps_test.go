package data_test

import (
	"testing"
	"fmt"
	//"github.com/go-playground/validator/v10"
    //"github.com/bybrisk/structs"
	"github.com/bybrisk/businessAccount-api/data"
)

/*func TestAddData(t *testing.T) {

	account := &data.BusinessAccountRequest{
		UserName: "TestForDuration",
		Email: "shashank@provider.com",
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
		Latitude: 22.784793,
		Longitude:77.8382493,
	}

	res:= data.AddData(account) 

	fmt.Println(res)
	if res==nil{
		t.Fail()
	}
}*/

/*func TestGetID(t *testing.T) {
	payload:= &data.PasswordAndUsername {
		Password : "shashank",
		UserName : "VrajMilk",
	}
	res := data.GetAccountIDByUUID(payload)
	fmt.Println(res)
}*/

/*func TestUsername(t *testing.T) {
	res := data.IsUserPresent("VrajMilk")
	fmt.Println(res)
}*/

/*func TestUserPresentOrNot (t *testing.T) {
	res := data.GetAvailabilityStatus("shashank2")
	fmt.Println(res)
}*/

func TestGetData(t *testing.T) {
	res:= data.GetData("6038bd0fc35e3b8e8bd9f81a")
	fmt.Println(res)
}