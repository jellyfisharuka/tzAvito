package model

import (
	"gorm.io/gorm"
)

type User struct {
	ID        string          `gorm:"primaryKey" json:"id"`
	Username  string          `gorm:"type:varchar(50);unique;not null" json:"username"`
	FirstName string          `gorm:"type:varchar(50)" json:"firstName"`
	LastName  string          `gorm:"type:varchar(50)" json:"lastName"`
	CreatedAt *gorm.DeletedAt `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt *gorm.DeletedAt `gorm:"autoUpdateTime" json:"updatedAt"`

	Organizations []Organization `gorm:"many2many:organization_responsible;" json:"organizations"`
}
