package definitions

import "github.com/google/uuid"

type BranchTracker interface {
	// AddProcessor track the processor as scheduled and pending completion
	AddProcessor(sessionID uuid.UUID, processorID uuid.UUID, nextProcessorIDs []uuid.UUID)

	// MarkComplete mark the processor as completed
	MarkComplete(sessionID uuid.UUID, processorID uuid.UUID) (allComplete bool)

	// IsComplete check if all processors in a session are completed
	IsComplete(sessionID uuid.UUID) bool

	RestoreState(sessionID uuid.UUID, completedProcessorIDs []uuid.UUID)

	GetCompletedProcessors(sessionID uuid.UUID) []uuid.UUID
}
