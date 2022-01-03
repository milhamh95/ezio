package main

import "fmt"


type TopicMessage string

type Message struct {
	Topic   TopicMessage
	Msg string
}

func NewMessage(topic TopicMessage,msg string) Message {
	return Message{
		Topic: topic,
		Msg: msg,
	}
}

func (m Message) Get() string {
	return fmt.Sprintf("topic->%s, msg:%s", m.Topic,m.Msg)
}
