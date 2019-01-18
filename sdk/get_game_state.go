/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// GetGameState TODO...
func (session *Session) GetGameState(ctx context.Context) (string, error) {
	request := &rpc.Nothing{}
	response, err := session.client.session.GetGameState(ctx, request)
	if err != nil {
		return "", errors.Wrap(err, "GetGameState failed")
	}
	return response.Value, nil
}
