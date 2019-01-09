/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// Ping sends a dummy packet to the master server.
func (client *Client) Ping(ctx context.Context) error {
	_, err := client.public.Ping(ctx, &rpc.Nothing{})
	if err != nil {
		return errors.Wrap(err, "Ping failed")
	}
	return nil
}
