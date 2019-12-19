package errors

import (
	gerrors "errors"
)

var (
	// ErrInconsistentIDs ...
	ErrInconsistentIDs = gerrors.New("inconsistent IDs")
	// ErrAlreadyExists ...
	ErrAlreadyExists = gerrors.New("already exists")
	// ErrNotFound ...
	ErrNotFound = gerrors.New("not found")
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = gerrors.New("inconsistent mapping between route and handler (programmer error)")

	ErrInvalidArgument = gerrors.New("invalid argument")
)
