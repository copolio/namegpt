package entity

import "gorm.io/gorm"

type QueryHistory struct {
	gorm.Model
	QueryID int
	Query   Query
	UserId  string
}
