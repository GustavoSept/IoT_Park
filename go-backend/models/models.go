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
	ID           uuid.UUID `form:"user_id"`
	First_Name   string    `form:"user_f_name" validate:"required,onlyNames,max=20"`
	Last_Name    string    `form:"user_l_name" validate:"required,onlyNames,max=60"`
	Office_Level string    `form:"user_oflvl" validate:"required,oneof=operador lavador vendedor dono gerente"`
}

type User_Auth struct {
	UserID       uuid.UUID `form:"uAuth_user_id"`
	Email        string    `form:"uAuth_email" validate:"required,email, max=127"`
	PasswordHash string    `validate:"required, max=255"`
	Salt         []byte    `validate:"required, max=255"`
}

type ParkingLot struct {
	ID         uuid.UUID `db:"id" form:"pl_id" json:"pl_id"`
	PLot_Name  string    `db:"pl_name" form:"pl_name" json:"pl_name"`
	AddrStreet string    `db:"addr_street" form:"pl_adstr" json:"pl_adstr" validate:"required,onlyNames,max=80"`
	AddrNumber int       `db:"addr_number" form:"pl_adnum" json:"pl_adnum" validate:"required,numeric,min=0,max=32767"`
	CEP        string    `db:"cep" form:"pl_cep" json:"cep" validate:"required"`
	OwnerID    uuid.UUID `db:"owner_id" form:"owner_id" json:"owner_id"`
}

// This struct is a sub-set of User, but when checking only an owner name
type Owner_onlyName struct {
	First_Name string `json:"owner_first_name" form:"owner_first_name" validate:"required,onlyNames,max=20"`
	Last_Name  string `json:"owner_last_name" form:"owner_last_name" validate:"required,onlyNames,max=60"`
}

var Users []User
var ParkingLots []ParkingLot
