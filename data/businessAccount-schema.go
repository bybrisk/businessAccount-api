package data

import (
	"github.com/go-playground/validator/v10"
	"github.com/bybrisk/structs"
)

// BusinessAccountRequest defines the structure for an API BusinessAccount
// swagger:model
type BusinessAccountRequest struct{
	// The url of the profile pic for this Account
	//
	// required: false
	// max length: 1000
	PicURL string `json: "picurl"`
	// The Username for this account
	//
	// required: true
	// max length: 1000
	UserName string `json: "username" validate:"required"`
	// The email Id associated with this account
	//
	// required: true
	// example: user@provider.com
	Email string `json: "email" validate:"required"`
	// Name of the business this account is for
	//
	// required: true
	// max length: 1000
	BusinessName string `json: "businessname" validate:"required"`
	// Category the business belongs to
	//
	// required: true
	// max length: 1000
	BusinessCategory string `json: "businessCat" validate:"required"`
	// Password for the account
	//
	// required: true
	// max length: 100
	Password string `json: "password" validate:"required"` //custom requirement
	// Address of the business associated with the account
	//
	// required: true
	// max length: 10000
	Address string `json: "address" validate:"required"`

	// Specify the latitude of the drop point (through your application) 
	//
	// required: true
	Latitude float64 `json:"latitude" validate:"required"`
	
	// Specify the longitude of the drop point (through your application) 
	//
	// required: true
	Longitude float64 `json:"longitude" validate:"required"`

	// Type of delivery the business requires
	//
	// required: true
	// example: Self Delivery or Bybrisk Delivery
	
	DeliveryConfig structs.DeliveryConfig `json: "deliveryConfig" validate:"required"`
	// Business plan ID
	//
	// required: true
	// example: 1, 2, 3 or 4	
	BusinessPlan string `json: "businessplan" validate:"required"`

	// Time (in seconds) a business spends on a regular customer to process after delivery
	//
	// required: false
	StandbyDuration int64 `json: "standbyDuration"`

}

//post update password
type UpdatePasswordRequest struct {
	BybID string `json:"bybID" validate:"required"`
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}

//update request
type UpdateBusinessAccountRequest struct{
	BybID string `json: "bybID" validate:"required"`
	PicURL string `json: "picurl"`
	UserName string `json: "username" validate:"required"`
	Email string `json: "email" validate:"required"`
	BusinessName string `json: "businessname" validate:"required"`
	Address string `json: "address" validate:"required"`
	DeliveryConfig structs.DeliveryConfig `json: "deliveryConfig" validate:"required"`
	StandbyDuration int64 `json: "standbyDuration" validate:"required"`
}

//get response
type BusinessAccountResponse struct{
	PicURL string `json: "picurl"`
	UserName string `json: "username"`
	Email string `json: "email"`
	BusinessName string `json: "businessname"`
	BusinessCategory string `json: "businessCat"`
	Address string `json: "address"`
	BusinessPlan string `json: "businessplan"`
	ProfileConfig structs.ProfileConfig `json:"profileConfiguration"`
	DeliveryPending string `json: "deliveryPending"`
	DeliveryDelivered string `json: "deliveryDelivered"`
	DeliveryCancelled string `json: "deliveryCancelled"`
	DeliveryTransit string `json: "deliveryTransit"`
	UserID string `json:"bybID"`
	DeliveryConfig structs.DeliveryConfig `json: "deliveryConfig"`
}

//getID request struct
type PasswordAndUsername struct {
	// Password for the account
	//
	// required: true
	Password string `json: "password" validate:"required"` //custom requirement

	// The Username for this account
	//
	// required: true
	UserName string `json: "username" validate:"required"`
}

//post response
type BusinessAccountPostSuccess struct {
	BybID string `json:"bybID"`
	Message string `json:"message"`
}

type ClusterIDArray struct {
	ClusterID []string `json:"clusterID"`
	BybID string `json:"bybID"`
}

type UserPresentOrNot struct {
	// If true then username is taken else false
	//
	IsPresent bool `json: "isPresent"`
}

func (d *BusinessAccountRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

func (d *UpdateBusinessAccountRequest) ValidateUpdateRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}

func (d *PasswordAndUsername) ValidatePasswordAndUsernameRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}

func (d *UpdatePasswordRequest) ValidateUpdatePasswordRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}