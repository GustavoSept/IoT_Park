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

// onlyNames custom tag
func validation_names(fl validator.FieldLevel) bool {
	// accept any string that doesn't contain numbers
	// also blocks common escape characters

	pattern := `[0-9";\\<>\{\}\[\]\/=]`
	return !regexp.MustCompile(pattern).MatchString(fl.Field().String())
}

type User struct {
	ID           uuid.UUID `form:"user_id" json:"user_id"`
	First_Name   string    `form:"user_f_name" json:"user_f_name" validate:"required,onlyNames,max=20"`
	Last_Name    string    `form:"user_l_name" json:"user_l_name" validate:"required,onlyNames,max=60"`
	Office_Level string    `form:"user_oflvl" json:"user_oflvl" validate:"required,oneof=operador lavador vendedor dono gerente"`
}

type User_Auth struct {
	UserID       uuid.UUID `db:"user_id" form:"u_auth_id" json:"u_auth_id"`
	Email        string    `db:"email" form:"u_auth_email" json:"u_auth_email" validate:"required,email,max=127"`
	PasswordHash string    `db:"password_hash" json:"u_auth_pass" validate:"max=255"`
	Salt         []byte    `db:"salt" json:"u_auth_salt" validate:"max=255"`
}

type ParkingLot struct {
	ID         uuid.UUID `db:"id"`
	PLot_Name  string    `db:"pl_name" form:"pl_name"`
	AddrStreet string    `db:"addr_street" form:"pl_adstr" validate:"required,onlyNames,max=80"`
	AddrNumber int       `db:"addr_number" form:"pl_adnum" validate:"required,numeric,min=0,max=32767"`
	CEP        string    `db:"cep" form:"pl_cep" validate:"required"`
	OwnerID    uuid.UUID `db:"owner_id"`
}

// This struct is a sub-set of User, but when checking only an owner name
type Owner_onlyName struct {
	First_Name string `form:"owner_f_name" json:"owner_f_name" validate:"required,onlyNames,max=20"`
	Last_Name  string `form:"owner_l_name" json:"owner_l_name" validate:"required,onlyNames,max=60"`
}

var Users []User
var UsersAuth []User_Auth
var ParkingLots []ParkingLot
