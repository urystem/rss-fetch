package domain

import "errors"

var (
	ErrConflict    = errors.New("ErrConflict")
	ErrNotAffected = errors.New("ErrNotAffected")
	ErrFlag        = errors.New("error usage")
)
