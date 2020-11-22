package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// Database contains objects for database communication.
type Database struct {
	*gorm.DB
}

// New create new database instance.
func New(config *Config) (*Database, error) {
	db, err := gorm.Open("postgres", config.DatabaseURI)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}
	return &Database{db}, nil
}
