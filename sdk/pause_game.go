/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// PauseGame TODO...
func (session *Session) PauseGame(ctx context.Context, notice string) error {
	request := &rpc.TextString{Value: notice}
	_, err := session.client.session.PauseGame(ctx, request)
	if err != nil {
		return errors.Wrap(err, "PauseGame failed")
	}
	return nil
}
