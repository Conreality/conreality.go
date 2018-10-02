/* This is free and unencumbered software released into the public domain. */

package model

import uuid "github.com/satori/go.uuid"

// Object
type Object struct {
	uuid uuid.UUID
}

// NewObject TODO...
func NewObject(objectUUID string) *Object {
	var uuid, err = uuid.FromString(objectUUID)
	if err != nil {
		return nil
	}
	return &Object{uuid: uuid}
}
