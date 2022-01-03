package main

import "fmt"

type TopicEvent string

type Event struct {
	Topic   TopicEvent
	Message Message
}

func NewEvent(topic TopicEvent, msg Message) (Event, error) {
	if topic == "" {
		return Event{}, fmt.Errorf("empty topic")
	}
	return Event{
		Topic:   topic,
		Message: msg,
	}, nil
}
func (e Event) Start() {
	msg := e.Message.Get()
	fmt.Printf("topic:%s, message:%s", e.Topic, msg)
}
