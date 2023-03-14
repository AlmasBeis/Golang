package structure

import (
	"database/sql"
	"fmt"
)

func (u *User) Register() (int64, error) {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/mydb")
	if err != nil {
		return 0, err
	}
	defer db.Close()

	// Insert the new user into the database
	var userID int64
	err = db.QueryRow("INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id", u.Username, u.Password, u.Email).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (u *User) Authenticate() (int64, error) {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/mydb")
	if err != nil {
		return 0, err
	}
	defer db.Close()

	// Retrieve the user from the database by username
	var userID int64
	err = db.QueryRow("SELECT id, password FROM users WHERE username = $1", u.Username).Scan(&userID, &u.Password)
	if err != nil {
		return 0, err
	}

	// Check if the password matches
	if u.Password != password {
		return 0, fmt.Errorf("incorrect password")
	}

	return userID, nil
}
