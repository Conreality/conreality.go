/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// ListUnits TODO...
func (session *Session) ListUnits(ctx context.Context, unitID UnitID) ([]*Unit, error) {
	request := &rpc.UnitID{Value: unitID}
	response, err := session.client.session.ListUnits(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "ListUnits failed")
	}
	result := make([]*Unit, len(response.Values))
	for index, element := range response.Values {
		result[index] = &Unit{ID: UnitID(element.Value)}
	}
	return result, nil
}
