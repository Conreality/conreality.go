/* This is free and unencumbered software released into the public domain. */

package sdk

// Close TODO...
func (session *Session) Close() error {
	session.client = nil
	return nil
}
