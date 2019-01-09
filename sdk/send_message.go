/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"
)

// SendMessage TODO...
func (session *Session) SendMessage(ctx context.Context, messageText string) (*Message, error) {
	/*
		request := &rpc.SendMessageRequest{SessionId: session.ID, Text: messageText}
		response, err := session.client.session.SendMessage(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "SendMessage failed")
		}
		return &Message{ID: response.Id}, nil
	*/
	return nil, nil // TODO
}
