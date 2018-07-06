package pg

import (
	"database/sql"
	"fmt"
	// To load pg driver.
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// InitStore initializes DB connection with settings within config file.
func InitStore() (*sql.DB, error) {
	dbConf := viper.GetStringMap("services.db")
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConf["host"].(string),
		dbConf["port"].(int),
		dbConf["username"].(string),
		dbConf["password"].(string),
		dbConf["dbname"].(string))
	db, err := sql.Open("postgres", connStr)
	return db, err
}
