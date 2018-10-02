/* This is free and unencumbered software released into the public domain. */

package client

import (
	"github.com/pkg/errors"
)

// Disconnect closes the connection to the master server.
func (client *Client) Disconnect() error {
	var err = client.Connection.Close()
	if err != nil {
		return errors.Wrap(err, "close failed")
	}
	client.Connection = nil
	return nil
}
