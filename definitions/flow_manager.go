package definitions

import (
	"github.com/go-streamline/interfaces/definitions/models"
	"github.com/google/uuid"
)

type FlowManager interface {
	GetFirstProcessorsForFlow(flowID uuid.UUID) ([]models.Processor, error)
	GetLastProcessorForFlow(flowID uuid.UUID) (*models.Processor, error)
	ListFlows() ([]models.Flow, error)
	GetFlowByID(flowID uuid.UUID) (*models.Flow, error)
	GetProcessorByID(flowID uuid.UUID, processorID uuid.UUID) (*models.Processor, error)
	GetNextProcessors(flowID uuid.UUID, processorID uuid.UUID) ([]models.Processor, error)
	ListProcessorsForFlow(flowID uuid.UUID) ([]models.Processor, error)
	AddProcessorToFlowBefore(flowID uuid.UUID, processor *models.Processor, referenceProcessorID uuid.UUID) error
	AddProcessorToFlowAfter(flowID uuid.UUID, processor *models.Processor, referenceProcessorID uuid.UUID) error
}
