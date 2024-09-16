package model

import (
	"time"

	"gorm.io/gorm"
)

type OrganizationType string

const (
	OrganizationTypeIE  OrganizationType = "IE"
	OrganizationTypeLLC OrganizationType = "LLC"
	OrganizationTypeJSC OrganizationType = "JSC"
)

type Organization struct {
	ID          string           `gorm:"primaryKey" json:"id"`
	Name        string           `gorm:"type:varchar(100);not null" json:"name"`
	Description string           `gorm:"type:text" json:"description"`
	Type        OrganizationType `gorm:"type:organization_type" json:"type"`

	CreatedAt *gorm.DeletedAt `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt *gorm.DeletedAt `gorm:"autoUpdateTime" json:"updatedAt"`

	Responsibles []User `gorm:"many2many:organization_responsible;" json:"responsibles"`
}

type OrganizationResponsible struct {
	ID             string    `gorm:"type:uuid;primary_key" json:"id"`
	OrganizationID string    `json:"organizationId"`
	UserID         string    `json:"userId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
