package utils

import (
	uuid "github.com/satori/go.uuid"
)

// NewUUIdV4 return uuid string of version 4
func NewUUIdV4() string {
	uid := uuid.NewV4()

	return uid.String()
}
