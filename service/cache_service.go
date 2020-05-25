package service

import (
	"github.com/otz1/pr/dal"
	"github.com/otz1/pr/entity"
)

type CacheService struct {
	cacheDAL dal.ResultCacheDAL
}

func (c *CacheService) Store(query string, results *entity.SearchResultSet) bool {
	err := c.cacheDAL.Store(query, results)
	return err == nil
}

func (c *CacheService) Contains(query string) bool {
	return c.cacheDAL.Query(query) == nil
}

func (c *CacheService) Query(query string) *entity.SearchResultSet {
	return c.cacheDAL.Query(query)
}

func NewCacheService() *CacheService {
	return &CacheService{
		cacheDAL: dal.NewSearchResultCacheDAL(),
	}
}
