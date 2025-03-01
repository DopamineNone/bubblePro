package model

import "time"

type Post struct {
	ID          int64     `gorm:"primaryKey" json:"-"`
	PostID      int64     `gorm:"unique;not null" json:"post_id"`
	AuthorID    int64     `json:"author_id" binding:"required,numeric"`
	CommunityID int64     `json:"community_id" binding:"required,numeric"`
	Title       string    `gorm:"type:varchar(128)" json:"title" binding:"required"`
	Content     string    `gorm:"type:varchar(8192)" json:"content" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Status      int8      `json:"status"`
}

type PostRequest struct {
	CommunityID int64  `json:"community_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}
