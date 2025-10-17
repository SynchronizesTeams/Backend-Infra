package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	NIG			string 	`json:"nig"`
	FullName	string 	`json:"full_name"`
	Position	string 	`json:"position"`
	Subject 	string 	`json:"subject"`
	Photo 		string 	`json:"photo"`
	Description	string 	`json:"description" `
	UserID 		uint	`json:"user_id"`
	User		User
	Eskuls 		[]Eskul `gorm:"foreignKey:PembinaID" json:"eskuls"`
}