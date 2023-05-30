package entity

import "gorm.io/gorm"

type Query struct {
	gorm.Model
	Keyword string `gorm:"uniqueIndex"`
}
