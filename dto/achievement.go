package dto

type AchievementRequest struct {
	Title 			string 		`form:"title"`
	Description		string 		`form:"description"`
	Image			string		`form:"image"`
	AchievementDate	string 		`form:"achievement_date"`
	NewsID			uint		`form:"news_id"`
}

type AchievementResponse struct {
	ID				uint		`json:"id"`
	Title 			string 		`json:"title"`
	Description		string 		`json:"description"`
	Image			string		`json:"image"`
	AchievementDate	string 		`json:"achievement_date"`
	News 			NewsAchievementResponse
}

type NewsAchievementResponse struct {
	ID 		uint		`json:"id"`
	Title 	string 		`json:"title"`
	Slug 	string 		`json:"slug"`
	Excerpt string		`json:"excerpt"`
}