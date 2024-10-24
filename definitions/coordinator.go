package definitions

import "github.com/google/uuid"

// Coordinator defines the interface for assigning leaders nodes for each trigger processor.
// The coordinator elects a leader for each trigger processor and monitors the leader nodes.
// If a node drops, the coordinator is responsible for reassigning the leader.
type Coordinator interface {
	// IsLeader checks if the current instance is the leader for the trigger processor.
	// If the current instance is the coordinator, it will also assign the leader for the trigger processor.
	// Parameters:
	// - tpID: The unique identifier of the trigger processor.
	// Returns:
	// - bool: True if the current instance is the leader, false otherwise.
	// - error: An error if the check fails.
	IsLeader(tpID uuid.UUID) (bool, error)

	// Close terminates the leader election process and stops being the coordinator(if currently is).
	// Returns:
	// - error: An error if the termination fails.
	Close() error
}
