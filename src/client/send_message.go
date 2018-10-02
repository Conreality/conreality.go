/* This is free and unencumbered software released into the public domain. */

package client

import (
	"github.com/conreality/conreality.go/src/model"
	rpc "github.com/conreality/conreality.go/src/rpc"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

// SendMessage
func (client *Client) SendMessage(ctx context.Context, messageText string) (*model.Message, error) {
	var err error
	_, err = client.RPC.Ping(ctx, &rpc.PingRequest{}) // TODO
	if err != nil {
		return nil, errors.Wrap(err, "SendMessage failed")
	}
	return &model.Message{}, nil // TODO
}
