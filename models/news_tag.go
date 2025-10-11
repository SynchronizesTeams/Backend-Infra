package models

import "gorm.io/gorm"

type NewsTag struct {
	gorm.Model
	Name 	string 	`json:"name" gorm:"unique;not null"`
	News  	[]News  `gorm:"many2many:news_news_tags;" json:"news"`
}