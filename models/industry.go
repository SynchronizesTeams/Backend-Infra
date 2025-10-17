package models

import "gorm.io/gorm"

type Industry struct {
	gorm.Model
	Name 		string
	Logo		string
	Website 	string
	Description string `gorm:"type:text"`
}