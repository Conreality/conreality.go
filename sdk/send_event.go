/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// SendEvent TODO...
func (session *Session) SendEvent(ctx context.Context, predicate string, subjectID ObjectID, objectID ObjectID) (*Event, error) {
	request := &rpc.Event{
		SessionId: &rpc.SessionID{Value: session.ID},
		Predicate: predicate,
		SubjectId: &rpc.ObjectID{Value: subjectID},
		ObjectId:  &rpc.ObjectID{Value: objectID},
	}
	response, err := session.client.session.SendEvent(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "SendEvent failed")
	}
	return &Event{ID: response.Value}, nil
}
