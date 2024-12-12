package definitions

import "github.com/sirupsen/logrus"

type LoggerFactory interface {
	GetLogger(name string) *logrus.Logger
}
