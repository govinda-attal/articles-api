package ds

import (
	"database/sql"
	"time"
	"strconv"

	"github.com/lib/pq"

	"github.com/govinda-attal/articles-api/internal/provider"
	"github.com/govinda-attal/articles-api/pkg/articles"
	"github.com/govinda-attal/articles-api/pkg/core/status"
)

// ArticleDS is an internal concrete implementation of articles.API interface.
type ArticleDS struct {
	db *sql.DB
}

// NewArticleDS returns an instance of ArticleDS.
func NewArticleDS() *ArticleDS {
	db := provider.DB()
	return &ArticleDS{db}
}

// Add method saves a new article in datastore.
func (ads *ArticleDS) Add(art *articles.Article) (string, error) {
	var id string
	db := ads.db
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

// Get method gets information for a give article.
func (ads *ArticleDS) Get(idS string) (*articles.Article, error) {
	id, err := strconv.Atoi(idS)
	if  err != nil {
		return nil, status.ErrBadRequest.WithMessage(err.Error())
	}
	db := ads.db
	a := articles.Article{}
	findStmt := `SELECT id, title, publish_date, body, tags FROM ARTICLES.ARTICLE WHERE id = $1`
	err = db.QueryRow(
		findStmt,
		id).Scan(&a.ID, &a.Title, &a.Date, &a.Body, pq.Array(&a.Tags))
	if err == sql.ErrNoRows {
		return nil, status.ErrNotFound.WithMessage(err.Error())
	}
	return &a, nil
}

// QueryTagSummaryForDay method returns a summary of articles tagged with a given tag name and published on give date.
func (ads *ArticleDS) QueryTagSummaryForDay(tag string, pd time.Time) (*articles.TagSummary, error) {
	tags := []string{tag}
	tagSumy := &articles.TagSummary{Tag: tag}

	qRelArtStmt := `SELECT ARRAY(
		SELECT  id FROM articles.article  WHERE $1 <@ tags AND publish_date = $2::date ORDER BY created_on DESC LIMIT 10
		)`
	qRelTagsStmt := `select ARTICLES.ARRAY_DISTINCT_MINUS(
		ARRAY(SELECT UNNEST(tags) FROM articles.article WHERE $1 <@ tags AND publish_date = $2::date ORDER BY created_on DESC LIMIT 10), $1
		)`

	db := ads.db

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
