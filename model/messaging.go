package model

type SimpleMessage struct {
	Recipient      SimpleMessageRecipient `json:"recipient"`
	Messaging_type string                 `json:"messaging_type"`
	Message        string                 `json:"message"`
}

type SimpleMessageRecipient struct {
	Id string `json:"id"`
}

type SimpleMessageText struct {
	Text string `json:"text"`
}

func NewSimpleMessage() *SimpleMessage {
	return &SimpleMessage{}
}
