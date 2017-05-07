package model

import (
	"app/shared/database"
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrAuthentication = errors.New("authentication failed")

type User struct {
	Id      int64
	Email   string
	Deleted uint8
}

func UserAuthenticate(email string, password string) (int64, error) {
	row := database.SQL.QueryRow(
		"SELECT id, password FROM user WHERE email = ? LIMIT 1",
		email)

	var id int64
	var hash string
	err := row.Scan(&id, &hash)
	if err == sql.ErrNoRows {
		return 0, ErrAuthentication
	} else if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return id, nil
	} else if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, ErrAuthentication
	}
	return 0, err
}

func UserCreate(email string, password string) (int64, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := database.SQL.Exec(
		"INSERT INTO user (email, password) VALUES (?, ?)",
		email, hash)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
