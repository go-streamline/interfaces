package definitions

type StateType string

const (
	StateTypeCluster StateType = "cluster"
	StateTypeLocal   StateType = "local"
)

type StateManager interface {
	// Reset resets the state manager including closing it before then.
	// Returns:
	// - error: An error if the reset fails.
	Reset() error

	// Close closes the state manager.
	// Returns:
	// - error: An error if the close operation fails.
	Close() error

	// GetState retrieves the state for the given state type.
	// Parameters:
	// - stateType: The type of state to retrieve.
	// Returns:
	// - map[string]any: A map representing the state.
	// - error: An error if the retrieval fails.
	GetState(stateType StateType) (map[string]any, error)

	// SetState sets the state for the given state type.
	// Parameters:
	// - stateType: The type of state to set.
	// - value: A map representing the state to set.
	// Returns:
	// - error: An error if the set operation fails.
	SetState(stateType StateType, value map[string]any) error

	// WatchState watches for changes in the given state type.
	// Parameters:
	// - state: The type of state to watch.
	// - callback: A callback function to invoke when the state changes.
	// Returns:
	// - error: An error if the watch operation fails.
	WatchState(state StateType, callback func()) error
}
