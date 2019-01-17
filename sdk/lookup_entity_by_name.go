/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// LookupEntityByName TODO...
func (session *Session) LookupEntityByName(ctx context.Context, entityName string) (*Entity, error) {
	request := &rpc.String{Value: entityName}
	response, err := session.client.session.LookupEntityByName(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "LookupEntityByName failed")
	}
	if response.Value == 0 {
		return nil, nil // entity not found
	}
	return &Entity{ID: response.Value}, nil
}
