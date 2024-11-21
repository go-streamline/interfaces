package definitions

import (
	"github.com/google/uuid"
	"time"
)

// FlowManager defines the interface for managing and interacting with flows and processors within those flows.
type FlowManager interface {
	// GetFirstProcessorsForFlow retrieves the first processors for a given flow.
	// Parameters:
	// - flowID: The unique identifier of the flow.
	// Returns:
	// - []SimpleProcessor: A slice of the first processors for the flow.
	// - error: An error if the retrieval fails.
	GetFirstProcessorsForFlow(flowID uuid.UUID) ([]*SimpleProcessor, error)

	// GetFlowProcessors retrieves all processors for a given flow.
	// Parameters:
	// - flowID: The unique identifier of the flow.
	// Returns:
	// - []SimpleProcessor: A slice of processors for the flow.
	// - error: An error if the retrieval fails.
	GetFlowProcessors(flowID uuid.UUID) ([]*SimpleProcessor, error)

	// GetTriggerProcessorsForFlow retrieves the trigger processors for a given flow.
	// Parameters:
	// - flowID: The unique identifier of the flow.
	// Returns:
	// - []*SimpleTriggerProcessor: A slice of the trigger processors for the flow.
	// - error: An error if the retrieval fails.
	GetTriggerProcessorsForFlow(flowID uuid.UUID) ([]*SimpleTriggerProcessor, error)

	// GetProcessors retrieves processors by their unique identifiers.
	// Parameters:
	// - processorIDs: A slice of unique identifiers for the processors.
	// Returns:
	// - []SimpleProcessor: A slice of processors with the specified IDs.
	// - error: An error if the retrieval fails.
	GetProcessors(processorIDs []uuid.UUID) ([]*SimpleProcessor, error)

	// ListFlows lists all flows with pagination and a time filter.
	// Parameters:
	// - pagination: The pagination request containing page number and items per page.
	// - since: The time filter to list flows updated since this time.
	// Returns:
	// - PaginatedData[*Flow]: The paginated data containing the flows.
	// - error: An error if the listing fails.
	ListFlows(pagination *PaginationRequest, since time.Time) (PaginatedData[*Flow], error)

	// GetFlowByID retrieves a flow by its unique identifier.
	// Parameters:
	// - flowID: The unique identifier of the flow.
	// Returns:
	// - *Flow: The flow with the specified ID.
	// - error: An error if the retrieval fails.
	GetFlowByID(flowID uuid.UUID) (*Flow, error)

	// GetProcessorByID retrieves a processor by its unique identifier within a flow.
	// Parameters:
	// - flowID: The unique identifier of the flow.
	// - processorID: The unique identifier of the processor.
	// Returns:
	// - *SimpleProcessor: The processor with the specified ID.
	// - error: An error if the retrieval fails.
	GetProcessorByID(flowID uuid.UUID, processorID uuid.UUID) (*SimpleProcessor, error)

	// AddProcessorToFlowBefore adds a processor to a flow before a reference processor.
	// Parameters:
	// - flowID: The unique identifier of the flow.
	// - processor: The processor to add.
	// - referenceProcessorID: The unique identifier of the reference processor.
	// Returns:
	// - error: An error if the addition fails.
	AddProcessorToFlowBefore(flowID uuid.UUID, processor *SimpleProcessor, referenceProcessorID uuid.UUID) error

	// AddProcessorToFlowAfter adds a processor to the flow after a reference processor.
	// Parameters:
	// - flowID: The unique identifier of the flow.
	// - processor: The processor to add.
	// - referenceProcessorID: The unique identifier of the reference processor.
	// Returns:
	// - error: An error if the addition fails.
	AddProcessorToFlowAfter(flowID uuid.UUID, processor *SimpleProcessor, referenceProcessorID uuid.UUID) error

	// SaveFlow saves a flow and its processors to the database.
	// Parameters:
	// - flow: The flow to save.
	// Returns:
	// - error: An error if the save operation fails.
	SaveFlow(flow *Flow) error

	// GetLastUpdateTime retrieves the last update time for a given flow.
	// Parameters:
	// - flowIDs: A slice of flow IDs.
	// Returns:
	// - map[uuid.UUID]time.Time: A map of flow IDs to their last update time.
	// - error: An error if the retrieval fails.
	GetLastUpdateTime(flowIDs []uuid.UUID) (map[uuid.UUID]time.Time, error)

	// SetFlowActive marks a flow as active or inactive.
	// Parameters:
	// - flowID: The unique identifier of the flow.
	// - active: A boolean indicating whether the flow should be marked as active.
	// Returns:
	// - error: An error if the operation fails.
	SetFlowActive(flowID uuid.UUID, active bool) error
}
