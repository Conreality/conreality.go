/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// DefineMission TODO...
func (session *Session) DefineMission(ctx context.Context, summary string) error {
	request := &rpc.TextString{Value: summary}
	_, err := session.client.session.DefineMission(ctx, request)
	if err != nil {
		return errors.Wrap(err, "DefineMission failed")
	}
	return nil
}
