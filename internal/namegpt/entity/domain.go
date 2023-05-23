package entity

import "gorm.io/gorm"

type Domain struct {
	gorm.Model
	QueryID int
	Query   Query
}
