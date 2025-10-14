package dto

import "time"

type ForumPostRequest struct {
	Title 		string 						`form:"title"`
	Content 	string 						`form:"content"`
	SectionID 	uint 						`form:"section_id"`
}

type ForumPostResponse struct {
	ID 			uint 						`json:"id"`
	Title 		string 						`json:"title"`
	Content 	string 						`json:"content"`
	Image 		string 						`json:"image"`
	File 		string 						`json:"file"`
	Status 		string 						`json:"status"`
	Upvote 		int 						`json:"upvote"`
	Downvote 	int 						`json:"downvote"`
	ReplyCount 	int 						`json:"reply_count"`
	IsHidden 	bool 						`json:"is_hidden"`
	CreatedAt 	time.Time 					`json:"created_at"`
	User 		UserForumPostResponse 		`json:"user"`
	Section 	SectionForumPostResponse 	`json:"section"`
}

type UserForumPostResponse struct {
	ID 		uint 	`json:"id"`
	Name	string	`json:"name"`
}

type SectionForumPostResponse struct {
	ID 			uint 		`json:"id"`
	Name 		string 		`json:"name"`
}