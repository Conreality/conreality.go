/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// Authenticate TODO...
func (client *Client) Authenticate(ctx context.Context, playerNick string) (*Session, error) {
	request := &rpc.AuthRequest{PlayerNick: playerNick}
	response, err := client.public.Authenticate(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "Ping failed")
	}
	return &Session{client: client, ID: response.SessionId}, nil
}
