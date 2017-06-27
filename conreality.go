/* This is free and unencumbered software released into the public domain. */

// Package conreality provides the Conreality Software Development Kit (SDK) for Go.
package conreality

import (
	"database/sql"
	_ "github.com/lib/pq" // for side effects only
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"strconv"
)

// Version contains the current package version, as a string.
const Version = "0.0.0"

// Asset
type Asset struct {
	Object
}

// Binary
type Binary struct {
	id uint64
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
	id uint64
}

// Message
type Message struct {
	id uint64
}

// Object
type Object struct {
	uuid uuid.UUID
}

// Player
type Player struct {
	Object
}

// Scope
type Scope struct {
	tx *sql.Tx
}

// Session
type Session struct{}

// Theater
type Theater struct {
	uuid uuid.UUID
}

// Connect attempts to connect to a local master server.
//
// The returned handle is safe for concurrent use by multiple goroutines and
// maintains its own internal pool of idle connections. Thus, the Connect
// function should be called just once. It is rarely necessary to close a
// handle.
func Connect(gameName string) (*Client, error) {
	var db, err = sql.Open("postgres", "sslmode=disable user=00000000-0000-0000-0000-000000000000 dbname="+gameName)
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

// Begin creates a new transaction scope.
func (client *Client) Begin() (*Scope, error) {
	var tx, err = client.db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "BEGIN failed")
	}
	return &Scope{tx: tx}, nil
}

// Abort TODO...
func (scope *Scope) Abort() error {
	var err = scope.tx.Rollback()
	if err != nil {
		return errors.Wrap(err, "ROLLBACK failed")
	}
	scope.tx = nil
	return nil
}

// Commit TODO...
func (scope *Scope) Commit() error {
	var err = scope.tx.Commit()
	if err != nil {
		return errors.Wrap(err, "COMMIT failed")
	}
	scope.tx = nil
	return nil
}

// SendMessage TODO...
func (scope *Scope) SendMessage(messageText string) (int64, error) {
	var result sql.NullString
	var err = scope.tx.QueryRow("SELECT conreality.message_send($1) AS id", messageText).Scan(&result)
	if err != nil {
		return -1, errors.Wrap(err, "conreality.message_send failed")
	}
	if !result.Valid {
		panic("unexpected NULL result from conreality.message_send")
	}
	var messageID, _ = strconv.Atoi(result.String)
	return int64(messageID), nil
}
