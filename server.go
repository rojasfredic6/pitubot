package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	auth "pitubot/data"
	sm "pitubot/model"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	e.GET("/webhook", func(c echo.Context) error {
		checkAuth := c.FormValue("hub.verify_token")
		challengeValue := c.FormValue("hub.challenge")
		if checkAuth == auth.AuthDataInfo.Auth {
			return c.String(http.StatusOK, challengeValue)
		} else {
			return c.String(http.StatusBadRequest, "Error!!")
		}
	})

	e.POST("/webhook", func(c echo.Context) error {
		fmt.Println(c.Request().GetBody())
		return nil
	})
	print(sm.NewSimpleMessage())

	e.Logger.Fatal(e.Start(":8080"))

}
