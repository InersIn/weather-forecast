package controllers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	fmt.Println("TEST")
	return c.File("index.html")
}
