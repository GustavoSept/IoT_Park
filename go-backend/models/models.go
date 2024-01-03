package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	First_Name   string    `json:"first_name"`
	Last_Name    string    `json:"last_name"`
	Office_Level string    `json:"office_level"`
}

type ParkingLot struct {
	ID         uuid.UUID `db:"id" form:"id" json:"id"`
	AddrStreet string    `db:"addr_street" form:"addr_street" json:"addr_street"`
	AddrNumber int       `db:"addr_number" form:"addr_number" json:"addr_number"`
	CEP        string    `db:"cep" form:"cep" json:"cep"`
	OwnerID    uuid.UUID `db:"owner_id" form:"owner_id" json:"owner_id"`
}

var Users []User
var ParkingLots []ParkingLot
