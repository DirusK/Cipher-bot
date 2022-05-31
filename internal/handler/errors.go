package handler

import "errors"

var (
	// ErrInvalidKeyLength is returned when length of key is not valid.
	ErrInvalidKeyLength = errors.New("key length is not valid")

	// ErrInvalidFormatKey is returned when format of key is not valid.
	ErrInvalidFormatKey = errors.New("key format is not valid")

	// ErrKeyExpired is returned when key is expired.
	ErrKeyExpired = errors.New("key expired")
)
