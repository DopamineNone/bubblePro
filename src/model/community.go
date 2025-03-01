package model

import (
	"time"
)

// Community map community table in mysql
type Community struct {
	ID            int64     `gorm:"primaryKey" json:"id"`
	CommunityID   int64     `gorm:"unqiue;not null" json:"community_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CommunityName string    `gorm:"varchar(50);unqiue;not null" json:"community_name"`
	Introduction  string    `json:"introduction"`
}

// CommunityOverview map community list item in http server
type CommunityOverview struct {
	CommunityID   int64  `json:"id"`
	CommunityName string `json:"name"`
}

type CommunityDetail struct {
	CommunityID   int64  `json:"id"`
	CommunityName string `json:"name"`
	Introduction  string `json:"introduction"`
	CreatedAt     string `json:"created_at"`
}
