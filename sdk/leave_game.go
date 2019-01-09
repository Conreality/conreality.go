/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// LeaveGame TODO...
func (session *Session) LeaveGame(ctx context.Context, notice string) error {
	request := &rpc.String{Value: notice}
	_, err := session.client.session.LeaveGame(ctx, request)
	if err != nil {
		return errors.Wrap(err, "LeaveGame failed")
	}
	return nil
}
