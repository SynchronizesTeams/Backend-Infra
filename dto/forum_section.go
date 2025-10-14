package dto

import "time"

type ForumSectionResponse struct {
	ID 			uint 		`json:"id"`
	Name 		string 		`json:"name"`
	Slug 		string 		`json:"slug"`
	Description string 		`json:"description"`
	Icon 		string 		`json:"icon"`
	IsActive 	bool 		`json:"is_active"`
	CreatedAt 	time.Time 	`json:"created_at"` 
}

