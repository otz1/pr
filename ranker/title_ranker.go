package ranker

import (
	"errors"
	"github.com/getsentry/sentry-go"
	"github.com/otz1/pr/entity"
	"github.com/schollz/closestmatch"
	"log"
	"strings"
)

// this pass will use fuzzy matching to see
// how close the query is to the results given to us
// we rank results higher if the title of the result given to
// us closely matches the query more than others.
//
// note that this comparison is done in lowercase
type titleRanker struct {
	basicRanker
}

func (t *titleRanker) Score(resultsToRank []entity.RankedSearchResult) []entity.RankedSearchResult {
	if t.context == nil {
		// do nothing no context.
		sentry.CaptureException(errors.New("trying to rank titles but no context given"))
		return resultsToRank
	}

	wordsToTest := make([]string, len(resultsToRank))
	for i, result := range resultsToRank {
		wordsToTest[i] = strings.ToLower(result.Result.Title)
	}

	cm := closestmatch.New(wordsToTest, []int{2})

	// the index is the ranking of the result
	ranking := cm.ClosestN(strings.ToLower(t.context.query), len(resultsToRank))
	log.Println("closes results are", ranking)

	for i, result := range resultsToRank {
		title := strings.ToLower(result.Result.Title)

		for idx, item := range ranking {
			if strings.Compare(title, item) == 0 {
				// convert a low ranking into a higher score.
				resultsToRank[i].Score += len(resultsToRank) - idx
				break
			}
		}
	}

	return resultsToRank
}

func newTitleRanker() *titleRanker {
	return &titleRanker{
		newBasicRanker(),
	}
}

