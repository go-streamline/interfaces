package definitions

import (
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"io"
)

type ScheduleType int

const (
	EventDriven ScheduleType = iota
	CronDriven
)

// BaseProcessor represents a basic processor.
type BaseProcessor struct {
}

type BaseProcessorInterface interface {
	// Name returns the name of the processor.
	// Returns:
	// - string: The name of the processor.
	Name() string

	// SetConfig sets the configuration for the processor.
	// Parameters:
	// - config: A map containing the configuration settings.
	// Returns:
	// - error: An error if setting the configuration fails.
	SetConfig(config map[string]interface{}) error

	// Close is called when the processor is being stopped or cleaned up.
	// Returns:
	// - error: An error if the close operation fails.
	Close() error
}

// TriggerProcessor defines the interface for a trigger processor.
type TriggerProcessor interface {
	BaseProcessorInterface

	// Execute executes the processor logic.
	// Parameters:
	// - info: The EngineFlowObject containing the execution information.
	// - fileHandler: The ProcessorFileHandler for handling file operations.
	// - log: The logger for logging information.
	// Returns:
	// - []*EngineFlowObject: A list of *EngineFlowObject after execution.
	// - error: An error if the execution fails.
	Execute(info *EngineFlowObject, fileHandler ProcessorFileHandler, log *logrus.Logger) ([]*EngineFlowObject, error)

	// GetScheduleType returns the scheduling type supported by the TriggerProcessor.
	GetScheduleType() ScheduleType

	// HandleSessionUpdate allows the TriggerProcessor to respond to session updates.
	HandleSessionUpdate(update SessionUpdate)
}

type ScheduleConfig struct {
	Type     ScheduleType
	CronExpr string // Used only if Type == CronDriven
}

// DecodeMap decodes a map into the output structure. This can be used to decode the configuration(which is a map) into its own struct.
// Parameters:
// - input: The input map to decode.
// - output: The output structure to decode into.
// Returns:
// - error: An error if the decoding fails.
func (b *BaseProcessor) DecodeMap(input interface{}, output interface{}) error {
	return mapstructure.Decode(input, output)
}

// Processor defines the interface for a processor.
type Processor interface {
	BaseProcessorInterface

	// Execute executes the processor logic.
	// Parameters:
	// - info: The EngineFlowObject containing the execution information.
	// - fileHandler: The ProcessorFileHandler for handling file operations.
	// - log: The logger for logging information.
	// Returns:
	// - *EngineFlowObject: The updated EngineFlowObject after execution.
	// - error: An error if the execution fails.
	Execute(info *EngineFlowObject, fileHandler ProcessorFileHandler, log *logrus.Logger) (*EngineFlowObject, error)
}

// ProcessorFileHandler defines the interface for handling the current contents.
type ProcessorFileHandler interface {
	// Read reads data from a file.
	// Returns:
	// - io.Reader: The reader for reading data.
	// - error: An error if the read operation fails.
	Read() (io.Reader, error)

	// Write writes data to a file.
	// Returns:
	// - io.Writer: The writer for writing data.
	// - error: An error if the write operation fails.
	Write() (io.Writer, error)
}
