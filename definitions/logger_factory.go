package definitions

import "github.com/sirupsen/logrus"

type LoggerFactory interface {
	// GetLogger retrieves a logger based on the specified type and name.
	// Parameters:
	// - loggerType: The type of logger to retrieve. For example: Human
	// - name: The name of the instance of the logger. For example, if I have multiple instances for Human,
	// I can use its name to differentiate between them.
	// Returns:
	// - *logrus.Logger: A pointer to the retrieved logger.
	GetLogger(loggerType string, name string) *logrus.Logger
}
