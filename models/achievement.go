package models

import "gorm.io/gorm"

type Achievement struct {
	gorm.Model
	Title 			string 		`json:"title"`
	Description		string 		`json:"description"`
	Image			string		`json:"image"`
	AchievementDate	string 		`json:"achievement_date"`
	NewsID			uint		`json:"news_id"`
	News 			News		`json:"news"`
}