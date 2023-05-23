package request

type CreateUser struct {
	Name string `json:"name" binding:"required"`
}
