package main

import (
	"net/http"

	"github.com/labstack/echo"

	auth "pitubot/data"
	messaging "pitubot/model"
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
		var message messaging.SimpleTextMessage

		err := c.Bind(&message)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		if len(message.Entry) > 0 && len(message.Entry[0].Messaging) > 0 {
			for _, entry := range message.Entry[0].Messaging {
				print(entry.Message.Text)
			}
		}

		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))

}
