/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// DefineTheater TODO...
func (session *Session) DefineTheater(ctx context.Context) error { // TODO
	request := &rpc.Box{} // TODO
	_, err := session.client.session.DefineTheater(ctx, request)
	if err != nil {
		return errors.Wrap(err, "DefineTheater failed")
	}
	return nil
}
