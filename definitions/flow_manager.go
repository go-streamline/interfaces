package definitions

import (
	"github.com/google/uuid"
)

type FlowManager interface {
	GetInitialProcessors(flowID uuid.UUID) ([]*SimpleProcessor, error)
	GetNextProcessors(processorID uuid.UUID) ([]*SimpleProcessor, error)
	SaveFlowState(flow *SimpleFlow) error
}
