// database/database.go
package database

import (
	"database/sql"
	"user_management_service/models"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{db: db}
}

func (d *Database) CreateUser(user *models.User) error {
	_, err := d.db.Exec("INSERT INTO  users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}

func (d *Database) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	row := d.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return user, nil
}

func (d *Database) UpdateUser(user *models.User) error {
	_, err := d.db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3",
		user.Name, user.Email, user.ID)
	return err
}

func (d *Database) DeleteUser(id int) error {
	_, err := d.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
