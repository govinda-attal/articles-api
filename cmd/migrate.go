// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	// Register Step: to migrate DB up/down
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/spf13/cobra"

	"github.com/govinda-attal/articles-api/internal/provider"
)

const (
	up   = "up"
	down = "down"
)

var migSrc string

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "reads migrations from sources and applies them to the database",
	Run:   migrateDB,
}

func migrateDB(cmd *cobra.Command, args []string) {
	provider.Setup()
	defer provider.Cleanup()
	db := provider.DB()

	action := strings.Join(args, "")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	srcURL := fmt.Sprintf("file://%s/%s/", dir, migSrc)
	if strings.HasPrefix(migSrc, "/") {
		srcURL = fmt.Sprintf("file://%s", migSrc)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(srcURL, "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	switch action {
	case up:
		err = m.Up()
	case down:
		err = m.Down()
	}
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.PersistentFlags().StringVar(&migSrc, "src", "migrations", "migrations source directory (migrations)")
}
