/* This is free and unencumbered software released into the public domain. */

package client

import (
	rpc "github.com/conreality/conreality.go/sdk/rpc"
	"google.golang.org/grpc"
)

const Version = "0.0.0"

// Client connection state.
type Client struct {
	Conn *grpc.ClientConn
	RPC  rpc.MasterClient
}
