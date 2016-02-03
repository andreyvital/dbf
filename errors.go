package main

import "errors"

var (
	ErrUnexpectedHeaderSize = errors.New("unexpected header size")
	ErrUnsupportedVersion   = errors.New("unsupported version")
	ErrInvalidDBF           = errors.New("invalid dBASE file")
)
