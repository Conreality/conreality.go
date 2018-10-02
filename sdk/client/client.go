/* This is free and unencumbered software released into the public domain. */

package client

import (
	rpc "github.com/conreality/conreality.go/sdk/rpc"
	"google.golang.org/grpc"
)

// Client connection state.
type Client struct {
	Connection *grpc.ClientConn
	public     rpc.PublicClient
	session    rpc.SessionClient
}
