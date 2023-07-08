package database

import (
	"database/sql"
	"name-storage/database/models"
)

type DB interface {
	GetConnection() *sql.DB
	AllUsers() ([]models.User, error)
	AddUser(user *models.User) error
}
