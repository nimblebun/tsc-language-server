package session

import (
	"context"
	"log"

	"github.com/creachadair/jrpc2"
)

// ServiceSession is an interface used to instantiate a JSON-RPC service
type ServiceSession interface {
	Assigner() (jrpc2.Assigner, error)
	Finish(jrpc2.Assigner, jrpc2.ServerStatus)
	SetLogger(*log.Logger)
}

// Factory is a callback function used for creating a ServiceSession with
// the provided context
type Factory func(context.Context) ServiceSession
