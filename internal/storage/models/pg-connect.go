package models

import "database/sql"

type Datastore interface {
	AllNames() ([]*User, error)
}

type DB struct {
	*sql.DB
}

func Connect(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func Close(db *DB) error {
	err := db.Close()
	return err
}
