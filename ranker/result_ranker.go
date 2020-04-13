package ranker

import "github.com/otz1/pr/entity"

type rankerContext struct {
	query string
}

type resultRanker interface {
	Context() *rankerContext
	SetContext(query string)
	Score(result []entity.RankedSearchResult) []entity.RankedSearchResult
}

type basicRanker struct {
	context *rankerContext
}

func newBasicRanker() basicRanker {
	return basicRanker{
		context: nil,
	}
}

func (b *basicRanker) SetContext(query string) {
	b.context = &rankerContext{
		query:     query,
	}
}

func (b *basicRanker) Context() *rankerContext {
	return b.context
}