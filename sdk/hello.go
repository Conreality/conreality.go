/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// Hello performs a version handshake with the master server.
func (client *Client) Hello(ctx context.Context) (string, error) {
	response, err := client.public.Hello(ctx, &rpc.HelloRequest{Version: Version})
	if err != nil {
		return "", errors.Wrap(err, "Hello failed")
	}
	return response.Version, nil
}
