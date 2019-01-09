/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// JoinGame TODO...
func (session *Session) JoinGame(ctx context.Context) (*Player, error) {
	request := &rpc.Nothing{}
	_, err := session.client.session.JoinGame(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "JoinGame failed")
	}
	return nil, nil
}
