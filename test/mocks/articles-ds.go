package mocks

import (
	"time"

	"github.com/govinda-attal/articles-api/pkg/articles"
)

// ArticleMockDS is an mock article datastore.
type ArticleMockDS struct {
	AddCall struct {
		Receives struct {
			Article *articles.Article
		}
		Returns struct {
			ArticleID string
			Error     error
		}
	}

	GetCall struct {
		Receives struct {
			ArticleID string
		}
		Returns struct {
			Article *articles.Article
			Error   error
		}
	}

	QueryTagSummaryForDayCall struct {
		Receives struct {
			Tag         string
			PublishDate time.Time
		}
		Returns struct {
			TagSummary *articles.TagSummary
			Error      error
		}
	}
}

// Add Mock.
func (am *ArticleMockDS) Add(art *articles.Article) (string, error) {
	am.AddCall.Receives.Article = art
	return am.AddCall.Returns.ArticleID, am.AddCall.Returns.Error
}

// Get Mock.
func (am *ArticleMockDS) Get(idS string) (*articles.Article, error) {
	am.GetCall.Receives.ArticleID = idS
	return am.GetCall.Returns.Article, am.GetCall.Returns.Error
}

// QueryTagSummaryForDay Mock.
func (am *ArticleMockDS) QueryTagSummaryForDay(tag string, pd time.Time) (*articles.TagSummary, error) {
	am.QueryTagSummaryForDayCall.Receives.Tag = tag
	am.QueryTagSummaryForDayCall.Receives.PublishDate = pd
	return am.QueryTagSummaryForDayCall.Returns.TagSummary, am.QueryTagSummaryForDayCall.Returns.Error
}
