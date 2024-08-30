package definitions

import (
	"github.com/google/uuid"
)

type Flow struct {
	ID          uuid.UUID
	Name        string
	Description string
	Processors  []SimpleProcessor
}

type SimpleProcessor struct {
	ID         uuid.UUID
	FlowID     uuid.UUID
	Name       string
	Type       string
	FlowOrder  int
	Config     map[string]interface{}
	MaxRetries int
	LogLevel   string
}
