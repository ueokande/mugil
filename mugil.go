package main

import (
	"app/controller"
	"app/shared/database"
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	*template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:CSRF_TOKEN",
	}))

	db, err := database.LoadConfig("config/development.json")
	if err != nil {
		e.Logger.Fatal(err)
	}
	err = database.Connect(db)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Renderer = &Template{
		Template: template.Must(template.ParseGlob("template/*.html")),
	}

	e.Static("/assets", "build")

	e.GET("/login", controller.LoginGet)
	e.POST("/login", controller.LoginPost)

	e.GET("/", controller.TaskIndex)
	e.GET("/tasks.json", controller.TaskIndexJson)
	e.POST("/tasks", controller.TaskCreate)
	e.Logger.Fatal(e.Start(":1323"))
}
