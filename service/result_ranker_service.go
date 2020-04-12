package service

import (
	"github.com/otz1/pr/dal"
	"github.com/otz1/pr/ranker"
)

type ResultRankerService struct {
	resultRanker *ranker.ResultRanker
}

func (r *ResultRankerService) Rank(resultSet *dal.SearchResultSet) {
	r.resultRanker.Rank(resultSet)
}

func NewResultRankerService() *ResultRankerService {
	return &ResultRankerService{
		resultRanker: ranker.NewResultRanker(),
	}
}
