package definitions

import (
	"github.com/go-streamline/interfaces/utils"
	"github.com/google/uuid"
	"io"
)

type EngineFlowObject struct {
	Metadata map[string]interface{} `json:"metadata"`
}

type EngineIncomingObject struct {
	FlowID    uuid.UUID
	Metadata  map[string]interface{}
	Reader    io.Reader
	SessionID uuid.UUID
}

func (e *EngineFlowObject) EvaluateExpression(input string) (string, error) {
	return utils.EvaluateExpression(input, e.Metadata)
}

type EngineFileHandler interface {
	ProcessorFileHandler
	GetInputFile() string
	GetOutputFile() string
	Close()
	GenerateNewFileHandler() (EngineFileHandler, error)
}

type SessionUpdate struct {
	SessionID uuid.UUID
	Finished  bool
	Error     error
}
