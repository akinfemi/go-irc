package server

import "fmt"

//Message each message sent
type Message struct {
	Body     string
	Sender   User
	Reciever interface{}
	Threads  []Message
}

func (m Message) String() string {
	return fmt.Sprintf("Sender: %s\nReciver: %s\nBody: %s\n", m.Sender, m.Reciever, m.Body)
}
