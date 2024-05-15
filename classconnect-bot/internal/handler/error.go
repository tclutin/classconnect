package handler

import "errors"

var (
	ErrInternal     = errors.New("⚠️ An internal error has occurred")
	ErrNotAvailable = errors.New("💋 This section is not available to you")
)
