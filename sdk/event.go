/* This is free and unencumbered software released into the public domain. */

package sdk

import "context"

// EventID is a unique identifier for a game event.
type EventID = uint64

// Event represents a game event.
type Event struct {
	ID EventID
}

// ReceiveEvents subscribes to the stream of game events.
func (connection *GameConnection) ReceiveEvents(ctx context.Context, startID EventID) (<-chan *Event, error) {
	channel := make(chan *Event)
	// TODO
	return channel, nil
}
