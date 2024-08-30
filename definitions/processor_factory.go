package definitions

// ProcessorFactory defines an interface for retrieving processors.
// This can be used to support all sorts of processors.
type ProcessorFactory interface {
	// GetProcessor retrieves a processor by its type name.
	// Parameters:
	// - typeName: The type name of the processor to retrieve.
	// Returns:
	// - Processor: The retrieved processor.
	// - error: An error if the retrieval fails.
	GetProcessor(typeName string) (Processor, error)

	// RegisterProcessor registers a new processor.
	// Parameters:
	// - processor: The processor to register.
	RegisterProcessor(processor Processor)

	// RegisterProcessorWithTypeName registers a new processor with a specific type name.
	// Parameters:
	// - typeName: The type name to associate with the processor.
	// - processor: The processor to register.
	RegisterProcessorWithTypeName(typeName string, processor Processor)
}
