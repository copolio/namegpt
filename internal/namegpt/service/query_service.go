package service

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/internal/namegpt/repository"
	"github.com/copolio/namegpt/pkg/client/chatgpt"
	"github.com/copolio/namegpt/pkg/dto/request"
)

type QueryUseCase interface {
	Handle(request request.SimilarDomainNames) (domainNames []entity.DomainName, err error)
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

func (q QueryService) Handle(request request.SimilarDomainNames) (domainNames []entity.DomainName, err error) {
	db := config.GetGormDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Find Or Create Query
	q.queryRepository.WithTransaction(tx)
	query, err1 := q.queryRepository.FindOrCreate(entity.Query{Keyword: request.Keyword})
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
		return query.DomainNames, tx.Commit().Error
	}

	// Ask ChatGPT for domains
	domains, err3 := chatgpt.GetSimilarDomains(request.Keyword)
	if err3 != nil {
		tx.Rollback()
		return nil, err3
	}

	q.domainNameRepository.WithTransaction(tx)
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

	return domainNames, tx.Commit().Error
}
