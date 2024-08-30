package definitions

import (
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"io"
)

type BaseProcessor struct {
	ID string
}

func (b *BaseProcessor) GetID() string {
	return b.ID
}

func (b *BaseProcessor) DecodeMap(input interface{}, output interface{}) error {
	return mapstructure.Decode(input, output)
}

type Processor interface {
	GetID() string
	Name() string
	Execute(info *EngineFlowObject, fileHandler ProcessorFileHandler, log *logrus.Logger) (*EngineFlowObject, error)
	SetConfig(config map[string]interface{}) error
}

type ProcessorFileHandler interface {
	Read() (io.Reader, error)
	Write() (io.Writer, error)
}
