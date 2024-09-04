package definitions

import "github.com/google/uuid"

// BranchTracker is an interface that tracks the state of the processor branches for a flow.
type BranchTracker interface {
	// MarkProcessorComplete marks a processor as complete for a session.
	// It returns whether the current processor was the last one for this level and if the next level should be scheduled.
	MarkProcessorComplete(sessionID uuid.UUID, processorID uuid.UUID) (bool, error)

	// AddProcessorsForSession adds a set of processors to be tracked for the session at the current level.
	AddProcessorsForSession(sessionID uuid.UUID, processorIDs []uuid.UUID)

	// HasPendingProcessors returns true if there are processors in the current level that haven't completed.
	HasPendingProcessors(sessionID uuid.UUID) bool

	// ClearSession removes a session from the tracker once it's completed.
	ClearSession(sessionID uuid.UUID)
}
