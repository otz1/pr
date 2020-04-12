package entity

type ScrapeRequest struct {
	Query string `json:"query"`
}

type ScrapeResult struct {
	Title   string `json:"title"`
	Href    string `json:"href"`
	Snippet string `json:"snippet"`
}

type ScrapeResponse struct {
	OriginalQuery string         `json:"originalQuery"`
	Results       []ScrapeResult `json:"results"`
}