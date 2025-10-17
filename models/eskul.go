package models

import "gorm.io/gorm"

type Eskul struct {
	gorm.Model
	Name 			string 		`json:"name"`
	Description 	string 		`json:"description" gorm:"type:text"`
	Image 			string		`json:"image"`
	PembinaID 		uint    	`json:"pembina_id"`
	Pembina   		Teacher 	`gorm:"foreignKey:PembinaID" json:"pembina"`
}