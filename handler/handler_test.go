package handler_test

import (
	"encoding/json"
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/govinda-attal/articles-api/handler"
	"github.com/govinda-attal/articles-api/test/mocks"
)

func TestHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Suite")
}

var _ = Describe("Article Handler's", func() {

	Describe("Basic Behaviour", func() {
		var (
			addArticleRqB = []byte(`{"title": "hello-world","date": "2006-01-02","body": "hello-world-body","tags": ["hello-tag","world-tag"]}`)
			articleB = []byte(`{"id": "1","title": "hello-world","date": "2006-01-02","body": "hello-world-body","tags": ["hello-tag","world-tag"]}`)
			fetchArticleTagSummaryRsB = []byte(`{"tag": "hello-tag","count": 1,"articles": ["1"],"related_tags": ["world-tag"]}`)
			
			testRouter = func (path string, handler http.HandlerFunc) *mux.Router {
				r := mux.NewRouter()
				r.HandleFunc(path, handler)
				return r 
			}
		)
		BeforeEach(func() {

		})
		Context("When adding a valid article", func() {
			mockADS := &mocks.ArticleMockDS{}
			mockADS.AddCall.Returns.ArticleID = "1"
			It("must return new article identifier and no error", func() {
				rr := httptest.NewRecorder()
				ah := NewArticleHandler(mockADS)
				handler := WrapperHandler(ah.AddArticle)
				req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(addArticleRqB))
				router := testRouter("/articles", handler)
				router.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusOK))
				respBytes, err := ioutil.ReadAll(rr.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(respBytes).To(MatchJSON(articleB))
			})
		})

		Context("When fetching an article with given valid identifier", func() {
			mockADS := &mocks.ArticleMockDS{}
			json.Unmarshal(articleB, &mockADS.GetCall.Returns.Article)
			It("must return article information and no error", func() {
				rr := httptest.NewRecorder()
				ah := NewArticleHandler(mockADS)
				handler := WrapperHandler(ah.FetchArticle)
				req, _ := http.NewRequest("GET", "/articles/1", nil)
				router := testRouter("/articles/{id:[0-9]+}", handler)
				router.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusOK))
				respBytes, err := ioutil.ReadAll(rr.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(respBytes).To(MatchJSON(articleB))
				
			})
		})

		Context("When fetching tag summary with given valid tag and publish date", func() {
			mockADS := &mocks.ArticleMockDS{}
			json.Unmarshal(fetchArticleTagSummaryRsB, &mockADS.QueryTagSummaryForDayCall.Returns.TagSummary)
			It("must return tag summary information and no error", func() {
				rr := httptest.NewRecorder()
				ah := NewArticleHandler(mockADS)
				handler := WrapperHandler(ah.FetchArticleTagSummary)
				req, _ := http.NewRequest("GET", "/tags/hello-tag/20060102", nil)
				router := testRouter("/tags/{tagName}/{date:\\d{4,4}\\d{2,2}\\d{2,2}}", handler)
				router.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusOK))
				respBytes, err := ioutil.ReadAll(rr.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(respBytes).To(MatchJSON(fetchArticleTagSummaryRsB))				
			})
		})
	})
})
