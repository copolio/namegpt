package request

type SearchDomainNames struct {
	Keyword string `form:"keyword" binding:"required"`
	UserId  string `form:"user"`
}
