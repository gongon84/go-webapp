package main

import (
	"html/template"
	"io"

	"example.com/go-webapp/cmd/config"
	"example.com/go-webapp/cmd/controllers/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	// Template
	t := &Template{
		templates: template.Must(template.ParseGlob("html/*.html")),
	}
	e.Renderer = t

	e.GET("/", makeHandler(api.Index))
	e.GET("/company", makeHandler(api.CompanyDetail))
	e.Logger.Fatal(e.Start(":8000"))

}

func makeHandler(fn func(c echo.Context, cfg config.Configs) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		return fn(c, config.ApplyConfig())
	}
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
