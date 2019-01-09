/* This is free and unencumbered software released into the public domain. */

package client

import (
	"context"

	"github.com/conreality/conreality.go/sdk/model"
	rpc "github.com/conreality/conreality.go/sdk/rpc"
	"github.com/pkg/errors"
)

// FormUnit TODO...
func (session *Session) FormUnit(ctx context.Context, unitName string) (*model.Unit, error) {
	request := &rpc.String{Value: unitName}
	response, err := session.client.session.FormUnit(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "FormUnit failed")
	}
	return &model.Unit{ID: response.Value}, nil
}
