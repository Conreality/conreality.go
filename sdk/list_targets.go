/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"
)

// ListTargets TODO...
func (session *Session) ListTargets(ctx context.Context, unitID UnitID) ([]Target, error) {
	return make([]Target, 0), nil // TODO
}
