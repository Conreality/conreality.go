/* This is free and unencumbered software released into the public domain. */

// Package conreality provides the Conreality Software Development Kit (SDK) for Go.
package conreality

import (
	"database/sql"
	_ "github.com/lib/pq" // for side effects only
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

////////////////////////////////////////////////////////////////////////////////
// Constants

// Version contains the current package version, as a string.
const Version = "0.0.0"

////////////////////////////////////////////////////////////////////////////////
// Type Definitions

// Action
type Action struct {
	client *Client
	tx     *sql.Tx
}

// Asset
type Asset struct {
	Object
}

// Binary
type Binary struct {
	id int64
}

// Camera
type Camera struct {
	Object
}

// Client
type Client struct {
	db *sql.DB
}

// Event
type Event struct {
	id int64
}

// Game
type Game struct {
	session *Session
}

// Message
type Message struct {
	id int64
}

// Object
type Object struct {
	uuid uuid.UUID
}

// Player
type Player struct {
	Object
}

// Session
type Session struct {
	client    *Client
	agentUUID uuid.UUID
}

// Theater
type Theater struct {
	uuid uuid.UUID
}

////////////////////////////////////////////////////////////////////////////////
// Object API

// NewObject TODO...
func NewObject(objectUUID string) *Object {
	var uuid, err = uuid.FromString(objectUUID)
	if err != nil {
		return nil
	}
	return &Object{uuid: uuid}
}

////////////////////////////////////////////////////////////////////////////////
// Client API

// Connect attempts to connect to a local master server.
//
// The returned handle is safe for concurrent use by multiple goroutines and
// maintains its own internal pool of idle connections. Thus, the Connect
// function should be called just once. It is rarely necessary to close a
// handle.
func Connect(masterHost string) (*Client, error) {
	var db, err = sql.Open("postgres", "sslmode=disable user=00000000-0000-0000-0000-000000000000 dbname="+masterHost) // FIXME
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "ping failed")
	}
	return &Client{db: db}, nil
}

// Disconnect closes the connection to the master server.
func (client *Client) Disconnect() error {
	var err = client.db.Close()
	if err != nil {
		return errors.Wrap(err, "close failed")
	}
	client.db = nil
	return nil
}

// Login TODO...
func (client *Client) Login(agentUUID string, secret string) (*Session, error) {
	var uuid, err = uuid.FromString(agentUUID)
	if err != nil {
		return nil, errors.Wrap(err, "UUID parsing failed")
	}
	return &Session{client: client, agentUUID: uuid}, nil
}

////////////////////////////////////////////////////////////////////////////////
// Session API

// Logout TODO...
func (session *Session) Logout() error {
	session.client = nil
	return nil
}

// Game returns the current game.
func (session *Session) Game() *Game {
	return &Game{session: session}
}

// NewAction creates a new transactional action.
func (session *Session) NewAction() (*Action, error) {
	var tx, err = session.client.db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "BEGIN failed")
	}
	return &Action{client: session.client, tx: tx}, nil
}

////////////////////////////////////////////////////////////////////////////////
// Game API

/* TODO */

////////////////////////////////////////////////////////////////////////////////
// Action API

// Abort TODO...
func (action *Action) Abort() error {
	var err = action.tx.Rollback()
	if err != nil {
		return errors.Wrap(err, "ROLLBACK failed")
	}
	action.tx = nil
	return nil
}

// Commit TODO...
func (action *Action) Commit() error {
	var err = action.tx.Commit()
	if err != nil {
		return errors.Wrap(err, "COMMIT failed")
	}
	action.tx = nil
	return nil
}

// SendEvent TODO...
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

// SendMessage TODO...
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
