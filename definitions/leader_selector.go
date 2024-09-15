package definitions

// LeaderSelector defines the interface for leader election mechanisms.
type LeaderSelector interface {
	// Start initiates the leader election process.
	// Returns:
	// - error: An error if the initiation fails.
	Start() error

	// IsLeader checks if the current instance is the leader.
	// Returns:
	// - bool: True if the current instance is the leader, false otherwise.
	// - error: An error if the check fails.
	IsLeader() (bool, error)

	// Stop terminates the leader election process.
	// Returns:
	// - error: An error if the termination fails.
	Stop() error
}
