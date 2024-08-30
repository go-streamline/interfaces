package definitions

import (
	"github.com/google/uuid"
)

type FlowManager interface {
	GetFirstProcessorsForFlow(flowID uuid.UUID) ([]SimpleProcessor, error)
	GetLastProcessorForFlow(flowID uuid.UUID) (*SimpleProcessor, error)
	ListFlows(pagination *PaginationRequest) (PaginatedData[Flow], error)
	GetFlowByID(flowID uuid.UUID) (*Flow, error)
	GetProcessorByID(flowID uuid.UUID, processorID uuid.UUID) (*SimpleProcessor, error)
	GetNextProcessors(flowID uuid.UUID, processorID uuid.UUID) ([]SimpleProcessor, error)
	AddProcessorToFlowBefore(flowID uuid.UUID, processor *SimpleProcessor, referenceProcessorID uuid.UUID) error
	AddProcessorToFlowAfter(flowID uuid.UUID, processor *SimpleProcessor, referenceProcessorID uuid.UUID) error
	SaveFlow(flow *Flow) error
}
