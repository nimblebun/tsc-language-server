package session

// State represents the state of the language server session (workspace)
type State int

const (
	// StateEmpty is the state of an empty session (before it starts)
	StateEmpty State = -1

	// StatePrepared is the state of a session that's ready to accept requests
	StatePrepared State = 0

	// StateInitializing is the state of a session after "initialize"
	StateInitializing State = 1

	// StateInitialized is the state of a session after "initialized"
	StateInitialized State = 2

	// StateShutdown is the state of a session after it has been shut down
	StateShutdown State = 3
)

// ToString will return a string representation of the given session state
func (state State) String() string {
	switch state {
	case StateEmpty:
		return "(empty)"
	case StatePrepared:
		return "prepared"
	case StateInitializing:
		return "initializing"
	case StateInitialized:
		return "initialized"
	case StateShutdown:
		return "shutdown"
	default:
		return "(unknown state)"
	}
}
