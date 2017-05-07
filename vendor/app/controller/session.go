package controller

import (
	"app/shared/session"

	"github.com/labstack/echo"
)

func CurrentUserId(c echo.Context) (int64, error) {
	s, err := session.DefaultSessionManager().StartSession(c)
	if err != nil {
		return 0, err
	}

	id, err := s.Get("current_user_id")
	if err != nil {
		return 0, err
	}
	return id.(int64), nil
}

func LoggedIn(c echo.Context) error {
	_, err := CurrentUserId(c)
	return err
}
