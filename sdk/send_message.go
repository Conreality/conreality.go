/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// SendMessage TODO...
func (session *Session) SendMessage(ctx context.Context, messageText string) (*Message, error) {
	request := &rpc.Message{
		SessionId: &rpc.SessionID{Value: session.ID},
		Text:      messageText,
	}
	response, err := session.client.session.SendMessage(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "SendMessage failed")
	}
	return &Message{ID: response.Value}, nil
}
