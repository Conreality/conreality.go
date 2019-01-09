/* This is free and unencumbered software released into the public domain. */

package sdk

// SessionID TODO...
type SessionID = uint64

// Session TODO...
type Session struct {
	client *Client
	ID     SessionID
}
