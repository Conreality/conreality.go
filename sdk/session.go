/* This is free and unencumbered software released into the public domain. */

package sdk

// Session
type Session struct {
	client *Client
	ID     uint64
}

// Close TODO...
func (session *Session) Close() error {
	session.client = nil
	return nil
}
