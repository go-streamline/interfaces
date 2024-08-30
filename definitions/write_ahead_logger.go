package definitions

import (
	"github.com/google/uuid"
)

type LogEntry struct {
	SessionID     uuid.UUID
	ProcessorName string
	ProcessorID   string
	FlowID        uuid.UUID
	InputFile     string
	OutputFile    string
	FlowObject    EngineFlowObject
	RetryCount    int
}
type WriteAheadLogger interface {
	WriteEntry(entry LogEntry)
	ReadEntries() ([]LogEntry, error)
}
