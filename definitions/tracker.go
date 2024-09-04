package definitions

import "github.com/google/uuid"

// BranchTrackerResult represents the result after a processor completes.
type BranchTrackerResult struct {
	IsLastInBranch     bool // IsLastInBranch indicates if this processor was the last in its branch.
	ShouldScheduleNext bool // ShouldScheduleNext indicates if the next level of processors should be scheduled.
}

// BranchTracker is an interface that tracks the state of processor branches for a flow.
type BranchTracker interface {
	// MarkProcessorComplete marks a processor as complete for a session.
	// It returns a BranchTrackerResult that indicates if the current processor was the last one in its branch to finish,
	// and whether the next level of processors should be scheduled.
	MarkProcessorComplete(sessionID uuid.UUID, processorID uuid.UUID) BranchTrackerResult

	// AddBranch adds a new processor branch for a session.
	AddBranch(sessionID uuid.UUID, processorIDs []uuid.UUID)

	// IsBranchComplete checks if all processors in a specific branch have completed for a session.
	IsBranchComplete(sessionID uuid.UUID) bool

	// ClearSession removes tracking data for a completed session.
	ClearSession(sessionID uuid.UUID)
}
