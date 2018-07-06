package provider

import (
	"database/sql"
	"log"

	"github.com/spf13/viper"

	"github.com/govinda-attal/articles-api/internal/provider/pg"
)

const (
	// PrvDB is constant literal and is used as a key to store db connection within viper config data map.
	PrvDB = "prv.db"
)

// Setup function loads providers for this microservice. For articles-api postgres db is the only provider.
// This function is called at the microservice startup.
func Setup() {
	db, err := pg.InitStore()
	if err != nil {
		log.Fatal(err)
	}
	viper.SetDefault(PrvDB, db)
}

// DB function returns sql DB connection for this microservice.
func DB() *sql.DB {
	return viper.Get(PrvDB).(*sql.DB)
}

// Cleanup function cleans up active provider resources if any.
// This function is to be called when the microservice is shutting down.
func Cleanup() {
	db := DB()
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Println("Close on db connection failed!")
		}
	}
}
