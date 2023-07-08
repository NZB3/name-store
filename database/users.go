package database

import (
	"log"
	"name-storage/database/models"
)

func (db *db) prepareUsers() error {
	const fn = "database.users.PrepareUsers: "

	stmt, err := db.connection.Prepare(`
		CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			name VARCHAR(255),
			email VARCHAR(255)
		);
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

func (db *db) AddUser(user *models.User) error {
	const fn = "database.users.AddUser: "

	err := db.prepareUsers()
	if err != nil {
		log.Println(fn, err)
		return err
	}

	query := db.connection.QueryRow(`
		INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id;
	`, user.Name, user.Email)

	err = query.Scan(&user.ID)
	if err != nil {
		log.Println(fn, err)
		return err
	}

	log.Println("User was successfully added to database")
	return nil
}

func (db *db) AllUsers() ([]models.User, error) {
	const fn = "database.users.AllUsers: "

	err := db.prepareUsers()
	if err != nil {
		log.Println(fn, err)
		return nil, err
	}

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
