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
	ID         uuid.UUID `db:"id"`
	AddrStreet string    `db:"addr_street"`
	AddrNumber int       `db:"addr_number"`
	CEP        string    `db:"cep"`
	OwnerID    uuid.UUID `db:"owner_id"`
}

var Users []User
var ParkingLots []ParkingLot
