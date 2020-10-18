package handlers

import "log"

// MethodHandler is a struct that contains a logger, as well as methods to
// individual method handlers
type MethodHandler struct {
	logger *log.Logger
}

// NewMethodHandler will instantiate a handler that contains callbacks to the
// server's methods
func NewMethodHandler(logger *log.Logger) *MethodHandler {
	return &MethodHandler{logger}
}
