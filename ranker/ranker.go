package ranker

import (
	"github.com/otz1/pr/entity"
	"sort"
)

// contains the series of passes that are executed on each result
// that affect the score of the individual result
var scoringPasses = []resultRanker{
	newSourceRanker(),
}

type ResultRanker struct{}

func (r *ResultRanker) Rank(resultSet []entity.ScrapeResult) []entity.RankedSearchResult {
	results := make([]entity.RankedSearchResult, len(resultSet))
	for idx, result := range resultSet {
		results[idx] = scoreResult(entity.RankedSearchResult{
			Result: result,
			Score:          0,
		})
	}

	// sort by the result scores
	sort.Slice(results[:], func(i int, j int) bool {
		return results[i].Score > results[j].Score
	})

	return results
}

func scoreResult(result entity.RankedSearchResult) entity.RankedSearchResult {
	var scoredResult = result
	for _, pass := range scoringPasses {
		scoredResult = pass.Score(result)
	}
	return scoredResult
}

func NewResultRanker() *ResultRanker {
	return &ResultRanker{}
}
