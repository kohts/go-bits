package main

import (
	"fmt"
	"time"
)

type Assistant struct {
	ID string

	messageBus AssistantsMessageBus
}

func (a *Assistant) init() {
	messages := a.messageBus.ListenMessages(a.ID)
	fmt.Printf("%s: Assistant %s got channel and will start listening\n",
		time.Now().Format("2006-01-02 15:04:05.000 MST"), a.ID)

	for message := range messages {
		fmt.Printf("%s: Assistant %s BLOCKED until received message: %s\n",
			time.Now().Format("2006-01-02 15:04:05.000 MST"), a.ID, message.Content)
	}
}

type Message struct {
	SenderID     string
	Content      string
	ThreadID     string
	ReceiversIDs []string
}

type AssistantsMessageBus interface {
	SendMessage(message Message) error
	ListenMessages(assistantID string) <-chan Message
}

type AssistantsMessageBusImpl struct {
	assistants map[string]chan Message
}

func (a *AssistantsMessageBusImpl) SendMessage(message Message) error {
	if len(message.ReceiversIDs) == 0 {
		// Send to all assistants
		for _, assistant := range a.assistants {
			assistant <- message
		}
	} else {
		for _, receiverID := range message.ReceiversIDs {
			if assistant, ok := a.assistants[receiverID]; ok {
				assistant <- message
			}
		}
	}

	return nil
}

func (a *AssistantsMessageBusImpl) ListenMessages(assistantID string) <-chan Message {
	if _, ok := a.assistants[assistantID]; !ok {
		a.assistants[assistantID] = make(chan Message)
	}

	return a.assistants[assistantID]
}

func runAssistants() {
	messageBus := &AssistantsMessageBusImpl{
		assistants: make(map[string]chan Message),
	}

	assistant1 := &Assistant{
		ID:         "1",
		messageBus: messageBus,
	}
	go assistant1.init()

	assistant2 := &Assistant{
		ID:         "2",
		messageBus: messageBus,
	}
	go assistant2.init()

	assistant3 := &Assistant{
		ID:         "3",
		messageBus: messageBus,
	}
	go assistant3.init()

	fmt.Printf("%s: Assistants initialized, sleeping for 2 seconds...\n",
		time.Now().Format("2006-01-02 15:04:05.000 MST"))
	time.Sleep(2 * time.Second)

	fmt.Printf("%s: Sending first message\n",
		time.Now().Format("2006-01-02 15:04:05.000 MST"))

	messageBus.SendMessage(Message{
		SenderID: "1",
		Content:  "Hello from 1",
	})

	time.Sleep(2 * time.Second)

	fmt.Printf("%s: Sending second message\n",
		time.Now().Format("2006-01-02 15:04:05.000 MST"))
	messageBus.SendMessage(Message{
		SenderID: "2",
		Content:  "Second message",
	})

	time.Sleep(1 * time.Second)
}

func main() {
	runAssistants()
}
