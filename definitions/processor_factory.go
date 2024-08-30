package definitions

// ProcessorFactory defines an interface for retrieving processors.
type ProcessorFactory interface {
	GetProcessor(typeName string) (Processor, error)
	RegisterProcessor(processor Processor)
}
