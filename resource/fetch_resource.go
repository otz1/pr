package resource

import (
	"github.com/otz1/pr/dal"
	"github.com/otz1/pr/service"
	"log"
	"time"
)

type FetchResource struct {
	cacheService *service.CacheService
	scraperService *service.ScraperService
	rankerService *service.ResultRankerService
}

func (f *FetchResource) Fetch(query string) {
	// 1. we try and check the DB for the query searched.
	// bonus points: we search by nearest word distance type shit
	// bonus bonus points: we do some kind of keyword (input) => results (output) type thing
	resultSet := f.cacheService.Query(query)

	// 2. if not in the database, then we scrape for the search
	// 3. if it is in the results, we see how old they are. if they are super old
	//    then we scrape again (super old > 3 hours?)
	if resultSet == nil || time.Now().Sub(resultSet.Created) > time.Hour * 3 {
		scrapedResults := f.scraperService.Scrape(query)
		resultSet = dal.BuildSearchResultSet(scrapedResults.Results)
	}

	log.Println("no results for query", query)

	// 4. now we have some nice results to serve up to the user.
	//    these results are stored in the database.
	// we can do this asynchronously?
	go f.cacheService.Store(query, resultSet)

	// 5. then we rank the results (should we do this every time? probably not)
	f.rankerService.Rank(resultSet)

	// 6. finally these results are given to the user.
	log.Println(resultSet)
}

func NewFetchResource() FetchResource {
	return FetchResource{
		cacheService: service.NewCacheService(),
		scraperService: service.NewScraperService(),
		rankerService: service.NewResultRankerService(),
	}
}
