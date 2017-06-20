/* This is free and unencumbered software released into the public domain. */

// Conreality Software Development Kit (SDK) for Go.
package conreality

import (
	"database/sql"
	"github.com/satori/go.uuid"
)

const Version = "0.0.0"

type Asset struct {
	uuid uuid.UUID
}

type Binary struct {
	id uint64
}

type Camera struct {
	uuid uuid.UUID
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
	uuid uuid.UUID
}

type Session struct{}

type Theater struct {
	uuid uuid.UUID
}
