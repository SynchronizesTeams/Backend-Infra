package models

import (
	"gorm.io/gorm"
)

type ForumReply struct {
	gorm.Model
	Content 	string 	`json:"content" gorm:"type:text; not null"`
	UserID uint `json:"user_id"`
	PostID uint `json:"post_id"`
	ParentID *uint `json:"parent_id"`

	User User `json:"user"`
	Post ForumPost 	`json:"post"`
	Parent *ForumReply 	`json:"parent" gorm:"foreignKey:ParentID"`
	Children []ForumReply `json:"children" gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE"`

}
