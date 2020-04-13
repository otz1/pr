package ranker

import "github.com/otz1/pr/entity"

// sourceRanker is a pass that will rank the results
// based on their source.
//
// at the moment the current way this works is that
// any result from a search source, e.g. yahoo or ddg gets
// ranked higher.
type sourceRanker struct {}

func sourceRank(source entity.ScrapeSource) int {
	switch source {
	case entity.DDG:
		return 1
	case entity.YAHOO:
		return 1
	case entity.BING:
		return 1
	default:
		return 0
	}
}

func (s *sourceRanker) Score(result entity.RankedSearchResult) entity.RankedSearchResult {
	newScore := result.Score
	newScore += sourceRank(result.Result.Source)
	return entity.RankedSearchResult{
		Result: result.Result,
		Score:  newScore,
	}
}

func newSourceRanker() *sourceRanker {
	return &sourceRanker{}
}

