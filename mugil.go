package main

import (
	"app/controller"
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type Template struct {
	*template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Renderer = &Template{
		Template: template.Must(template.ParseGlob("template/*.html")),
	}

	e.GET("/", controller.TaskIndex)
	e.Logger.Fatal(e.Start(":1323"))
}
