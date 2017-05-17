package controller

import (
	"app/model"
	"app/shared/session"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func LoginGet(c echo.Context) error {
	err := LoggedIn(c)
	if err == nil {
		return c.Redirect(http.StatusFound, "/")
	} else if err != session.ErrNoSuchValue {
		log.Error(err)
		return err
	}

	return c.Render(http.StatusOK, "login.html", nil)
}

type LoginPostForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginPost(c echo.Context) error {
	var form LoginPostForm
	err := c.Bind(&form)
	if err != nil {
		log.Error(err)
		return err
	}
	id, err := model.UserAuthenticate(form.Email, form.Password)
	if err == model.ErrAuthentication {
		return c.JSON(http.StatusForbidden, MessageJsonDto{"authentication error"})
	} else if err != nil {
		log.Error(err)
		return err
	}

	sess, err := session.DefaultSessionManager().StartSession(c)
	if err != nil {
		log.Error(err)
		return err
	}
	sess.Set("current_user_id", id)

	return c.JSON(http.StatusOK, MessageJsonDto{"OK"})
}
