package entity

import "gorm.io/gorm"

type Domain struct {
	gorm.Model
	Name string
}
