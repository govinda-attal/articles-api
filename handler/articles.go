package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/govinda-attal/articles-api/pkg/articles"
	"github.com/govinda-attal/articles-api/pkg/core/status"
)

// AddArticleHandler on HTTP POST request for article data in json format and stores it within the system.
func AddArticleHandler(w http.ResponseWriter, r *http.Request) error {
	article := articles.Article{}
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		return status.ErrBadRequest.WithMessage(err.Error())
	}
	api := articles.NewAPI()
	id, err := api.Add(&article)
	if err != nil {
		return err
	}
	article.ID = id
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
	return nil
}

// FetchArticleHandler on HTTP GET request for a given article {id} and returns back its JSON representation.
func FetchArticleHandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	val, ok := vars["id"]
	if !ok {
		return status.ErrBadRequest
	}
	id, err := strconv.Atoi(val)
	if err != nil {
		return status.ErrBadRequest
	}
	api := articles.NewAPI()
	article, err := api.Get(id)
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
	return nil
}

// FetchArticleTagSummaryHandler on HTTP GET request and returns the list of articles that have that tag name on the given date.
// It also returns some summary data about that tag for that day.
func FetchArticleTagSummaryHandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	tag := vars["tagName"]
	pdVar := vars["date"]
	pd, err := time.Parse("20060102", pdVar)
	if err != nil {
		return status.ErrBadRequest
	}
	api := articles.NewAPI()
	tagSumy, err := api.QueryTagSummaryForDay(tag, pd)
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tagSumy)
	return nil
}
