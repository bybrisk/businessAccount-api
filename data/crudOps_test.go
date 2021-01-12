package data_test

import (
	"testing"
	//"github.com/go-playground/validator/v10"
	"github.com/bybrisk/structs"
	"github.com/bybrisk/businessAccount-api/data"
)

func TestAddData(t *testing.T) {

	account := &data.BusinessAccountRequest{
		UserName: "Psy Patna",
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
	if res==nil{
		t.Fail()
	}
}