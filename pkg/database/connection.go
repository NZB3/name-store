package database

import (
	"database/sql"
	"log"
)

type db struct {
	connection *sql.DB
}

var instance *db = nil

func (db *db) GetConnection() *sql.DB {
	return instance.connection
}

func NewConnection(connectionURI string) (*db, error) {
	if instance == nil {
		instance = &db{
			connection: new(connectionURI),
		}
	}

	return instance, nil
}

func new(connectionURI string) *sql.DB {
	conn, err := sql.Open("postgres", connectionURI)
	if err != nil {
		log.Panicf("Error with connection to database: %s", err)
	}
	return conn
}
