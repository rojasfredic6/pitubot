package main

import (
	"net/http"

	"github.com/labstack/echo"

	auth "pitubot/data"
	messaging "pitubot/model"
	handler "pitubot/services"
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
			handler.HandleMessage(message, c)
			// response := &messaging.SimpleTextMessageResponse{}
			// response.Recipient.Id = message.Entry[0].Messaging[0].Sender.Id
			// for _, entry := range message.Entry[0].Messaging {
			// 	response.Message.Text = entry.Message.Text
			// 	response.Messaging_Type = "RESPONSE"
			// 	bytesObj, err1 := json.Marshal(response)
			// 	if err1 != nil {
			// 		return c.String(http.StatusBadRequest, err1.Error())

			// 	}
			// 	req, err2 := http.NewRequest("POST", auth.AuthDataInfo.SendURL(), bytes.NewBuffer(bytesObj))
			// 	if err2 != nil {
			// 		return c.String(http.StatusBadRequest, err2.Error())
			// 	}

			// 	req.Header.Set("Content-Type", "application/json")

			// 	client := &http.Client{}
			// 	resp, err3 := client.Do(req)
			// 	if err3 != nil {
			// 		return c.String(http.StatusBadRequest, err3.Error())
			// 	}
			// 	defer resp.Body.Close()
			// }
		}
		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))

}
