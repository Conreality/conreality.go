/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// DesignateTarget TODO...
func (session *Session) DesignateTarget(ctx context.Context) (*Target, error) {
	request := &rpc.Target{
		SessionId: &rpc.SessionID{Value: session.ID},
		// TODO
	}
	targetID, err := session.client.session.DesignateTarget(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "DesignateTarget failed")
	}
	return &Target{ID: targetID.Value}, nil
}
