package dto

type DomainInfo struct {
	Domain    string `json:"domain"`
	Suffix    string `json:"suffix"`
	Available bool   `json:"available"`
	Price     string `json:"price"`
}

type GenerateDomainNameResult struct {
	DomainName string       `json:"domainName"`
	Info       []DomainInfo `json:"info"`
}

type GenerateDomainNames struct {
	Keyword string `form:"keyword" binding:"required"`
	UserId  string `form:"user"`
}

type RecommendDomainNames struct {
	Description     string   `json:"description" binding:"required"`
	UserId          string   `json:"userId"`
	PreviousResults []string `json:"previousResults"`
}
