package models

import "gorm.io/gorm"

type Product struct {
	*gorm.Model

	ProductCategoryId int             `json:"productCategoryId"`
	ProductCategory   ProductCategory `json:"productCategory" gorm:"foreignKey:ProductCategoryId"`
	Title             string          `json:"title"`
}
