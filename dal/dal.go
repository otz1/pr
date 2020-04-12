package dal

import (
	"fmt"
)

type ResultCacheDAL interface {
	// Query will look in the cache for the results
	// stored of the given query
	Query(query string) *SearchResultSet

	// Store will store the results for the given query
	// in the cache.
	Store(query string, results *SearchResultSet) error

	// Hash will hash the given query into a key
	// to index the cache with only relevant for the redis
	// cache at the moment however
	Hash(query string) string
}

type BasicResultCacheDAL struct {}

func NewBasicResultCacheDAL() BasicResultCacheDAL {
	return BasicResultCacheDAL{}
}

func (b BasicResultCacheDAL) Hash(query string) string {
	// perhaps we should encode the searches somehow
	// into a string.
	return fmt.Sprintf("search:%s")
}