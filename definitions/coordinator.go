package definitions

// Coordinator defines the interface for coordinating leader election and handling leader failures across a cluster.
type Coordinator interface {
	// IsLeader checks if the current instance is the leader for the trigger processor.
	// If the current instance is the coordinator, it will also assign the leader for the trigger processor.
	// Parameters:
	// - tpID: The unique identifier of the trigger processor.
	// Returns:
	// - bool: True if the current instance is the leader, false otherwise.
	// - error: An error if the check fails.
	IsLeader(tpID string) (bool, error)

	// Close terminates the leader election process and stops being the coordinator(if currently is).
	// Returns:
	// - error: An error if the termination fails.
	Close() error
}
