package definitions

import (
	"github.com/google/uuid"
)

// LogEntry is a struct that represents a log entry in the write ahead log.
type LogEntry struct {
	SessionID             uuid.UUID        `json:"session_id"`
	ProcessorName         string           `json:"processor_name"`
	ProcessorID           string           `json:"processor_id"`
	FlowID                uuid.UUID        `json:"flow_id"`
	InputFile             string           `json:"input_file"`
	OutputFile            string           `json:"output_file"`
	FlowObject            EngineFlowObject `json:"flow_object"`
	RetryCount            int              `json:"retry_count"`
	IsComplete            bool             `json:"is_complete"` // indicates if the processor completed
	CompletedProcessorIDs []uuid.UUID      `json:"completed_processor_ids"`
}

// WriteAheadLogger is an interface for writing log entries to a file and reading them back.
// The log entries are used to keep track of the state of the processing of a flow.
type WriteAheadLogger interface {
	WriteEntry(entry LogEntry)
	ReadEntries() ([]LogEntry, error)
	ReadLastEntries() ([]LogEntry, error)
}
