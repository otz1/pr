package entity

type SiteCode string

const (
	OTZIT_UK SiteCode = "UK"
	OTZIT_US          = "US"
	OTZIT_FR          = "FR"
	OTZIT_IT          = "IT"
	OTZIT_ES          = "ES"
)

type ScrapeRequest struct {
	Query string `json:"query"`
}

// ScrapeSource is a source to scrape from
// e.g. wikipedia.
type ScrapeSource string

// The sources available to scrape
// that are currently supported
const (
	DDG       ScrapeSource = "DDG"
	STARTPAGE              = "STARTPAGE"
	YAHOO                  = "YAHOO"
	BING                   = "BING"
	WIKIPEDIA              = "WIKIPEDIA"
)

type ScrapeResult struct {
	Title   string       `json:"title"`
	Href    string       `json:"href"`
	Snippet string       `json:"snippet"`
	Source  ScrapeSource `json:"source"`
}

type ScrapeResponse struct {
	OriginalQuery string         `json:"originalQuery"`
	Results       []ScrapeResult `json:"results"`
}
