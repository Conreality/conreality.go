/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// UpdatePlayer TODO...
func (session *Session) UpdatePlayer(ctx context.Context, heartbeat uint32) error {
	request := &rpc.PlayerStatus{
		SessionId: &rpc.SessionID{Value: session.ID},
		Heartbeat: heartbeat,
	}
	_, err := session.client.session.UpdatePlayer(ctx, request)
	if err != nil {
		return errors.Wrap(err, "UpdatePlayer failed")
	}
	return nil
}
