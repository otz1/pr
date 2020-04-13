package ranker

import "github.com/otz1/pr/entity"

type resultRanker interface {
	Score(result entity.RankedSearchResult) entity.RankedSearchResult
}