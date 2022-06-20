package logs

import (
	"github.com/sirupsen/logrus"
)

type service struct {
	logger *logrus.Logger
}

func LogrusAdapter() *service {
	return &service{
		logger: logrus.New(),
	}
}

func (s service) Info(args ...interface{}) {
	s.logger.Info(args)
}

func (s service) Debug(args ...interface{}) {
	s.logger.Debug(args)
}

func (s service) Error(args ...interface{}) {
	s.logger.Error(args)
}

func (s service) Warn(args ...interface{}) {
	s.logger.Warn(args)
}
