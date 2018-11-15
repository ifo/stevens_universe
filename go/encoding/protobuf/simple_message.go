package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

/* Remember, Message looks something like this:
type Message struct {
	Text   string
	Number int
}
*/

func main() {
	message := &Message{
		Text:   "text",
		Number: 1,
	}

	m, err := proto.Marshal(message)
	if err != nil {
		log.Fatalf("Marshalling error: %s\n", err)
	}

	newMessage := &Message{}
	err = proto.Unmarshal(m, newMessage)
	if err != nil {
		log.Fatalf("Unmarshalling error: %s\n", err)
	}

	if message.Text != newMessage.Text || message.Number != newMessage.Number {
		log.Fatalln("Oh no! Our messages did not match.")
	}

	fmt.Printf("Successfully Unmarshalled: %+v\n", newMessage)
}
