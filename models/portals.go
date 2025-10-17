package models

import "gorm.io/gorm"

type Portal struct {
	gorm.Model
	Name  			string
	Description 	string
	URL 			string
	Logo 			string
	Category 		string
	IsSSOEnabled 	bool `gorm:"default:false"`	
	IsActive 		bool `gorm:"default:true"`
}