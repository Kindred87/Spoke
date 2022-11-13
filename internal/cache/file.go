package cache

import (
	"fmt"
	"os"

	"github.com/Kindred87/quickbolt"
)

var (
	db quickbolt.DB
)

// createDB creates a new database and assigns it to the package db variable.
func createDB() error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error while resolving working directory: %w", err)
	}

	d, err := quickbolt.Create("spoke.db", dir)
	if err != nil {
		return fmt.Errorf("error while creating database at %s: %w", dir, err)
	}

	db = d

	return nil
}

// Delete deletes the on-disk cache.
func Delete() {
	if db != nil {
		db.RemoveFile()
	}
}
