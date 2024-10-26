package definitions

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Flow represents a flow of processors that are executed in a specific order.
// It contains the flow's metadata such as its ID, name, description, and the processors that belong to it.
type Flow struct {
	ID                uuid.UUID                `json:"id"`                 // ID uniquely identifies the flow.
	Name              string                   `json:"name"`               // Name is the human-readable name of the flow.
	Description       string                   `json:"description"`        // Description provides a brief explanation of the flow's purpose.
	Processors        []SimpleProcessor        `json:"processors"`         // Processors is the list of processors that are part of the flow, executed in order.
	TriggerProcessors []SimpleTriggerProcessor `json:"trigger_processors"` // TriggerProcessors is the list of trigger processors associated with the flow.
	Active            bool                     `json:"active"`             // Active indicates whether the flow is currently active.
	LastUpdated       int64                    `json:"last_updated"`       // LastUpdated indicates the last update timestamp for the flow.
}

// SimpleProcessor represents a processor within a flow.
// It contains metadata about the processor, including its configuration and execution parameters.
type SimpleProcessor struct {
	ID               uuid.UUID              `json:"id"`              // ID uniquely identifies the processor.
	FlowID           uuid.UUID              `json:"flow_id"`         // FlowID is the ID of the flow to which the processor belongs.
	Name             string                 `json:"name"`            // Name is the human-readable name of the processor.
	Type             string                 `json:"type"`            // Type indicates the type of the processor, used to determine its implementation.
	Config           map[string]interface{} `json:"config"`          // Config holds the specific configuration for the processor.
	MaxRetries       int                    `json:"max_retries"`     // MaxRetries specifies the maximum number of retries allowed for this processor in case of failure.
	BackoffSeconds   int                    `json:"backoff_seconds"` // Specifies the delay in seconds between retry attempts for this processor in case of failure.
	LogLevel         logrus.Level           `json:"log_level"`       // LogLevel defines the logging level used for this processor.
	Enabled          bool                   `json:"enabled"`         // Enabled indicates if the processor is enabled.
	NextProcessorIDs []uuid.UUID            `json:"next_processors"` // NextProcessorIDs explicitly defines the processors that should run after this one.
}

// SimpleTriggerProcessor represents a trigger processor within a flow.
// It contains metadata about the trigger processor, including its configuration and scheduling parameters.
type SimpleTriggerProcessor struct {
	ID           uuid.UUID              `json:"id"`            // ID uniquely identifies the trigger processor.
	FlowID       uuid.UUID              `json:"flow_id"`       // FlowID is the ID of the flow to which the trigger processor belongs.
	Name         string                 `json:"name"`          // Name is the human-readable name of the trigger processor.
	Type         string                 `json:"type"`          // Type indicates the type of the trigger processor, used to determine its implementation.
	Config       map[string]interface{} `json:"config"`        // Config holds the specific configuration for the trigger processor.
	ScheduleType ScheduleType           `json:"schedule_type"` // ScheduleType defines the type of scheduling (event-driven or cron-driven).
	CronExpr     string                 `json:"cron_expr"`     // CronExpr is the cron expression used for scheduling, applicable if ScheduleType is CronDriven.
	EventName    string                 `json:"event_name"`    // EventName is the name of the event to listen to, applicable if ScheduleType is EventDriven.
	LogLevel     logrus.Level           `json:"log_level"`     // LogLevel defines the logging level used for this trigger processor.
	Enabled      bool                   `json:"enabled"`       // Enabled indicates if the processor is enabled.
}
