/* This is free and unencumbered software released into the public domain. */

package client

import (
	rpc "github.com/conreality/conreality.go/sdk/rpc"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

// Hello performs a version handshake with the master server.
func (client *Client) Hello(ctx context.Context) (string, error) {
	response, err := client.RPC.Hello(ctx, &rpc.HelloRequest{Version: Version})
	if err != nil {
		return "", errors.Wrap(err, "Hello failed")
	}
	return response.Version, nil
}
