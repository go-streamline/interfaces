package definitions

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type SimpleFlow struct {
	ID         uuid.UUID
	Name       string
	Processors []*SimpleProcessor
}

// SimpleProcessor represents a minimal processor structure for the engine to process.
type SimpleProcessor struct {
	ID         uuid.UUID
	Name       string
	Type       string
	Config     map[string]interface{}
	MaxRetries int
	LogLevel   logrus.Level
	Next       []*SimpleProcessor
}
