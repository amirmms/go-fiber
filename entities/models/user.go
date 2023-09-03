package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model

	ID     int    `json:"ID,omitempty"`
	Name   string `json:"name,omitempty"`
	Family string `json:"family,omitempty"`
	Mobile string `json:"mobile,omitempty"`
}
