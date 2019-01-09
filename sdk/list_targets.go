/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// ListTargets TODO...
func (session *Session) ListTargets(ctx context.Context, unitID UnitID) ([]*Target, error) {
	request := &rpc.UnitID{Value: unitID}
	response, err := session.client.session.ListTargets(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "ListTargets failed")
	}
	result := make([]*Target, len(response.Values))
	for index, element := range response.Values {
		result[index] = &Target{ID: TargetID(element.Value)}
	}
	return result, nil
}
