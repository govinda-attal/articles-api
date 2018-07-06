package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/urfave/negroni"

	"github.com/govinda-attal/articles-api/handler"
	"github.com/govinda-attal/articles-api/internal/ds"
	"github.com/govinda-attal/articles-api/internal/provider"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the articles api micro service",
	Run:   startServer,
}

func startServer(cmd *cobra.Command, args []string) {	
	provider.Setup()
	ads := ds.NewArticleDS()
	artHandler := handler.NewArticleHandler(ads)

	r := mux.NewRouter()

	r.PathPrefix("/api/").Handler(http.StripPrefix("/api", http.FileServer(http.Dir("./api"))))

	r.HandleFunc("/articles", handler.WrapperHandler(artHandler.AddArticle)).Methods("POST")
	r.HandleFunc("/articles/{id:[0-9]+}", handler.WrapperHandler(artHandler.FetchArticle)).Methods("GET")
	r.HandleFunc("/tags/{tagName}/{date:\\d{4,4}\\d{2,2}\\d{2,2}}", handler.WrapperHandler(artHandler.FetchArticleTagSummary)).Methods("GET")

	h := cors.Default().Handler(r)
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(h)

	srv := &http.Server{
		Addr:         "0.0.0.0:9080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      n,
	}

	go func() {
		defer provider.Cleanup()
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("articles server shutdown ...")
	os.Exit(0)
}

func init() {
	rootCmd.AddCommand(startCmd)
}
