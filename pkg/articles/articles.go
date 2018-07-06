package articles

import (
	"database/sql"
	"time"

	"github.com/lib/pq"

	"github.com/govinda-attal/articles-api/internal/provider"
	"github.com/govinda-attal/articles-api/pkg/core/status"
	"github.com/govinda-attal/articles-api/pkg/core/types"
)

// API interface lists methods for Articles microservice.
type API interface {
	Add(art *Article) (int, error)
	Get(id int) (*Article, error)
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

type articlesAPI struct {
	db *sql.DB
}

// Refer to Add method within API interface
func (api *articlesAPI) Add(art *Article) (string, error) {
	var id string
	db := api.db
	newStmt := `INSERT INTO ARTICLES.ARTICLE (title, publish_date, body, tags) VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(
		newStmt,
		art.Title,
		time.Time(*art.Date),
		art.Body,
		pq.Array(art.Tags)).Scan(&id)
	if err != nil {
		return id, status.ErrInternal.WithMessage(err.Error())
	}
	return id, err
}

// Refer to Get method within API interface
func (api *articlesAPI) Get(id int) (*Article, error) {
	db := api.db
	a := Article{}
	findStmt := `SELECT id, title, publish_date, body, tags FROM ARTICLES.ARTICLE WHERE id = $1`
	err := db.QueryRow(
		findStmt,
		id).Scan(&a.ID, &a.Title, &a.Date, &a.Body, pq.Array(&a.Tags))
	if err == sql.ErrNoRows {
		return nil, status.ErrNotFound.WithMessage(err.Error())
	}
	return &a, nil
}

// Refer to QueryTagSummaryForDay method within API interface
func (api *articlesAPI) QueryTagSummaryForDay(tag string, pd time.Time) (*TagSummary, error) {
	tags := []string{tag}
	tagSumy := &TagSummary{Tag: tag}

	qRelArtStmt := `SELECT ARRAY(
		SELECT  id FROM articles.article  WHERE $1 <@ tags AND publish_date = $2::date ORDER BY created_on DESC LIMIT 10
		)`
	qRelTagsStmt := `select ARTICLES.ARRAY_DISTINCT_MINUS(
		ARRAY(SELECT UNNEST(tags) FROM articles.article WHERE $1 <@ tags AND publish_date = $2::date ORDER BY created_on DESC LIMIT 10), $1
		)`

	db := api.db

	err := db.QueryRow(qRelArtStmt, pq.Array(tags), pd).Scan(pq.Array(&tagSumy.ArticleIDs))
	if err != nil || len(tagSumy.ArticleIDs) == 0 {
		return nil, status.ErrNotFound
	}
	err = db.QueryRow(qRelTagsStmt, pq.Array(tags), pd).Scan(pq.Array(&tagSumy.RelatedTags))
	if err != nil {
		return nil, status.ErrNotFound.WithMessage(err.Error())
	}
	tagSumy.Count = len(tagSumy.ArticleIDs)
	return tagSumy, nil
}

// NewAPI returns an implemenation of API interface.
// The type returned is intentionally unexported to restrict developers to create variable of this type.
func NewAPI() *articlesAPI {
	db := provider.DB()
	return &articlesAPI{db}
}
