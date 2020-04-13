package conv

import (
	"github.com/getsentry/sentry-go"
	jsoniter "github.com/json-iterator/go"
	"github.com/otz1/pr/entity"
)

// JSONToSearchResultSet will convert json data from the cache
// into an object we can work with
func JSONToSearchResultSet(jsonData []byte) *entity.SearchResultSet {
	result := &entity.SearchResultSet{}
	if err := jsoniter.Unmarshal(jsonData, result); err != nil {
		sentry.CaptureException(err)
		return nil
	}
	return result
}
