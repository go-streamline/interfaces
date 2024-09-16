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

	// Nodes retrieves the list of nodes participating in the leader election.
	// Returns:
	// - []string: A slice of node identifiers.
	// - error: An error if the retrieval fails.
	Nodes() ([]string, error)

	// Close terminates the leader election process.
	// Returns:
	// - error: An error if the termination fails.
	Close() error
}
