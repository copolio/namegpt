package dto

type CreateUser struct {
	Name string `json:"name" binding:"required"`
}
