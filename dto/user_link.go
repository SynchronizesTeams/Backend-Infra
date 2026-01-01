package dto

type UserLinksRequest struct {
	Title	string 	`form:"title"`
	Url 	string 	`form:"url"`
	Icon 	string 	`form:"icon"`
}

type UserLinksResponse struct {
	ID 		uint	`json:"id"`
	Title	string 	`json:"title"`
	Url 	string 	`json:"url"`
	Icon 	string 	`json:"icon"`
	User 	UserLinksWithUserResponse `json:"user"`
}

type UserLinksWithUserResponse struct {
	ID 					uint	`json:"id"`
	Name     			string 	`json:"name"`
	Email    			string 	`json:"email"`
	PhotoUrl 			string 	`json:"photo_url"`
	Jabatan				string 	`json:"jabatan"`
}
