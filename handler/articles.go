package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/govinda-attal/articles-api/pkg/articles"
	"github.com/govinda-attal/articles-api/pkg/core/status"
)

// ArticlesHandler implements methods to handle HTTP requests for the microservice.
type ArticlesHandler struct {
	artAPI articles.API
}
// NewArticleHandler returns a new HTTP handler instance for articles API
func NewArticleHandler(artAPI articles.API) (*ArticlesHandler){
	return &ArticlesHandler {artAPI}
}

// AddArticle on HTTP POST request for article JSON data stores it within the system.
func (ah *ArticlesHandler) AddArticle(w http.ResponseWriter, r *http.Request) error {
	article := articles.Article{}
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		return status.ErrBadRequest.WithMessage(err.Error())
	}
	id, err := ah.artAPI.Add(&article)
	if err != nil {
		return err
	}
	article.ID = id
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
	return nil
}

// FetchArticle on HTTP GET request for a given article {id} and returns back its JSON representation.
func (ah *ArticlesHandler)FetchArticle(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return status.ErrBadRequest
	}
	article, err := ah.artAPI.Get(id)
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
	return nil
}

// FetchArticleTagSummary on HTTP GET request and returns the list of articles that have that tag name on the given date.
// It also returns some summary data about that tag for that day.
func (ah *ArticlesHandler)FetchArticleTagSummary(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	tag := vars["tagName"]
	pdVar := vars["date"]
	pd, err := time.Parse("20060102", pdVar)
	if err != nil {
		return status.ErrBadRequest
	}
	tagSumy, err := ah.artAPI.QueryTagSummaryForDay(tag, pd)
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tagSumy)
	return nil
}
