package service

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/internal/namegpt/repository"
	"github.com/copolio/namegpt/pkg/client/chatgpt"
	"github.com/copolio/namegpt/pkg/client/gabia"
	"github.com/copolio/namegpt/pkg/dto"
)

type QueryUseCase interface {
	Handle(request dto.SimilarDomainNames) (domainNames [][]*gabia.RegistCheckResult, err error)
}

func NewQueryUseCase() QueryUseCase {
	return &QueryService{
		queryRepository:        repository.NewQueryRepository(),
		queryHistoryRepository: repository.NewQueryHistoryRepository(),
		domainNameRepository:   repository.NewDomainNameRepository(),
	}
}

type QueryService struct {
	queryRepository        repository.QueryRepository
	queryHistoryRepository repository.QueryHistoryRepository
	domainNameRepository   repository.DomainNameRepository
}

func (q QueryService) Handle(request dto.SimilarDomainNames) (registCheckResults [][]*gabia.RegistCheckResult, err error) {
	db := config.GetGormDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Find Or Create Query
	q.queryRepository.WithTransaction(tx)
	query, err1 := q.queryRepository.FindOrCreate(entity.Query{Keyword: request.Keyword, Type: entity.SIMILAR})
	if err1 != nil {
		tx.Rollback()
		return nil, err1
	}
	// Create query history
	q.queryHistoryRepository.WithTransaction(tx)
	_, err2 := q.queryHistoryRepository.Save(entity.QueryHistory{
		QueryID: query.ID,
		Query:   entity.Query{},
		UserId:  request.UserId,
	})
	if err2 != nil {
		tx.Rollback()
		return nil, err2
	}
	// Return cached result if exists
	if len(query.DomainNames) > 0 {
		result := mapDomainNameToGenerateDomainNamesResult(query.DomainNames)
		return result, tx.Commit().Error
	}

	// Ask ChatGPT for domains
	domains, err3 := chatgpt.GetSimilarDomains(request.Keyword)
	if err3 != nil {
		tx.Rollback()
		return nil, err3
	}

	q.domainNameRepository.WithTransaction(tx)
	var domainNames []entity.DomainName
	for _, domain := range domains {
		domainName, err4 := q.domainNameRepository.Save(entity.DomainName{
			Name:    domain,
			QueryID: query.ID,
		})
		if err4 != nil {
			tx.Rollback()
			return nil, err4
		}
		domainNames = append(domainNames, *domainName)
	}

	return mapDomainNameToGenerateDomainNamesResult(domainNames), tx.Commit().Error
}

func mapDomainNameToGenerateDomainNamesResult(domainNames []entity.DomainName) [][]*gabia.RegistCheckResult {
	suffixes := []string{".com", ".co.kr", ".kr", ".shop", ".store", ".net", ".site", ".org", ".me", ".한국", ".io",
		".biz", ".tv", ".info", ".xyz", ".ai", ".company", ".app", ".us", ".jp", ".cn", ".vn", ".tw", ".im", ".club", ".co"}

	result := make([][]*gabia.RegistCheckResult, len(domainNames))
	for i := 0; i < len(domainNames); i++ {
		result[i] = make([]*gabia.RegistCheckResult, len(suffixes))
	}

	for i, domainName := range domainNames {
		for j, suffix := range suffixes {
			domain := domainName.Name + suffix
			registCheckResult, err := gabia.CheckDomainRegist(domain)
			if err != nil {
				result[i][j] = nil
			}
			result[i][j] = registCheckResult
		}
	}

	return result
}
