package helpers

import (
	"go-api-infra/dto"
	"go-api-infra/models"
	"time"
)

func UpdateForumPostFromDTO(post *models.ForumPost, dto dto.ForumPostRequest) {
	if dto.Title != "" {
		post.Title = dto.Title
	}
	if dto.Content != "" {
		post.Content = dto.Content
	}
	post.UpdatedAt = time.Now()
}