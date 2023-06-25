package request

type SimilarDomainNames struct {
	Keyword string `form:"keyword" binding:"required"`
	UserId  string `form:"user"`
}

type RecommendDomainNames struct {
	Description     string   `json:"description" binding:"required"`
	UserId          string   `json:"userId"`
	PreviousResults []string `json:"previousResults"`
}
