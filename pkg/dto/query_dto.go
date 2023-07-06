package dto

import "github.com/copolio/namegpt/internal/namegpt/entity"

type GenerateDomainNames struct {
	Description string           `form:"description" binding:"required"`
	UserId      string           `form:"user"`
	Type        entity.QueryType `form:"type"`
}

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

type SimilarDomainNames struct {
	Keyword string `form:"keyword" binding:"required"`
	UserId  string `form:"user"`
}

type RecommendDomainNames struct {
	Description     string   `json:"description" binding:"required"`
	UserId          string   `json:"userId"`
	PreviousResults []string `json:"previousResults"`
}
