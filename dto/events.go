package dto

import "time"

type EventRequest struct {
	Title 		string 		`json:"title"`
	Description string 		`json:"description" gorm:"type:text"`
	Category 	string 		`json:"category"`
	StartDate 	time.Time 	`json:"start_date"`
	EndDate 	time.Time 	`json:"end_date"`
	Location 	string 		`json:"location"`
	Organizer 	string 		`json:"organizer"`
}

type EventResponse struct {
	ID			uint 		`json:"id"`
	Title 		string 		`json:"title"`
	Description string 		`json:"description" `
	Category 	string 		`json:"category"`
	Visibility 	string 		`json:"visibility"`
	StartDate 	time.Time 	`json:"start_date"`
	EndDate 	time.Time 	`json:"end_date"`
	Location 	string 		`json:"location"`
	Organizer 	string 		`json:"organizer"`
	Image		string 		`json:"image"`
	CreatedAt 	time.Time 	`json:"created_at"`
}