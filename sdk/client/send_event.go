/* This is free and unencumbered software released into the public domain. */

package client

import (
	"context"

	"github.com/conreality/conreality.go/sdk/model"
)

// SendEvent TODO...
func (session *Session) SendEvent(ctx context.Context, predicate string, subject model.Object, object model.Object) (*model.Event, error) {
	/*
		request := &rpc.SendEventRequest{SessionId: session.ID, Predicate: predicate, SubjectUuid: subject.UUID.String(), ObjectUuid: object.UUID.String()}
		response, err := session.client.session.SendEvent(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "SendEvent failed")
		}
		return &model.Event{ID: response.Id}, nil
	*/
	return nil, nil // TODO
}
