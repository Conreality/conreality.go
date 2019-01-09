/* This is free and unencumbered software released into the public domain. */

package client

import (
	"context"

	"github.com/conreality/conreality.go/sdk/model"
)

// ListTargets TODO...
func (session *Session) ListTargets(ctx context.Context, unitID model.UnitID) ([]model.Target, error) {
	return make([]model.Target, 0), nil // TODO
}
