package controller

import (
	"app/model"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func LoginGet(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

type LoginPostForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func LoginPost(c echo.Context) error {
	var form LoginPostForm
	err := c.Bind(&form)
	if err != nil {
		log.Error(err)
		return err
	}
	ok, err := model.UserAuthenticate(form.Email, form.Password)
	if err != nil {
		log.Error(err)
		return err
	}
	if !ok {
		c.Redirect(http.StatusFound, "/login")
	}
	return c.String(http.StatusOK, "OK")
}
