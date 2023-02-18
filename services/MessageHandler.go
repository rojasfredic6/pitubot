package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	auth "pitubot/data"
	messaging "pitubot/model"

	"github.com/labstack/echo"
)

func HandleMessage(event messaging.SimpleTextMessage, c echo.Context) error {
	response := &messaging.SimpleTextMessageResponse{}
	response.Recipient.Id = event.Entry[0].Messaging[0].Sender.Id
	response.Messaging_Type = "RESPONSE"
	for _, entry := range event.Entry[0].Messaging {
		response.Message.Text = entry.Message.Text
		bytesObj, err1 := json.Marshal(response)
		if err1 != nil {
			return c.String(http.StatusBadRequest, err1.Error())
		}
		req, err2 := http.NewRequest("POST", auth.AuthDataInfo.SendURL(), bytes.NewBuffer(bytesObj))
		if err2 != nil {
			return c.String(http.StatusBadRequest, err2.Error())
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err3 := client.Do(req)
		if err3 != nil {
			return c.String(http.StatusBadRequest, err3.Error())
		}
		defer resp.Body.Close()
	}
	return nil
}
