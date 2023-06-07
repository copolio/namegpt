package service

import (
	"github.com/copolio/namegpt/config"
	"github.com/copolio/namegpt/internal/namegpt/entity"
	"github.com/copolio/namegpt/internal/namegpt/repository"
	"github.com/copolio/namegpt/pkg/client/chatgpt"
	"github.com/copolio/namegpt/pkg/dto/request"
	"gorm.io/gorm"
)

type QueryUseCase interface {
	Handle(request request.SearchDomainNames) (domainNames []entity.DomainName, err error)
}

func NewQueryUseCase() QueryUseCase {
	return &QueryService{
		queryRepository:        repository.NewQueryRepository(),
		queryHistoryRepository: repository.NewQueryHistoryRepository(),
	}
}

type QueryService struct {
	queryRepository        repository.QueryRepository
	queryHistoryRepository repository.QueryHistoryRepository
}

func (q QueryService) Handle(request request.SearchDomainNames) (domainNames []entity.DomainName, err error) {
	db := config.GetGormDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		// Find Or Create Query
		query, err1 := q.queryRepository.FindOrCreate(tx, &entity.Query{Keyword: request.Keyword})
		if err1 != nil {
			return err1
		}

		// Create query history
		_, err2 := q.queryHistoryRepository.Save(tx, &entity.QueryHistory{
			QueryID: query.ID,
			Query:   entity.Query{},
			UserId:  request.UserId,
		})
		if err2 != nil {
			return err2
		}

		// Return cached result if exists
		if len(query.DomainNames) > 0 {
			domainNames = query.DomainNames
			return nil
		}

		domains, err3 := chatgpt.GetSimilarDomains(request.Keyword)
		if err3 != nil {
			return err3
		}

		// Cache and return ChatGPT response
		err5 := tx.Transaction(func(tx2 *gorm.DB) error {
			for _, domain := range domains {
				domainName := &entity.DomainName{
					Name:    domain,
					QueryID: query.ID,
				}
				saveResult := tx2.Save(domainName)
				if saveResult.Error != nil {
					return saveResult.Error
				}
				domainNames = append(domainNames, *domainName)
			}
			return nil
		})
		if err5 != nil {
			return err5
		}
		return nil
	})
	return domainNames, err
}
