package ranker

import (
	"github.com/otz1/pr/entity"
	"sort"
)

// contains the series of passes that are executed on each result
// that affect the score of the individual result
var scoringPasses = []resultRanker{
	newSourceRanker(),
	newTitleRanker(),
}

type ResultRanker struct{}

// buildRankedResultSet ...
func buildRankedResultSet(set []entity.ScrapeResult) []entity.RankedSearchResult {
	rankedResults := make([]entity.RankedSearchResult, len(set))
	for i, result := range set {
		rankedResults[i] = entity.RankedSearchResult{
			Result: result,
			Score:  0,
		}
	}
	return rankedResults
}

func (r *ResultRanker) Rank(query string, resultSet []entity.ScrapeResult) []entity.RankedSearchResult {
	resultsToRank := buildRankedResultSet(resultSet)

	for _, pass := range scoringPasses {
		pass.SetContext(query)
		resultsToRank = pass.Score(resultsToRank[:])
	}

	// sort by the result scores from highest to lowest.
	sort.Slice(resultsToRank[:], func(i int, j int) bool {
		return resultsToRank[i].Score > resultsToRank[j].Score
	})

	return resultsToRank
}

func scoreResult(result entity.RankedSearchResult) entity.RankedSearchResult {
	var scoredResult = result

	return scoredResult
}

func NewResultRanker() *ResultRanker {
	return &ResultRanker{}
}
