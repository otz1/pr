package ranker

import "github.com/otz1/pr/entity"

type sourceRanker struct {}

func (s *sourceRanker) Score(result entity.RankedSearchResult) entity.RankedSearchResult {
	newScore := result.Score

	// TODO do some stuff like decrease if wikipedia.

	return entity.RankedSearchResult{
		Result: result.Result,
		Score:  newScore,
	}
}

func newSourceRanker() *sourceRanker {
	return &sourceRanker{}
}

