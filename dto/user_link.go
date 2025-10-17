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
}

