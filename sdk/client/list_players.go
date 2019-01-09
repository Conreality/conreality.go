/* This is free and unencumbered software released into the public domain. */

package client

import (
	"context"

	"github.com/conreality/conreality.go/sdk/model"
)

// ListPlayers TODO...
func (session *Session) ListPlayers(ctx context.Context, unitID model.UnitID) ([]model.Player, error) {
	return make([]model.Player, 0), nil // TODO
}
