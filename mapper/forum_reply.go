package mapper

import (
	"go-api-infra/dto"
	"go-api-infra/models"
)

func ForumReplyToDTO(reply models.ForumReply) dto.ForumReplyResponse {
	// Map anak-anak reply secara rekursif
	children := make([]dto.ForumReplyResponse, 0)
	for _, child := range reply.Children {
		children = append(children, ForumReplyToDTO(child))
	}

	return dto.ForumReplyResponse{
		ID:       reply.ID,
		Content:  reply.Content,
		User: dto.UserForumReplyResponse{
			ID:   reply.User.ID,
			Name: reply.User.Name,
		},
		PostID:    reply.PostID,
		ParentID:  reply.ParentID,
		CreatedAt: reply.CreatedAt,
		Children:  children,
	}
}

func ForumReplyListToDTO(replies []models.ForumReply) []dto.ForumReplyResponse {
	result := make([]dto.ForumReplyResponse, 0)
	for _, reply := range replies {
		result = append(result, ForumReplyToDTO(reply))
	}
	return result
}
