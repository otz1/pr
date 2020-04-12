package dal

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/go-redis/redis"
	"log"
	"os"
	"time"
)

type SearchResultCacheDAL struct {
	BasicResultCacheDAL
	client *redis.Client
}

func (s *SearchResultCacheDAL) Store(query string, resultSet *SearchResultSet) error {
	jsonData, err := resultSet.ToJSON()
	if err == nil {
		s.client.Set(s.Hash(query), jsonData, time.Hour * 3)
		return nil
	}
	sentry.CaptureException(err)
	return fmt.Errorf("failed to cache search results for query '%s'", query)
}

// TODO should we distinguish nil from cache miss to genuine error
func (s *SearchResultCacheDAL) Query(query string) *SearchResultSet {
	result, err := s.client.Get(s.Hash(query)).Result()
	if err != nil {
		sentry.CaptureException(err)
		return nil
	}
	return jsonToSearchResultSet([]byte(result))
}

func NewSearchResultCacheDAL() *SearchResultCacheDAL {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URI"))
	if err != nil {
		sentry.CaptureException(err)
		log.Fatal(err)
	}
	return &SearchResultCacheDAL{
		NewBasicResultCacheDAL(),
		redis.NewClient(opt),
	}
}
