/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"
)

// ListUnits TODO...
func (session *Session) ListUnits(ctx context.Context, unitID UnitID) ([]Unit, error) {
	return make([]Unit, 0), nil // TODO
}
