/* This is free and unencumbered software released into the public domain. */

// Package conreality provides the Conreality Software Development Kit (SDK) for Go.
package conreality

////////////////////////////////////////////////////////////////////////////////
// Client API

// Login TODO...
/*
func (client *Client) Login(agentUUID string, secret string) (*Session, error) {
	var uuid, err = uuid.FromString(agentUUID)
	if err != nil {
		return nil, errors.Wrap(err, "UUID parsing failed")
	}
	return &Session{client: client, agentUUID: uuid}, nil
}
*/

////////////////////////////////////////////////////////////////////////////////
// Session API

// Logout TODO...
/*
func (session *Session) Logout() error {
	session.client = nil
	return nil
}
*/

// Game returns the current game.
/*
func (session *Session) Game() *Game {
	return &Game{session: session}
}
*/

// NewAction creates a new transactional action.
/*
func (session *Session) NewAction() (*Action, error) {
	var tx, err = session.client.db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "BEGIN failed")
	}
	return &Action{client: session.client, tx: tx}, nil
}
*/

////////////////////////////////////////////////////////////////////////////////
// Game API

/* TODO */

////////////////////////////////////////////////////////////////////////////////
// Action API

// Abort TODO...
/*
func (action *Action) Abort() error {
	var err = action.tx.Rollback()
	if err != nil {
		return errors.Wrap(err, "ROLLBACK failed")
	}
	action.tx = nil
	return nil
}
*/

// Commit TODO...
/*
func (action *Action) Commit() error {
	var err = action.tx.Commit()
	if err != nil {
		return errors.Wrap(err, "COMMIT failed")
	}
	action.tx = nil
	return nil
}
*/

// SendEvent TODO...
/*
func (action *Action) SendEvent(predicate string, subject, object *Object) (*Event, error) {
	var result sql.NullInt64
	var err = action.tx.QueryRow("SELECT conreality.event_send($1, $2, $3) AS id",
		predicate, subject.uuid, object.uuid).Scan(&result)
	if err != nil {
		return nil, errors.Wrap(err, "CALL conreality.event_send failed")
	}
	if !result.Valid {
		panic("unexpected NULL result from conreality.event_send")
	}
	return &Event{id: result.Int64}, nil
}
*/

// SendMessage TODO...
/*
func (action *Action) SendMessage(messageText string) (*Message, error) {
	var result sql.NullInt64
	var err = action.tx.QueryRow("SELECT conreality.message_send($1) AS id", messageText).Scan(&result)
	if err != nil {
		return nil, errors.Wrap(err, "CALL conreality.message_send failed")
	}
	if !result.Valid {
		panic("unexpected NULL result from conreality.message_send")
	}
	return &Message{id: result.Int64}, nil
}
*/
