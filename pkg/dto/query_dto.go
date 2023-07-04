package dto

import "github.com/copolio/namegpt/internal/namegpt/entity"

type GenerateDomainNames struct {
	Description string           `form:"description" binding:"required"`
	UserId      string           `form:"user"`
	Type        entity.QueryType `form:"type"`
}

type GenerateDomainNameResult struct {
	DomainName string `json:"domain"`
	Info       []struct {
		Tld       string  `json:"tld"`
		Available bool    `json:"available"`
		Price     float64 `json:"price"`
		Currency  string  `json:"currency"`
	} `json:"info"`
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
