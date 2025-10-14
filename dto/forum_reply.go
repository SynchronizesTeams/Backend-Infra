package dto

import (
	"time"
)

// DTO untuk user di reply
type ForumReplyResponse struct {
	ID        uint                   `json:"id"`
	Content   string                 `json:"content"`
	User      UserForumReplyResponse `json:"user"`
	PostID    uint                   `json:"post_id"`
	ParentID  *uint                  `json:"parent_id"`
	CreatedAt time.Time              `json:"created_at"`
	Children  []ForumReplyResponse   `json:"children,omitempty"`
}

type UserForumReplyResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
