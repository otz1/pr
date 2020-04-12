package dal

import (
	"errors"
	"github.com/getsentry/sentry-go"
	jsoniter "github.com/json-iterator/go"
	"github.com/otz1/pr/entity"
	"time"
)

type SearchResultSet struct {
	Results []entity.ScrapeResult `json:"results"`
	Created time.Time             `json:"created"`
}

func (s *SearchResultSet) ToJSON() (string, error) {
	jsonData, err := jsoniter.Marshal(s)
	if err != nil {
		sentry.CaptureException(err)
		return "", errors.New("failed to marshal into json object")
	}
	return string(jsonData), nil
}

// jsonToSearchResultSet will convert json data from the cache
// into an object we can work with
func jsonToSearchResultSet(jsonData []byte) *SearchResultSet {
	result := &SearchResultSet{}
	if err := jsoniter.Unmarshal(jsonData, result); err != nil {
		sentry.CaptureException(err)
		return nil
	}
	return result
}

func BuildSearchResultSet(results []entity.ScrapeResult) *SearchResultSet {
	return &SearchResultSet{
		results,
		time.Now(),
	}
}
