package service

import "github.com/otz1/pr/dal"

type CacheService struct {
	cacheDAL dal.ResultCacheDAL
}

// TODO site code!
func (c *CacheService) Store(query string, results *dal.SearchResultSet) bool {
	err := c.cacheDAL.Store(query, results)
	return err == nil
}

func (c *CacheService) Contains(query string) bool {
	return c.cacheDAL.Query(query) == nil
}

// TODO site code!
func (c *CacheService) Query(query string) *dal.SearchResultSet {
	return c.cacheDAL.Query(query)
}

func NewCacheService() *CacheService {
	return &CacheService{
		cacheDAL: dal.NewSearchResultCacheDAL(),
	}
}
