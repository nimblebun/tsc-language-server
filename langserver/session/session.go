package session

import (
	"context"
	"fmt"
	"time"

	"github.com/creachadair/jrpc2"
)

// Session is a struct that represents an individual language server workspace
type Session struct {
	initializingReq     *jrpc2.Request
	initializingReqTime time.Time

	initializedReq     *jrpc2.Request
	initializedReqTime time.Time

	downReq     *jrpc2.Request
	downReqTime time.Time

	state      State
	cancelFunc context.CancelFunc
}

func (s *Session) isPrepared() bool {
	return s.state == StatePrepared
}

func (s *Session) isInitializing() bool {
	return s.state == StateInitializing
}

func (s *Session) isInitialized() bool {
	return s.state == StateInitialized
}

func (s *Session) isShutdown() bool {
	return s.state == StateShutdown
}

func (s *Session) isExitable() bool {
	return s.isShutdown() || s.isPrepared()
}

// Init will attempt to initialize a session
func (s *Session) Init(req *jrpc2.Request) error {
	if !s.isPrepared() {
		if s.isInitializing() {
			return fmt.Errorf("Session already initializing, req ID: %s", s.initializingReq.ID())
		}

		return fmt.Errorf("Session is not ready. Current state: %s", s.state)
	}

	s.initializingReq = req
	s.initializedReqTime = time.Now()
	s.state = StateInitializing

	return nil
}

// Prepare will attempt to set a session's state to "prepared"
func (s *Session) Prepare() error {
	if s.state != StateEmpty {
		return fmt.Errorf("Failed to prepare session. Expected state %s, got state %s",
			StateInitialized, s.state)
	}

	s.state = StatePrepared

	return nil
}

// EnsureInitialized will return an error if the current session is not
// initialized
func (s *Session) EnsureInitialized() error {
	if !s.isInitialized() {
		return fmt.Errorf("Session not initialized. Current state: %d", s.state)
	}

	return nil
}

// FinishInitialization will attempt to set the state of the session to
// "initialized"
func (s *Session) FinishInitialization(req *jrpc2.Request) error {
	if !s.isInitializing() {
		if s.isInitialized() {
			return fmt.Errorf("Session already initialized at %s via request %s",
				s.initializedReqTime, s.initializedReq.ID())
		}

		return fmt.Errorf("Cannot initialize session because it's not ready (current state: %s)", s.state)
	}

	s.initializedReq = req
	s.initializedReqTime = time.Now()
	s.state = StateInitialized

	return nil
}

// Shutdown will attempt to shut down an active session
func (s *Session) Shutdown(req *jrpc2.Request) error {
	if s.isShutdown() {
		return fmt.Errorf("Session already shut down in request %s", s.downReq.ID())
	}

	s.downReq = req
	s.downReqTime = time.Now()
	s.state = StateShutdown

	return nil
}

// Exit will attempt to exit out of an exitable session
func (s *Session) Exit() error {
	if !s.isExitable() {
		return fmt.Errorf("Session is not ready to be exited. Current state: %s", s.state)
	}

	s.cancelFunc()

	return nil
}

// GetState will return the current state of a session
func (s *Session) GetState() State {
	return s.state
}

// New instantiates a new language server workspace session
func New(cancelFunc context.CancelFunc) *Session {
	return &Session{
		state:      StateEmpty,
		cancelFunc: cancelFunc,
	}
}
