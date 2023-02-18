package model

type SimpleTextMessage struct {
	Object string      `json:"object"`
	Entry  []EntryData `json:"entry"`
}

type EntryData struct {
	Id        string          `json:"id"`
	Time      int             `json:"time"`
	Messaging []MessagingData `json:"messaging"`
}

type MessagingData struct {
	Sender    SenderData    `json:"sender"`
	Recipient RecipientData `json:"recipient"`
	Timestamp int           `json:"timestamp"`
	Message   MessageData   `json:"message"`
}

type SenderData struct {
	Id string `json:"id"`
}

type RecipientData struct {
	Id string `json:"id"`
}

type MessageData struct {
	Mid  string `json:"mid"`
	Text string `json:"text"`
}

type SimpleTextMessageResponse struct {
	Recipient      RecipientResponse `json:"recipient"`
	Message        MessageResponse   `json:"message"`
	Messaging_Type string            `json:"messaging_type"`
}

type RecipientResponse struct {
	Id string `json:"id"`
}

type MessageResponse struct {
	Text string `json:"text"`
}
