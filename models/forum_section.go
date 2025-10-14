package models

import "gorm.io/gorm"

type ForumSection struct {
	gorm.Model
	Name string 
	Slug string
	Description string `gorm:"type:text"`
	Icon string
	IsActive bool `gorm:"default:true"`
}