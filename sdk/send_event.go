/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"
)

// SendEvent TODO...
func (session *Session) SendEvent(ctx context.Context, predicate string, subject Object, object Object) (*Event, error) {
	/*
		request := &rpc.SendEventRequest{SessionId: session.ID, Predicate: predicate, SubjectUuid: subject.UUID.String(), ObjectUuid: object.UUID.String()}
		response, err := session.client.session.SendEvent(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "SendEvent failed")
		}
		return &Event{ID: response.Id}, nil
	*/
	return nil, nil // TODO
}
