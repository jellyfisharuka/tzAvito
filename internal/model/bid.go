package model

import (
	//"time"

	//"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bid struct {
	UUID        string     ` json:"id"`
	TenderID    string     `gorm:"not null" json:"tenderId"`
	Name        string     `gorm:"type:varchar(100);not null" json:"name"`
	Description string     `gorm:"type:text" json:"description"`
	Status      StatusType `gorm:"type:varchar(50);not null" json:"status"`
	AuthorType  string     `json:"authorType"`
	AuthorID    string     `gorm:"type:varchar(50)" json:"authorId"`
	Version     int
	CreatedAt   *gorm.DeletedAt `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   *gorm.DeletedAt `gorm:"autoUpdateTime" json:"updatedAt"`

	// Foreign keys
	Tender  Tender `gorm:"foreignKey:TenderID;references:UUID" json:"tender"`
	Creator *User  `gorm:"foreignKey:AuthorID;references:ID" json:"creator"`
	//Organization Organization `gorm:"foreignKey:AuthorID;references:ID" json:"organization"`
}

type BidResponse struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Status     StatusType `json:"status"`
	AuthorType string     `json:"authorType"` // Например, "User" или "Organization"
	AuthorID   string     `json:"authorId"`
	Version    int
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

const (
	StatusTypeAPPROVED StatusType = "Approved"
	StatusTypeREJECTED StatusType = "Rejected"
)
