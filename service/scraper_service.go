package service

import (
	"github.com/otz1/pr/client"
	"github.com/otz1/pr/entity"
)

type ScraperService struct {
	scraperClient *client.ScraperClient
}

func (s *ScraperService) Scrape(query string) *entity.ScrapeResponse {
	return s.scraperClient.Scrape(query)
}

func NewScraperService() *ScraperService {
	return &ScraperService{
		scraperClient: client.NewScraperClient(),
	}
}
