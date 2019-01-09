/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// ListPlayers TODO...
func (session *Session) ListPlayers(ctx context.Context, unitID UnitID) ([]*Player, error) {
	request := &rpc.UnitID{Value: unitID}
	response, err := session.client.session.ListPlayers(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "ListPlayers failed")
	}
	result := make([]*Player, len(response.Values))
	for index, element := range response.Values {
		result[index] = &Player{Object: Object{ID: ObjectID(element.Value)}}
	}
	return result, nil
}
