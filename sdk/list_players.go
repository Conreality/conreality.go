/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"
)

// ListPlayers TODO...
func (session *Session) ListPlayers(ctx context.Context, unitID UnitID) ([]Player, error) {
	return make([]Player, 0), nil // TODO
}
