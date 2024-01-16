package models

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// using a single instance of Validate, it caches struct info
var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())

	// custom validations
	Validate.RegisterValidation("onlyNames", validation_names)
}

func validation_names(fl validator.FieldLevel) bool {
	// accept any string that doesn't contain numbers
	// also blocks common escape characters

	pattern := `[0-9";\\<>\{\}\[\]\/=]`
	return !regexp.MustCompile(pattern).MatchString(fl.Field().String())
}

type User struct {
	ID           uuid.UUID `json:"id"`
	First_Name   string    `json:"first_name" validate:"required,onlyNames,max=20"`
	Last_Name    string    `json:"last_name" validate:"required,onlyNames,max=60"`
	Office_Level string    `json:"office_level" validate:"required,oneof=operador lavador vendedor dono gerente"`
}

type ParkingLot struct {
	ID         uuid.UUID `db:"id" form:"id" json:"id"`
	AddrStreet string    `db:"addr_street" form:"addr_street" json:"addr_street" validate:"required,onlyNames,max=80"`
	AddrNumber int       `db:"addr_number" form:"addr_number" json:"addr_number" validate:"required,numeric,min=0,max=32767"`
	CEP        string    `db:"cep" form:"cep" json:"cep" validate:"required"`
	OwnerID    uuid.UUID `db:"owner_id" form:"owner_id" json:"owner_id"`
}

// This struct is a sub-set of User, but when checking only an owner name
type Owner_onlyName struct {
	First_Name string `json:"owner_first_name" form:"owner_first_name" validate:"required,onlyNames,max=20"`
	Last_Name  string `json:"owner_last_name" form:"owner_last_name" validate:"required,onlyNames,max=60"`
}

var Users []User
var ParkingLots []ParkingLot
