/* This is free and unencumbered software released into the public domain. */

package sdk

// Close TODO...
func (session *Session) Close() error {
	if session.client != nil {
		session.client = nil
	}
	return nil
}
