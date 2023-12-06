package models

type User struct {
	ID           int    `json:"id"`
	First_Name   string `json:"first_name"`
	Last_Name    string `json:"last_name"`
	Office_Level string `json:"office_level"`
}

var Users []User
