/* This is free and unencumbered software released into the public domain. */

package client

import (
	rpc "github.com/conreality/conreality.go/sdk/rpc"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

// Authenticate TODO...
func (client *Client) Authenticate(ctx context.Context, playerNick string) (*Session, error) {
	response, err := client.public.Authenticate(ctx, &rpc.AuthRequest{PlayerNick: playerNick})
	if err != nil {
		return nil, errors.Wrap(err, "Ping failed")
	}
	return &Session{client: client, ID: response.SessionId}, nil
}
