package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplateRenderer() *TemplateRenderer {
	return &TemplateRenderer{templates: template.Must(template.ParseGlob("views/*.templ"))}
}

type App struct {
	Title string
	Name  string
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplateRenderer()

	app := &App{
		Title: "Conway's The Game of Life",
		Name:  "World",
	}

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", app)
	})

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":8386"))
}
