//go:build wireinject

package main

import "github.com/google/wire"

func InitializeEvent(evtTopic TopicEvent,msgTopic TopicMessage ,userMessage string) (Event, error) {
	wire.Build(
		NewEvent,
		NewMessage,
		)
	return Event{}, nil
}
