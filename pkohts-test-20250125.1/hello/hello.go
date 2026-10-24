package main

import (
	"fmt"
	"log"
	"pkohts/greetings"
)

type Command struct {
	Name      string
	Arguments string
}

type Message struct {
	SenderID     string
	Type         string
	Content      string
	Command      *Command
	ThreadID     string
	ReceiversIDs []string
	TraceID      string
	// Stores a slack channel name associated with a message for EPMBotMessageRouter
	Metadata map[string]string
}

type openAIThreadData struct {
	messages       []Message
	OpenAIThreadId string
}
type openAIRequiredAction struct {
	Name      string
	Arguments string
	Result    string
	Done      bool
}
type openAIActiveRun struct {
	RunID           string
	ThreadData      *openAIThreadData
	RequiredActions map[string]*openAIRequiredAction
	done            chan struct{}
}

func main() {

	StudentAges := map[string]int{
		"Peter":    44,
		"Victoria": 43,
		"Olivia":   12,
	}
	log.Printf("StudentAges: %v", StudentAges)

	if _, ok := StudentAges["John"]; !ok {
		StudentAges["John"] = 15
	}
	delete(StudentAges, "Peter")
	log.Printf("StudentAges (excercise 2): %v", StudentAges)

	Courses := map[string][]string{
		"Math": {"Peter", "Victoria"},
		"Art":  {"Olivia", "Victoria"},
	}
	Courses["Math"] = append(Courses["Math"], "Olivia")
	for course, students := range Courses {
		log.Printf("Course %v: %v students", course, len(students))
	}

	EmptyMap := map[string]int{}
	log.Printf("EmptyMap len: %v", len(EmptyMap))

	threadData := &openAIThreadData{
		messages:       []Message{},
		OpenAIThreadId: "1",
	}
	activeRun := &openAIActiveRun{
		RunID:           "1",
		ThreadData:      threadData,
		done:            make(chan struct{}),
		RequiredActions: map[string]*openAIRequiredAction{},
	}
	log.Printf("activeRun: %v", activeRun)

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
