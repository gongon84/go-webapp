package api

import (
	"net/http"

	"example.com/go-webapp/cmd/config"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context, cfg config.Configs) error {
	return c.Render(http.StatusOK, "index.html", "ok")
}
