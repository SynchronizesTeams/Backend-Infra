package dto

type EskulRequest struct {
	Name 			string 		`form:"name"`
	Description 	string 		`form:"description"`
	PembinaID 		uint    	`form:"pembina_id"`
}


type EskulResponse struct {
	ID				uint		`json:"id"`
	Name 			string 		`json:"name"`
	Description 	string 		`json:"description" gorm:"type:text"`
	Image 			string		`json:"image"`
	PembinaID 		uint    	`json:"pembina_id"`
	Pembina 		TeacherEskulResponse
}

type TeacherEskulResponse struct {
	ID 				uint 		`json:"id"`
	FullName		string		`json:"full_name"`
}