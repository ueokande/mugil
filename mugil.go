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

	db, err := database.LoadConfig("database.json")
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

	e.GET("/", controller.TaskIndex)
	e.POST("/tasks", controller.TaskCreate)
	e.Logger.Fatal(e.Start(":1323"))
}
