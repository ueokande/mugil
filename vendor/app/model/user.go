package model

import (
	"app/shared/database"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id      uint64
	Email   string
	Deleted uint8
}

func UserAuthenticate(email string, password string) (bool, error) {
	row := database.SQL.QueryRow(
		"SELECT password FROM user WHERE email = ? LIMIT 1",
		email)

	var hash string
	err := row.Scan(&hash)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return true, nil
	} else if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}
	return false, err
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
