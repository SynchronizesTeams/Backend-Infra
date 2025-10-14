package mapper

import (
	"go-api-infra/dto"
	"go-api-infra/models"
)

func ForumPostToDTO(fp models.ForumPost) dto.ForumPostResponse {
	return dto.ForumPostResponse{
		ID:         fp.ID,
		Title:      fp.Title,
		Content:    fp.Content,
		Image:      fp.Image,
		File:       fp.File,
		Status:     fp.Status,
		Upvote:     fp.Upvote,
		Downvote:   fp.Downvote,
		ReplyCount: fp.ReplyCount,
		IsHidden:   fp.IsHidden,
		User: dto.UserForumPostResponse{
			ID:   fp.UserID,
			Name: fp.User.Name,
		},
		Section: dto.SectionForumPostResponse{
			ID:   fp.SectionID,
			Name: fp.Section.Name,
		},
		CreatedAt: fp.CreatedAt,
	}
}

func ForumPostListToDTO(posts []models.ForumPost) []dto.ForumPostResponse {
	var result []dto.ForumPostResponse
	for _, post := range posts {
		result = append(result, ForumPostToDTO(post))
	}
	return result
}
