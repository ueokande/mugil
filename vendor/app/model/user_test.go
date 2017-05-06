package model

import (
	"app/shared/database"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func createUser(email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = database.SQL.Exec(
		"INSERT INTO user (email, password) VALUES (?, ?)",
		email, hash)
	if err != nil {
		return err
	}
	return nil
}

func TestUserAuthenticate(t *testing.T) {
	err := createUser("abc@example.com", "secret")
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		email    string
		password string
		expected bool
	}{
		{"abc@example.com", "secret", true},
		{"abc@example.com", "non-secret", false},
		{"nonuser@example.com", "non-secret", false},
	}

	for _, c := range cases {
		ok, err := UserAuthenticate(c.email, c.password)
		if err != nil {
			t.Fatal(err)
		}
		if ok != c.expected {
			t.Errorf("Unexpected authentication (email=%v, password=%v)", c.email, c.password)
		}
	}

}
