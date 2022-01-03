package main

import "fmt"

func main() {
	msg := "this is a message abc"
	evtTopic := TopicEvent("topic.event")
	msgTopic := TopicMessage("topic.message")
	e, err := InitializeEvent(evtTopic,msgTopic,msg)
	if err != nil {
		fmt.Println("error")
		return
	}
	e.Start()
}
