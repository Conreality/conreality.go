/* This is free and unencumbered software released into the public domain. */

// Package conreality provides the Conreality Software Development Kit (SDK) for Go.
package conreality

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
)

const Version = "0.0.0"

type Asset struct {
	Object
}

type Binary struct {
	id uint64
}

type Camera struct {
	Object
}

type Client struct {
	db *sql.DB
}

type Event struct {
	id uint64
}

type Message struct {
	id uint64
}

type Object struct {
	uuid uuid.UUID
}

type Player struct {
	Object
}

type Session struct{}

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
	var db, err = sql.Open("postgres", "sslmode=disable dbname="+gameName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &Client{db: db}, nil
}

// Disconnect closes the connection to the master server.
func (client *Client) Disconnect() error {
	var err = client.db.Close()
	if err != nil {
		return err
	}
	client.db = nil
	return nil
}
