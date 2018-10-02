/* This is free and unencumbered software released into the public domain. */

package client

import (
	rpc "github.com/conreality/conreality.go/sdk/rpc"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

// Ping sends a dummy packet to the master server.
func (client *Client) Ping(ctx context.Context) error {
	var err error
	_, err = client.public.Ping(ctx, &rpc.PingRequest{})
	if err != nil {
		return errors.Wrap(err, "Ping failed")
	}
	return nil
}
