/* This is free and unencumbered software released into the public domain. */

package client

import (
	"github.com/conreality/conreality.go/sdk/model"
	rpc "github.com/conreality/conreality.go/sdk/rpc"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

// Ping sends a dummy packet to the master server.
func (client *Client) SendEvent(ctx context.Context, messageText string) (*model.Message, error) {
	var err error
	_, err = client.RPC.Ping(ctx, &rpc.PingRequest{}) // TODO
	if err != nil {
		return nil, errors.Wrap(err, "SendEvent failed")
	}
	return &model.Message{}, nil // TODO
}
