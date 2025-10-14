package models

import (
	"gorm.io/gorm"
	"go-api-infra/dto"
)

type ForumPost struct {
	gorm.Model
	Title 		string 			`json:"title" gorm:"not null"`
	Content 	string 			`json:"content" gorm:"type:text"`
	Image 		string 			`json:"image"`
	File 		string 			`json:"file"`
	Status 		string 			`gorm:"type:enum('pending','approved','rejected','hidden');default:'pending'"`
	Upvote 		int 			`json:"upvote"`
	Downvote 	int 			`json:"downvote"`
	ReplyCount 	int 			`json:"reply_count"`
	IsHidden 	bool 			`json:"is_hidden"`

	//FK
	UserID 		uint 			`json:"user_id"`
	SectionID 	uint 			`json:"section_id"`

	// Relation
	User 		User			`json:"user"`
	Section 	ForumSection 	`json:"section"`

}

func (fp *ForumPost) ToDTO() dto.ForumPostResponse {
	return dto.ForumPostResponse{
		ID: fp.ID,
		Title: fp.Title,
		Content: fp.Content,
		Image: fp.Image,
		File: fp.File,
		Status: fp.Status,
		Upvote: fp.Upvote,
		Downvote: fp.Downvote,
		ReplyCount: fp.ReplyCount,
		IsHidden: fp.IsHidden,
		User: dto.UserForumPostResponse {
			ID: fp.UserID,
			Name: fp.User.Name,
		},
		Section: dto.SectionForumPostResponse{
			ID: fp.SectionID,
			Name: fp.Section.Name,
		},
		CreatedAt: fp.CreatedAt,
	}
}

func ToDTOList(posts []ForumPost) []dto.ForumPostResponse {
	var result []dto.ForumPostResponse
	for _, post := range posts {
		result = append(result, post.ToDTO())
	}
	return result
}

