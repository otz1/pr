package ranker

type sourceRanker struct {}

func (s *sourceRanker) Score(result RankedResult) RankedResult {
	newScore := result.score

	// TODO do some stuff like decrease if wikipedia.

	return RankedResult{
		result.originalResult,
		newScore,
	}
}

func newSourceRanker() *sourceRanker {
	return &sourceRanker{}
}

