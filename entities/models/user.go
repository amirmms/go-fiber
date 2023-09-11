package models

import "time"

type User struct {
	ID        int       `json:"id,omitempty" gorm:"primaryKey"`
	Name      string    `json:"name,omitempty"`
	Family    string    `json:"family,omitempty"`
	Mobile    string    `json:"mobile,omitempty" gorm:"unique"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
