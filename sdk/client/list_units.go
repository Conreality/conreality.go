/* This is free and unencumbered software released into the public domain. */

package client

import (
	"context"

	"github.com/conreality/conreality.go/sdk/model"
)

// ListUnits TODO...
func (session *Session) ListUnits(ctx context.Context, unitID model.UnitID) ([]model.Unit, error) {
	return make([]model.Unit, 0), nil // TODO
}
