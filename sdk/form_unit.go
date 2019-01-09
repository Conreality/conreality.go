/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// FormUnit TODO...
func (session *Session) FormUnit(ctx context.Context, unitName string) (*Unit, error) {
	request := &rpc.String{Value: unitName}
	response, err := session.client.session.FormUnit(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "FormUnit failed")
	}
	return &Unit{ID: response.Value}, nil
}
