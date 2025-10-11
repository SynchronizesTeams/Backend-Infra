package models

import "gorm.io/gorm"

type UserLinks struct {
	gorm.Model
	Title 	string `json:"title"`
	Url 	string `json:"url"`
	Icon 	string `json:"icon"`

	UserID 	uint `json:"user_id"`
	User 	User 	`json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}