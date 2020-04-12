package ranker

import (
	"github.com/otz1/pr/dal"
	"github.com/otz1/pr/entity"
)

// contains the series of passes that are executed on each result
// that affect the score of the individual result
var scoringPasses = []resultRanker{
	newSourceRanker(),
}

type RankedResult struct {
	originalResult entity.ScrapeResult
	score          int
}

type ResultRanker struct{}

func (r *ResultRanker) Rank(resultSet *dal.SearchResultSet) {
	for _, result := range resultSet.Results {
		scoreResult(RankedResult{
			originalResult: result,
			score:          0,
		})
	}
}

func scoreResult(result RankedResult) RankedResult {
	var scoredResult = result
	for _, pass := range scoringPasses {
		scoredResult = pass.Score(result)
	}
	return scoredResult
}

func NewResultRanker() *ResultRanker {
	return &ResultRanker{}
}
