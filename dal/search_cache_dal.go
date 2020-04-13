package dal

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/go-redis/redis"
	"github.com/otz1/pr/conv"
	"github.com/otz1/pr/entity"
	"log"
	"os"
	"time"
)

type SearchResultCacheDAL struct {
	BasicResultCacheDAL
	client *redis.Client
}

func (s *SearchResultCacheDAL) Store(query string, resultSet *entity.SearchResultSet) error {
	jsonData, err := resultSet.ToJSON()
	if err == nil {
		s.client.Set(s.Hash(query), jsonData, time.Hour * 3)
		return nil
	}
	sentry.CaptureException(err)
	return fmt.Errorf("failed to cache search results for query '%s'", query)
}

// TODO should we distinguish nil from cache miss to genuine error
func (s *SearchResultCacheDAL) Query(query string) *entity.SearchResultSet {
	result, err := s.client.Get(s.Hash(query)).Result()
	if err != nil {
		sentry.CaptureException(err)
		return nil
	}
	return conv.JSONToSearchResultSet([]byte(result))
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
