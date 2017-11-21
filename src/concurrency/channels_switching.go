package main

import (
	"fmt"
)

func main() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	// msg := Message{
	// 	To: []string{"apollo.arkwright@gmail.com"},
	// 	From: "root@localhost",
	// 	Content: "Status update",
	// }

	// failedMessage := FailedMessage{
	// 	ErrorMessage: "Generic error encountered!",
	// 	OriginalMessage: Message{},
	// }

	// msgCh <- msg
	// errCh <- failedMessage

	select {
		case receivedMsg := <- msgCh:
			fmt.Println(receivedMsg)
		case receivedErr := <- errCh:
			fmt.Println(receivedErr)
		default:
			fmt.Println("No messages received!")
	}
}

type Message struct {
	To 		[]string
	From 	string
	Content string
}

type FailedMessage struct {
	ErrorMessage	string
	OriginalMessage	Message
}