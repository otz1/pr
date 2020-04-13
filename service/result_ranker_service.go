package service

import (
	"github.com/otz1/pr/entity"
	"github.com/otz1/pr/ranker"
)

type ResultRankerService struct {
	resultRanker *ranker.ResultRanker
}

func (r *ResultRankerService) Rank(query string, resultSet []entity.ScrapeResult) []entity.RankedSearchResult {
	return r.resultRanker.Rank(query, resultSet)
}

func NewResultRankerService() *ResultRankerService {
	return &ResultRankerService{
		resultRanker: ranker.NewResultRanker(),
	}
}
