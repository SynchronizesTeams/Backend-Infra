package models

import "gorm.io/gorm"

type Mading struct {
	gorm.Model
	Title 		string 		`json:"title"`
	Type 		string 		`json:"type" gorm:"type:enum('infographic', 'announcement');default:'infographic'"`
	Content 	string 		`json:"content" gorm:"type:text"`
	Image		string		`json:"image"`
	Status 		string 		`json:"status" gorm:"type:enum('draft', 'published', 'archieve');default:'draft'"`
	IsActive	bool		`json:"is_active"`
}