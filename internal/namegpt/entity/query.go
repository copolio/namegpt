package entity

import "gorm.io/gorm"

type Query struct {
	gorm.Model
	keyword string
}
