package definitions

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Flow represents a flow of processors that are executed in a specific order.
// It contains the flow's metadata such as its ID, name, description, and the processors that belong to it.
type Flow struct {
	ID          uuid.UUID         // ID uniquely identifies the flow.
	Name        string            // Name is the human-readable name of the flow.
	Description string            // Description provides a brief explanation of the flow's purpose.
	Processors  []SimpleProcessor // Processors is the list of processors that are part of the flow, executed in order.
}

// SimpleProcessor represents a processor within a flow.
// It contains metadata about the processor, including its configuration and execution parameters.
type SimpleProcessor struct {
	ID         uuid.UUID              // ID uniquely identifies the processor.
	FlowID     uuid.UUID              // FlowID is the ID of the flow to which the processor belongs.
	Name       string                 // Name is the human-readable name of the processor.
	Type       string                 // Type indicates the type of the processor, typically used to determine its implementation.
	FlowOrder  int                    // FlowOrder specifies the order in which the processor is executed within the flow.
	Config     map[string]interface{} // Config holds the specific configuration for the processor.
	MaxRetries int                    // MaxRetries specifies the maximum number of retries allowed for this processor in case of failure.
	LogLevel   logrus.Level           // LogLevel defines the logging level used for this processor.
}
