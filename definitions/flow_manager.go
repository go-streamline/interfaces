package definitions

import (
	"github.com/google/uuid"
)

// FlowManager defines the interface for managing and interacting with flows and processors within those flows.
type FlowManager interface {
	// GetFirstProcessorsForFlow retrieves the first processors in the execution order for a given flow.
	// Parameters:
	//   - flowID: The UUID of the flow.
	// Returns:
	//   - A slice of SimpleProcessor representing the first processors in the flow.
	//   - An error if the operation fails.
	GetFirstProcessorsForFlow(flowID uuid.UUID) ([]SimpleProcessor, error)

	// GetLastProcessorForFlow retrieves the last processor in the execution order for a given flow.
	// Parameters:
	//   - flowID: The UUID of the flow.
	// Returns:
	//   - A pointer to a SimpleProcessor representing the last processor in the flow.
	//   - An error if the operation fails.
	GetLastProcessorForFlow(flowID uuid.UUID) (*SimpleProcessor, error)

	// ListFlows retrieves a paginated list of flows.
	// Parameters:
	//   - pagination: A pointer to a PaginationRequest containing pagination information.
	// Returns:
	//   - A PaginatedData struct containing a slice of Flow pointers and the total count of flows.
	//   - An error if the operation fails.
	ListFlows(pagination *PaginationRequest) (PaginatedData[*Flow], error)

	// GetFlowByID retrieves a flow by its UUID.
	// Parameters:
	//   - flowID: The UUID of the flow to retrieve.
	// Returns:
	//   - A pointer to the Flow struct representing the retrieved flow.
	//   - An error if the operation fails.
	GetFlowByID(flowID uuid.UUID) (*Flow, error)

	// GetProcessorByID retrieves a processor by its UUID within a specific flow.
	// Parameters:
	//   - flowID: The UUID of the flow containing the processor.
	//   - processorID: The UUID of the processor to retrieve.
	// Returns:
	//   - A pointer to a SimpleProcessor representing the retrieved processor.
	//   - An error if the operation fails.
	GetProcessorByID(flowID uuid.UUID, processorID uuid.UUID) (*SimpleProcessor, error)

	// GetNextProcessors retrieves the processors that follow a specific processor in a flow.
	// Parameters:
	//   - flowID: The UUID of the flow containing the processors.
	//   - processorID: The UUID of the processor whose successors are to be retrieved.
	// Returns:
	//   - A slice of SimpleProcessor representing the next processors in the flow.
	//   - An error if the operation fails.
	GetNextProcessors(flowID uuid.UUID, processorID uuid.UUID) ([]SimpleProcessor, error)

	// AddProcessorToFlowBefore adds a processor to a flow before a specified reference processor.
	// Parameters:
	//   - flowID: The UUID of the flow to which the processor will be added.
	//   - processor: A pointer to the SimpleProcessor to add to the flow.
	//   - referenceProcessorID: The UUID of the processor before which the new processor will be added.
	// Returns:
	//   - An error if the operation fails.
	AddProcessorToFlowBefore(flowID uuid.UUID, processor *SimpleProcessor, referenceProcessorID uuid.UUID) error

	// AddProcessorToFlowAfter adds a processor to a flow after a specified reference processor.
	// Parameters:
	//   - flowID: The UUID of the flow to which the processor will be added.
	//   - processor: A pointer to the SimpleProcessor to add to the flow.
	//   - referenceProcessorID: The UUID of the processor after which the new processor will be added.
	// Returns:
	//   - An error if the operation fails.
	AddProcessorToFlowAfter(flowID uuid.UUID, processor *SimpleProcessor, referenceProcessorID uuid.UUID) error

	// SaveFlow saves the given flow and its associated processors to the persistent storage.
	// Parameters:
	//   - flow: A pointer to the Flow struct representing the flow to be saved.
	// Returns:
	//   - An error if the operation fails.
	SaveFlow(flow *Flow) error
}
