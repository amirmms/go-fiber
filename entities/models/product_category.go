package models

import "gorm.io/gorm"

type ProductCategory struct {
	*gorm.Model

	ID    int    `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}
