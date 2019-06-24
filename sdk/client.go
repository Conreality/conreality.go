/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	//rpc "github.com/conreality/conreality.go/rpc"
	"google.golang.org/grpc"
)

// Client connection state.
type Client struct {
	Connection *grpc.ClientConn
	//public     rpc.PublicClient
	//session    rpc.SessionClient
}
