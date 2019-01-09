/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// AddPlayer TODO...
func (session *Session) AddPlayer(ctx context.Context, playerNick string) (*Player, error) {
	request := &rpc.Player{Nick: playerNick}
	response, err := session.client.session.AddPlayer(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "AddPlayer failed")
	}
	return &Player{Object: Object{ID: ObjectID(response.Value)}}, nil
}
