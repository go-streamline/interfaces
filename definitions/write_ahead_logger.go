package definitions

import (
	"github.com/go-streamline/core/definitions"
	"github.com/google/uuid"
)

type LogEntry struct {
	SessionID     uuid.UUID
	ProcessorName string
	ProcessorID   string
	FlowID        uuid.UUID
	InputFile     string
	OutputFile    string
	FlowObject    definitions.EngineFlowObject
	RetryCount    int
}
type WriteAheadLogger interface {
	WriteEntry(entry LogEntry)
	ReadEntries() ([]LogEntry, error)
}
