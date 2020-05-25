package entity

import (
	"errors"
	"time"

	"github.com/getsentry/sentry-go"
	jsoniter "github.com/json-iterator/go"
)

type RankedSearchResult struct {
	Result ScrapeResult `json:"result"`
	Score  int          `json:"score"`
}

type SearchResultSet struct {
	Results []RankedSearchResult `json:"rankedResults"`
	Created time.Time            `json:"created"`
}

type PageRankRequest struct {
	Query string `json:"query"`
}

type PageRankResponse struct {
	Results []RankedSearchResult `json:"results"`
}

func (s *SearchResultSet) ToJSON() (string, error) {
	jsonData, err := jsoniter.Marshal(s)
	if err != nil {
		sentry.CaptureException(err)
		return "", errors.New("failed to marshal into json object")
	}
	return string(jsonData), nil
}

func BuildSearchResultSet(results []RankedSearchResult) *SearchResultSet {
	return &SearchResultSet{
		results,
		time.Now(),
	}
}
