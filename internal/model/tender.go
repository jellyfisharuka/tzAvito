package model

import (
	"time"
	//"github.com/google/uuid"
	//"gorm.io/gorm"
)

type Tender struct { // ID             string  `json:"id"` uuid
	UUID            string     `gorm:"primaryKey" json:"id"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	Status          StatusType `json:"status"`
	ServiceType     string     `gorm:"column:service_type" json:"serviceType"`
	Version         int        `json:"version"`
	OrganizationId  string     `json:"organizationId"`
	CreatorUsername string     `json:"creatorUsername"`
}
type StatusType string

const (
	StatusTypeCREATED   StatusType = "Created"
	StatusTypePUBLISHED StatusType = "Published"
	StatusTypeCLOSED    StatusType = "Closed"
)

type TenderVersion struct {
	ID          uint   `gorm:"primaryKey"`
	TenderID    string `gorm:"type:uuid;not null;index"` 
	Version     int    `gorm:"not null"`
	Name        string `gorm:"size:255"`
	Description string
	ServiceType string    `gorm:"size:255"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
type TenderResponse struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      StatusType `json:"status"`
	ServiceType string     `json:"serviceType"`
	Version     int        `json:"version"`
	CreatedAt   string     `json:"createdAt"`
}
