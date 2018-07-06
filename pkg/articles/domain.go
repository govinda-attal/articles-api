package articles

import (
	"time"

	"github.com/govinda-attal/articles-api/pkg/core/types"
)

// API interface lists methods for Articles microservice.
type API interface {
	// Add method saves a new article in datastore.
	Add(art *Article) (string, error)
	// Get method gets information for a give article.
	Get(id string) (*Article, error)
	// QueryTagSummaryForDay method returns a summary of articles tagged with a given tag name and published on give date.
	QueryTagSummaryForDay(tag string, date time.Time) (*TagSummary, error)
}

// Article represents a set of attributes for an article managed by ArticlesAPI microservice.
type Article struct {
	ID    string      `json:"id,omitempty"`
	Title string      `json:"title"`
	Date  *types.Date `json:"date"`
	Body  string      `json:"body"`
	Tags  []string    `json:"tags,omitempty"`
}

// TagSummary represents a set of attributes for an result set returned by Articles API microservice on query by a tag for a given date.
type TagSummary struct {
	Tag         string   `json:"tag"`
	Count       int      `json:"count"`
	ArticleIDs  []string `json:"articles"`
	RelatedTags []string `json:"related_tags,omitempty"`
}
