/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// LeaveUnit TODO...
func (session *Session) LeaveUnit(ctx context.Context, unitID UnitID) error {
	request := &rpc.UnitID{Value: unitID}
	_, err := session.client.session.LeaveUnit(ctx, request)
	if err != nil {
		return errors.Wrap(err, "LeaveUnit failed")
	}
	return nil
}
