package main

import (
  "github.com/labstack/echo/v4"
  "net/http"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  //e.Use(middleware.Logger())
  //e.Use(middleware.Recover())

  // Routes
  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello World")
})

e.Logger.Fatal(e.Start(":1323"))

}

