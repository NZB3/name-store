package database

import (
	"log"
	"name-storage/database/models"
)

func (db *DB) PrepareUsers() error {
	const fn = "database.users.PrepareUsers: "

	stmt, err := db.connection.Prepare(`
		CREATE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255),
			email VARCHAR(255)
		);
		CREATE INDEX IF NOT EXISTS idx_email on users(email)
	`)
	if err != nil {
		log.Println(fn, err)
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(fn, err)
		return err
	}

	return nil
}

func (db *DB) AddUser(user *models.User) error {
	const fn = "database.users.AddUser: "

	query := db.connection.QueryRow(`
		INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id;
	`, user.Name, user.Email)

	err := query.Scan(&user.ID)
	if err != nil {
		log.Println(fn, err)
		return err
	}

	log.Println("User was successfully added to database")
	return nil
}

func (db *DB) AllUsers() ([]models.User, error) {
	const fn = "database.users.AllUsers: "

	rows, err := db.connection.Query(`
		SELECT * FROM users
	`)

	if err != nil {
		log.Println(fn, err)
		return nil, err
	}

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Println(fn, err)
			return nil, err
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		log.Println(fn, err)
		return nil, err
	}

	return users, nil
}
