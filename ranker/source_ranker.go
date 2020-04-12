package ranker

type sourceRanker struct {}

func (s *sourceRanker) Score(result RankedResult) RankedResult {
	return RankedResult{}
}

func newSourceRanker() *sourceRanker {
	return &sourceRanker{}
}

