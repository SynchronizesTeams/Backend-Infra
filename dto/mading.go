package dto

type MadingRequest struct {
	Title 		string 		`form:"title"`
	Type 		string 		`form:"type"`
	Content 	string 		`form:"content"`
	Image		uint		`form:"image"`
}

type MadingResponse struct {
	ID			uint		`json:"id"`
	Title 		string 		`json:"title"`
	Type 		string 		`json:"type"`
	Content 	string 		`json:"content" gorm:"type:text"`
	Image		string		`json:"image"`
	Status 		string 		`json:"status" gorm:"type:enum('draft', 'published', 'archieve');default:'draft'"`
	IsActive	bool		`json:"is_active"`
}