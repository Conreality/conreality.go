/* This is free and unencumbered software released into the public domain. */

package client

import (
	rpc "github.com/conreality/conreality.go/src/rpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Connect attempts to connect to a local master server.
func Connect(masterHost string) (*Client, error) {
	conn, err := grpc.Dial(masterHost, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "connect failed")
	}
	return &Client{Conn: conn, RPC: rpc.NewMasterClient(conn)}, nil
}
