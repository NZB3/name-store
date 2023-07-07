package database

import (
	"database/sql"
)

type DB struct {
	connection *sql.DB
}

func New(connectionURI string) (*DB, error) {
	db, err := sql.Open("postgres", connectionURI)
	if err != nil {
		return nil, err
	}
	return &DB{
		connection: db,
	}, nil
}
