package entity

import "github.com/google/uuid"

type Category struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`

	// Relationships
	Products []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
}
