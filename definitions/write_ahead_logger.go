package definitions

import (
	"github.com/google/uuid"
)

// LogEntry is a struct that represents a log entry in the write ahead log.
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

// WriteAheadLogger is an interface for writing log entries to a file and reading them back.
// The log entries are used to keep track of the state of the processing of a flow.
type WriteAheadLogger interface {
	WriteEntry(entry LogEntry)
	ReadEntries() ([]LogEntry, error)
}
