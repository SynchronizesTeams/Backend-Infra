package models

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Title     				string `json:"title"`
	Slug  					string `json:"slug"`
	Content 				string `json:"content" gorm:"type:text"`
	Excerpt    				string `json:"excerpt"`
	Thumbnail       		string `json:"thumbnail"`
	Status 					string `json:"status" gorm:"type:enum('draft', 'published', 'archived');default:'draft'"`
	ViewCount 				int `json:"view_count"`
	Tags        			[]NewsTag  `gorm:"many2many:news_news_tags;" json:"tags"`
	AuthorID 				uint 	`json:"author_id"`
	Author 					User 	`json:"author" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
