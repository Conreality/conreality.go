/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// PauseGame TODO...
func (session *Session) PauseGame(ctx context.Context) error {
	request := &rpc.Nothing{}
	_, err := session.client.session.PauseGame(ctx, request)
	if err != nil {
		return errors.Wrap(err, "PauseGame failed")
	}
	return nil
}
