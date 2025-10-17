package dto

type TeacherRequest struct {
	NIG			string 	`form:"nig"`
	FullName	string 	`form:"full_name"`
	Position	string 	`form:"position"`
	Subject 	string 	`form:"subject"`
	Description	string 	`form:"description"`
}

type TeacherResponse struct {
	ID			uint	`json:"id"`
	NIG			string 	`json:"nig"`
	FullName	string 	`json:"full_name"`
	Position	string 	`json:"position"`
	Subject 	string 	`json:"subject"`
	Photo 		string 	`json:"photo"`
	Description	string 	`json:"description"`
	User 		UserTeacherResponse `json:"user"`
}

type UserTeacherResponse struct {
	ID		uint 	`json:"id"`
	Name 	string 	`json:"name"`
}