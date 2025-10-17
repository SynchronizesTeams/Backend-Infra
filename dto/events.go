package dto

import "time"

type EventRequest struct {
	Title 		string 		`form:"title"`
	Description string 		`form:"description" gorm:"type:text"`
	Category 	string 		`form:"category"`
	StartDate 	string 		`form:"start_date"`
	EndDate 	string 		`form:"end_date"`
	Location 	string 		`form:"location"`
	Organizer 	string 		`form:"organizer"`
}

type EventResponse struct {
	ID			uint 		`json:"id"`
	Title 		string 		`json:"title"`
	Description string 		`json:"description" `
	Category 	string 		`json:"category"`
	Visibility 	string 		`json:"visibility"`
	StartDate 	string 		`json:"start_date"`
	EndDate 	string 		`json:"end_date"`
	Location 	string 		`json:"location"`
	Organizer 	string 		`json:"organizer"`
	Image		string 		`json:"image"`
	CreatedAt 	time.Time 	`json:"created_at"`
}