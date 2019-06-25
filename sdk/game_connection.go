/* This is free and unencumbered software released into the public domain. */

package sdk

import "context"

// GameConnection represents an active connection to a game endpoint.
type GameConnection struct {
	Endpoint *GameEndpoint
}

// ConnectToGame TODO...
func ConnectToGame(ctx context.Context, endpoint *GameEndpoint) (*GameConnection, error) {
	// TODO
	return &GameConnection{Endpoint: endpoint}, nil
}
