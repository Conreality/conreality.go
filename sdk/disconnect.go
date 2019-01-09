/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"github.com/pkg/errors"
)

// Disconnect closes the connection to the master server.
func (client *Client) Disconnect() error {
	if client.Connection != nil {
		if err := client.Connection.Close(); err != nil {
			return errors.Wrap(err, "Disconnect failed")
		}
		client.Connection = nil
	}
	return nil
}
