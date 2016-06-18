package models

import (
	"time"
	"strconv"
)


// Message type that will be transmitted over the wire
// MessageType can be one of admin, cmd, info where
// admin can contain info such as how many sockets are connected, etc
// cmd may tell the browser to refresh or load a specific part
// info to push some string message to browser
type Message struct {
	Timestamp string
	Data map[string]string
}

func NewMessage(data map[string]string) Message  {
	var message Message
	message.Timestamp = strconv.FormatInt(time.Now().UnixNano(), 10)
	message.Data = data
	return message
}
