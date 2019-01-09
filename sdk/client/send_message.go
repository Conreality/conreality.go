/* This is free and unencumbered software released into the public domain. */

package client

import (
	"github.com/conreality/conreality.go/sdk/model"
	"golang.org/x/net/context"
)

// SendMessage TODO...
func (session *Session) SendMessage(ctx context.Context, messageText string) (*model.Message, error) {
	/*
		request := &rpc.SendMessageRequest{SessionId: session.ID, Text: messageText}
		response, err := session.client.session.SendMessage(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "SendMessage failed")
		}
		return &model.Message{ID: response.Id}, nil
	*/
	return nil, nil // TODO
}
