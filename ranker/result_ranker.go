package ranker

type resultRanker interface {
	Score(result RankedResult) RankedResult
}