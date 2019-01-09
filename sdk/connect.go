/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Connect attempts to connect to a local master server.
func Connect(masterURL string) (*Client, error) {
	// TODO: parse the URL
	conn, err := grpc.Dial(masterURL, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "connect failed")
	}
	return &Client{Connection: conn, public: rpc.NewPublicClient(conn), session: rpc.NewSessionClient(conn)}, nil
}
