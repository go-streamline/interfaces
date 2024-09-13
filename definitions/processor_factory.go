package definitions

// ProcessorFactory defines an interface for retrieving processors and trigger processors.
type ProcessorFactory interface {
	// GetProcessor retrieves a processor by its type name.
	// Parameters:
	// - id: The unique identifier of the processor.
	// - typeName: The type name of the processor to retrieve.
	// Returns:
	// - Processor: The retrieved processor.
	// - error: An error if the retrieval fails.
	GetProcessor(id string, typeName string) (Processor, error)

	// GetTriggerProcessor retrieves a trigger processor by its type name.
	// Parameters:
	// - id: The unique identifier of the trigger processor.
	// - typeName: The type name of the trigger processor to retrieve.
	// Returns:
	// - TriggerProcessor: The retrieved trigger processor.
	// - error: An error if the retrieval fails.
	GetTriggerProcessor(id string, typeName string) (TriggerProcessor, error)
}
